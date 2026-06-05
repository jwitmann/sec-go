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

func TestGetFactsheetSubscriptionRedemptionMinimums(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/subscription-redemption-minimums" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("latest") != "true" {
			t.Errorf("expected latest=true, got %q", q.Get("latest"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","fund_class_name":"main","start_date":"2022-06-30","end_date":"2022-07-26","prospectus_type":"Monthly","minimum_sub_ipo":5000,"minimum_sub_ipo_cur":"THB","minimum_sub":100,"minimum_sub_cur":"THB","minimum_sub_unit":"","minimum_redempt":0,"minimum_redempt_cur":"THB","minimum_redempt_unit":"","lowbal_val":0,"lowbal_val_cur":"THB","lowbal_unit":"","last_upd_date":"2022-07-26T07:53:25Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	minimums, _, err := c.GetFactsheetSubscriptionRedemptionMinimums(ctx, FactsheetOptions{ProjID: "M0000_2552", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(minimums) != 1 {
		t.Fatalf("expected 1 minimum record, got %d", len(minimums))
	}
	if minimums[0].MinimumSubIPO != 5000 {
		t.Errorf("unexpected minimum_sub_ipo: %f", minimums[0].MinimumSubIPO)
	}
	if minimums[0].MinimumSub != 100 {
		t.Errorf("unexpected minimum_sub: %f", minimums[0].MinimumSub)
	}
	if minimums[0].LowbalValCur != "THB" {
		t.Errorf("unexpected lowbal_val_cur: %s", minimums[0].LowbalValCur)
	}
}

func TestGetFactsheetSubscriptionRedemptionPeriods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/subscription-redemption-periods" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0512_2564" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("fund_class_name") != "AIA-ICA" {
			t.Errorf("unexpected fund_class_name: %q", q.Get("fund_class_name"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0512_2564","fund_class_name":"AIA-ICA","start_date":"2023-01-31","end_date":"2023-02-27","prospectus_type":"Monthly","type":"subscription","period":"ทุกวันทำการ","redemp_period_oth":"","settlement_period":"","last_upd_date":"2023-02-27T03:14:20Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	periods, _, err := c.GetFactsheetSubscriptionRedemptionPeriods(ctx, FactsheetOptions{ProjID: "M0512_2564", FundClassName: "AIA-ICA"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(periods) != 1 {
		t.Fatalf("expected 1 period record, got %d", len(periods))
	}
	if periods[0].Type != "subscription" {
		t.Errorf("unexpected type: %s", periods[0].Type)
	}
	if periods[0].Period != "ทุกวันทำการ" {
		t.Errorf("unexpected period: %s", periods[0].Period)
	}
	if periods[0].SettlementPeriod != "" {
		t.Errorf("unexpected settlement_period: %s", periods[0].SettlementPeriod)
	}
}

func TestGetFactsheetStatistics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/statistics" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0027_2541" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0027_2541","fund_class_name":"ABCC","start_date":"2023-07-31","end_date":"2023-08-30","prospectus_type":"Monthly","portfolio_turnover_ratio":"24.63","recovering_period":"1 เดือน","portfolio_duration_period":"1 เดือน 13 วัน","maximum_drawdown":"-0.02","sharpe_ratio":"0","beta":"0","alpha":"0","fx_hedging":"0","tracking_error":"0","yield_to_maturity":"2026-01-05","last_upd_date":"2023-08-28T11:15:36Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	stats, _, err := c.GetFactsheetStatistics(ctx, FactsheetOptions{ProjID: "M0027_2541"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stats) != 1 {
		t.Fatalf("expected 1 statistics record, got %d", len(stats))
	}
	if stats[0].PortfolioTurnoverRatio != "24.63" {
		t.Errorf("unexpected portfolio_turnover_ratio: %s", stats[0].PortfolioTurnoverRatio)
	}
	if stats[0].MaximumDrawdown != "-0.02" {
		t.Errorf("unexpected maximum_drawdown: %s", stats[0].MaximumDrawdown)
	}
	if stats[0].YieldToMaturity != "2026-01-05" {
		t.Errorf("unexpected yield_to_maturity: %s", stats[0].YieldToMaturity)
	}
}

func TestGetFundSpecifications(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/specifications" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","fund_class_name":"main","spec_code":"EQ","spec_desc":"Equity Fund","last_upd_date":"2022-07-26T07:53:25Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	specs, _, err := c.GetFundSpecifications(ctx, FeeOptions{ProjID: "M0000_2552"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(specs) != 1 {
		t.Fatalf("expected 1 specification, got %d", len(specs))
	}
	if specs[0].SpecCode != "EQ" {
		t.Errorf("unexpected spec_code: %s", specs[0].SpecCode)
	}
	if specs[0].SpecDesc != "Equity Fund" {
		t.Errorf("unexpected spec_desc: %s", specs[0].SpecDesc)
	}
}

func TestGetFactsheetBenchmarks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/benchmarks" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("latest") != "true" {
			t.Errorf("expected latest=true, got %q", q.Get("latest"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","start_date":"2022-06-30","end_date":"2022-07-26","prospectus_type":"Monthly","group_seq":1,"benchmark":"ดัชนีผลตอบแทนรวมตลาดหลักทรัพย์แห่งประเทศไทย (SET TRI)","benchmark_remark":"","last_upd_date":"2022-07-26T07:53:25Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	benchmarks, _, err := c.GetFactsheetBenchmarks(ctx, FactsheetOptions{ProjID: "M0000_2552", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(benchmarks) != 1 {
		t.Fatalf("expected 1 benchmark, got %d", len(benchmarks))
	}
	if benchmarks[0].GroupSeq != 1 {
		t.Errorf("unexpected group_seq: %d", benchmarks[0].GroupSeq)
	}
	if benchmarks[0].Benchmark != "ดัชนีผลตอบแทนรวมตลาดหลักทรัพย์แห่งประเทศไทย (SET TRI)" {
		t.Errorf("unexpected benchmark: %s", benchmarks[0].Benchmark)
	}
}

func TestGetFundInvolveParties(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/involve-parties" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("entity_type") != "R" {
			t.Errorf("unexpected entity_type: %q", q.Get("entity_type"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","entity_type":"R","entity_name_en":"MFC ASSET MANAGEMENT PUBLIC COMPANY LIMITED","entity_name_th":"บริษัทหลักทรัพย์จัดการกองทุน เอ็มเอฟซี จำกัด (มหาชน)","address":"เลขที่ 199","last_upd_date":"2025-11-19T07:22:16Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	parties, _, err := c.GetFundInvolveParties(ctx, InvolvePartyOptions{ProjID: "M0000_2552", EntityType: "R"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(parties) != 1 {
		t.Fatalf("expected 1 involve party, got %d", len(parties))
	}
	if parties[0].EntityType != "R" {
		t.Errorf("unexpected entity_type: %s", parties[0].EntityType)
	}
	if parties[0].EntityNameEN != "MFC ASSET MANAGEMENT PUBLIC COMPANY LIMITED" {
		t.Errorf("unexpected entity_name_en: %s", parties[0].EntityNameEN)
	}
}

func TestGetFundFactsheetURLs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/urls" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("fund_class_name") != "HIDIV-AR" {
			t.Errorf("unexpected fund_class_name: %q", q.Get("fund_class_name"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","fund_class_name":"HIDIV-AR","prospectus_type":"Monthly","amc_url_factsheet":"https://documents.mfcfund.com/Website/FundFiles/Q&A/QA_HIDIV-AR.pdf","pdf_factsheet":"https://secdocumentstorage.blob.core.windows.net/fundfactsheet/M0000_2552.pdf","as_of_date":"2025-09-30","last_upd_date":"2025-10-31T03:32:17Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	urls, _, err := c.GetFundFactsheetURLs(ctx, FeeOptions{ProjID: "M0000_2552", FundClassName: "HIDIV-AR"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(urls) != 1 {
		t.Fatalf("expected 1 URL record, got %d", len(urls))
	}
	if urls[0].AMCURLFactsheet != "https://documents.mfcfund.com/Website/FundFiles/Q&A/QA_HIDIV-AR.pdf" {
		t.Errorf("unexpected amc_url_factsheet: %s", urls[0].AMCURLFactsheet)
	}
	if urls[0].PDFFactsheet != "https://secdocumentstorage.blob.core.windows.net/fundfactsheet/M0000_2552.pdf" {
		t.Errorf("unexpected pdf_factsheet: %s", urls[0].PDFFactsheet)
	}
}

func TestGetFundIPOs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/ipos" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("latest") != "true" {
			t.Errorf("expected latest=true, got %q", q.Get("latest"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","start_date":"2022-06-30","end_date":"2022-07-26","prospectus_type":"Monthly","first_sell_start_date":"14/01/2552","first_sell_end_date":"20/01/2552","last_upd_date":"2022-07-26T07:53:25Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	ipos, _, err := c.GetFundIPOs(ctx, FactsheetOptions{ProjID: "M0000_2552", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ipos) != 1 {
		t.Fatalf("expected 1 IPO record, got %d", len(ipos))
	}
	if ipos[0].FirstSellStartDate != "14/01/2552" {
		t.Errorf("unexpected first_sell_start_date: %s", ipos[0].FirstSellStartDate)
	}
	if ipos[0].FirstSellEndDate != "20/01/2552" {
		t.Errorf("unexpected first_sell_end_date: %s", ipos[0].FirstSellEndDate)
	}
}

func TestGetFactsheetDividendPolicy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/factsheet/dividend-policy" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("proj_id") != "M0000_2552" {
			t.Errorf("unexpected proj_id: %q", q.Get("proj_id"))
		}
		if q.Get("latest") != "true" {
			t.Errorf("expected latest=true, got %q", q.Get("latest"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"M0000_2552","fund_class_name":"main","start_date":"2022-06-30","end_date":"2022-07-26","prospectus_type":"Monthly","dividend_policy":"N","last_upd_date":"2022-07-26T07:53:25Z"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	policies, _, err := c.GetFactsheetDividendPolicy(ctx, FactsheetOptions{ProjID: "M0000_2552", Latest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(policies) != 1 {
		t.Fatalf("expected 1 dividend policy, got %d", len(policies))
	}
	if policies[0].DividendPolicy != "N" {
		t.Errorf("unexpected dividend_policy: %s", policies[0].DividendPolicy)
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
