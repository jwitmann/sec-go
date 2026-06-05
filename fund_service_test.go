package sec

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestListAMCs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/amcs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":10,"next_cursor":"","items":[{"unique_id":"C0000000021","comp_name_th":"บลจ. กรุงศรี","comp_name_en":"Krungthai Asset Management","last_upd_date":"2024-01-15T00:00:00Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	amcs, cursor, err := c.ListAMCs(ctx, 10, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(amcs) != 1 {
		t.Fatalf("expected 1 AMC, got %d", len(amcs))
	}
	if amcs[0].CompNameEN != "Krungthai Asset Management" {
		t.Errorf("unexpected name: %s", amcs[0].CompNameEN)
	}
	if cursor != "" {
		t.Errorf("expected empty cursor, got %q", cursor)
	}
}

func TestListAMCsWithPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("page_size") != "5" {
			t.Errorf("expected page_size=5, got %q", q.Get("page_size"))
		}
		if q.Get("next_cursor") != "abc" {
			t.Errorf("expected cursor=abc, got %q", q.Get("next_cursor"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":5,"next_cursor":"def","items":[]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	_, cursor, err := c.ListAMCs(ctx, 5, "abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cursor != "def" {
		t.Errorf("expected cursor def, got %q", cursor)
	}
}

func TestGetFundProfiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/profiles" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("company_info") != "C0000000021" {
			t.Errorf("unexpected company_info: %q", q.Get("company_info"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":10,"items":[{"proj_id":"PRINCIPALi9","proj_name_th":"พรินซิเพิล แอ็กทีฟ อินคัม","proj_name_en":"Principal Active Income","fund_status":"RG"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	profiles, _, err := c.GetFundProfiles(ctx, ProfileOptions{CompanyInfo: "C0000000021"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(profiles) != 1 {
		t.Fatalf("expected 1 profile, got %d", len(profiles))
	}
	if profiles[0].ProjID != "PRINCIPALi9" {
		t.Errorf("unexpected proj_id: %s", profiles[0].ProjID)
	}
}

func TestGetDailyNAV(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/daily-info/nav" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "PRINCIPALi9" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("start_nav_date") != "2024-01-01" {
			t.Errorf("unexpected start_nav_date: %q", q.Get("start_nav_date"))
		}
		if q.Get("end_nav_date") != "2024-01-02" {
			t.Errorf("unexpected end_nav_date: %q", q.Get("end_nav_date"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"PRINCIPALi9","nav_date":"2024-01-02","last_val":10.5,"net_asset":1000000.0}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	navs, _, err := c.GetDailyNAV(ctx, NAVOptions{
		ProjID:    "PRINCIPALi9",
		StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(navs) != 1 {
		t.Fatalf("expected 1 NAV, got %d", len(navs))
	}
	if navs[0].LastVal != 10.5 {
		t.Errorf("unexpected last_val: %f", navs[0].LastVal)
	}
}

func TestGetMutualFundFees(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/mutual-fund-fees" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"PRINCIPALi9","fee_type_desc":"Management Fee","rate":1.5}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	fees, _, err := c.GetMutualFundFees(ctx, FeeOptions{ProjID: "PRINCIPALi9"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fees) != 1 || fees[0].Rate != 1.5 {
		t.Errorf("unexpected fees: %+v", fees)
	}
}

func TestGetFactsheetFees(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/fees" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "PRINCIPALi9" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("latest") != "true" {
			t.Errorf("expected latest=true, got %q", q.Get("latest"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"PRINCIPALi9","fee_type_desc":"Management Fee","rate":1.5}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	fees, _, err := c.GetFactsheetFees(ctx, FactsheetOptions{ProjID: "PRINCIPALi9", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fees) != 1 || fees[0].Rate != 1.5 {
		t.Errorf("unexpected fees: %+v", fees)
	}
}

func TestGetDividendHistory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/daily-info/dividend-history" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("class_abbr_name") != "PRINCIPAL-A" {
			t.Errorf("unexpected class_abbr_name: %q", q.Get("class_abbr_name"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"PRINCIPALi9","dividend_value":0.5}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	dividends, _, err := c.GetDividendHistory(ctx, DividendHistoryOptions{ClassAbbrName: "PRINCIPAL-A"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dividends) != 1 || dividends[0].DividendValue != 0.5 {
		t.Errorf("unexpected dividends: %+v", dividends)
	}
}

func TestGetAssetAllocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/asset-allocation" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"PRINCIPALi9","asset_name":"Equity","asset_ratio":60.0}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	allocations, _, err := c.GetAssetAllocation(ctx, FactsheetOptions{ProjID: "PRINCIPALi9", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(allocations) != 1 || allocations[0].AssetRatio != 60.0 {
		t.Errorf("unexpected allocations: %+v", allocations)
	}
}

func TestServiceEndpointError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"internal"}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL), WithMaxRetries(0))
	ctx := context.Background()

	_, _, err := c.ListAMCs(ctx, 10, "")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "list AMCs") {
		t.Errorf("expected error to mention operation, got: %v", err)
	}
}

func TestServiceInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`not json`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	_, _, err := c.ListAMCs(ctx, 10, "")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "unmarshal") {
		t.Errorf("expected unmarshal error, got: %v", err)
	}
}
