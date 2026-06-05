package sec

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestSearchFunds(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if r.URL.Path != "/v2/fund/general-info/profiles" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("project_info") != "alpha" {
			t.Errorf("unexpected project_info: %q", q.Get("project_info"))
		}

		cursor := `""`
		items := `[{"proj_id":"KT-Alpha","proj_name_th":"กรุงศรี อัลฟ่า","proj_name_en":"Krungthai Alpha","fund_status":"RG"}]`
		if callCount == 1 {
			cursor = `"next"`
			items = `[{"proj_id":"KT-Beta","proj_name_th":"กรุงศรี เบต้า","proj_name_en":"Krungthai Beta","fund_status":"RG"}]`
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":100,"next_cursor":` + cursor + `,"items":` + items + `}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	profiles, err := c.SearchFunds(context.Background(), "alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(profiles) != 2 {
		t.Fatalf("expected 2 profiles across pages, got %d", len(profiles))
	}
	if profiles[0].ProjID != "KT-Beta" {
		t.Errorf("unexpected first proj_id: %s", profiles[0].ProjID)
	}
	if profiles[1].ProjID != "KT-Alpha" {
		t.Errorf("unexpected second proj_id: %s", profiles[1].ProjID)
	}
}

func TestGetFundsByCompany(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/profiles" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("company_info") != "Krungthai Asset Management" {
			t.Errorf("unexpected company_info: %q", q.Get("company_info"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":100,"next_cursor":"","items":[{"proj_id":"KT-Alpha","proj_name_th":"กรุงศรี อัลฟ่า","proj_name_en":"Krungthai Alpha","fund_status":"RG"},{"proj_id":"KT-Beta","proj_name_th":"กรุงศรี เบต้า","proj_name_en":"Krungthai Beta","fund_status":"RG"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	profiles, err := c.GetFundsByCompany(context.Background(), "Krungthai Asset Management")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(profiles) != 2 {
		t.Fatalf("expected 2 profiles, got %d", len(profiles))
	}
	if profiles[0].ProjID != "KT-Alpha" || profiles[1].ProjID != "KT-Beta" {
		t.Errorf("unexpected profiles: %v", profiles)
	}
}

func TestFindAMC(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/amcs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":100,"next_cursor":"","items":[{"unique_id":"C0000000021","comp_name_th":"บลจ. กรุงศรี","comp_name_en":"Krungthai Asset Management","last_upd_date":"2024-01-15T00:00:00Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))

	tests := []struct {
		name  string
		query string
	}{
		{"by english name", "Krungthai"},
		{"by thai name", "กรุงศรี"},
		{"by unique_id", "C0000000021"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amc, err := c.FindAMC(context.Background(), tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if amc.UniqueID != "C0000000021" {
				t.Errorf("unexpected unique_id: %s", amc.UniqueID)
			}
		})
	}
}

func TestFindAMC_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":100,"next_cursor":"","items":[]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	_, err := c.FindAMC(context.Background(), "NonExistent")
	if err == nil {
		t.Fatal("expected error for not found")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected not found error, got: %v", err)
	}
}

func TestGetFundProfile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("proj_id") != "KT-Alpha" {
			t.Errorf("unexpected proj_id: %q", r.URL.Query().Get("proj_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","proj_name_en":"Krungthai Alpha","fund_status":"RG"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	profile, err := c.GetFundProfile(context.Background(), "KT-Alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if profile.ProjID != "KT-Alpha" {
		t.Errorf("unexpected proj_id: %s", profile.ProjID)
	}
}

func TestGetFundProfile_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	_, err := c.GetFundProfile(context.Background(), "MISSING")
	if err == nil {
		t.Fatal("expected error for not found")
	}
}

func TestGetFundLatestNAV(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/daily-info/nav" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("proj_id") != "KT-Alpha" {
			t.Errorf("unexpected proj_id: %q", r.URL.Query().Get("proj_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","nav_date":"2024-01-15","last_val":10.5}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	nav, err := c.GetFundLatestNAV(context.Background(), "KT-Alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if nav.LastVal != 10.5 {
		t.Errorf("unexpected last_val: %f", nav.LastVal)
	}
}

func TestGetFundRiskSpectrum(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/risk-spectrum" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "KT-Alpha" || q.Get("latest") != "true" {
			t.Errorf("unexpected query: %v", q)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","risk_spectrum":"RS5","risk_spectrum_desc":"Moderately high risk"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	rs, err := c.GetFundRiskSpectrum(context.Background(), "KT-Alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rs.RiskSpectrum != "RS5" {
		t.Errorf("unexpected risk spectrum: %s", rs.RiskSpectrum)
	}
}

func TestGetFundPortfolio(t *testing.T) {
	var requestCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.Add(1)
		q := r.URL.Query()
		if q.Get("proj_id") != "KT-Alpha" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v2/fund/factsheet/asset-allocation":
			_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","asset_seq":1,"asset_name":"หุ้น","asset_ratio":60.0}]}`))
		case "/v2/fund/factsheet/top5-holdings":
			_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","asset_seq":1,"asset_name":"PTT","asset_ratio":10.0}]}`))
		case "/v2/fund/outstanding/portfolio":
			_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","period":202401,"assetliab_desc":"หุ้น","assetliab_value":1000000.0}]}`))
		case "/v2/fund/outstanding/portfolio-asset-type":
			_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"KT-Alpha","period":202401,"assetliab_desc":"หุ้น","market_value":1000000.0}]}`))
		default:
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	portfolio, err := c.GetFundPortfolio(context.Background(), "KT-Alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if portfolio.ProjID != "KT-Alpha" {
		t.Errorf("unexpected proj_id: %s", portfolio.ProjID)
	}
	if len(portfolio.AssetAllocation) != 1 {
		t.Errorf("expected 1 asset allocation, got %d", len(portfolio.AssetAllocation))
	}
	if len(portfolio.Top5Holdings) != 1 {
		t.Errorf("expected 1 top 5 holding, got %d", len(portfolio.Top5Holdings))
	}
	if len(portfolio.QuarterlyPortfolio) != 1 {
		t.Errorf("expected 1 quarterly portfolio, got %d", len(portfolio.QuarterlyPortfolio))
	}
	if len(portfolio.MonthlyAssetBreakdown) != 1 {
		t.Errorf("expected 1 monthly asset breakdown, got %d", len(portfolio.MonthlyAssetBreakdown))
	}
	if requestCount.Load() != 4 {
		t.Errorf("expected 4 concurrent requests, got %d", requestCount.Load())
	}
}

func TestGetFundPortfolio_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v2/fund/factsheet/asset-allocation" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL), WithMaxRetries(0))
	_, err := c.GetFundPortfolio(context.Background(), "KT-Alpha")
	if err == nil {
		t.Fatal("expected error")
	}
}
