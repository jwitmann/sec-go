package sec

import (
	"net/http"
	"testing"
)

func TestListPVDs(t *testing.T) {
	c, ctx := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/pvd/general-info/list" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":10,"next_cursor":"","items":[{"unique_id":"C0000000021","comp_name_th":"บริษัท ทดสอบ","comp_name_en":"Test Company","last_upd_date":"2024-01-15T00:00:00Z"}]}`))
	})

	pvds, cursor, err := c.ListPVDs(ctx, PVDListOptions{PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pvds) != 1 {
		t.Fatalf("expected 1 PVD, got %d", len(pvds))
	}
	if pvds[0].CompNameEN != "Test Company" {
		t.Errorf("unexpected name: %s", pvds[0].CompNameEN)
	}
	if cursor != "" {
		t.Errorf("expected empty cursor, got %q", cursor)
	}
}

func TestGetPVDFundInfo(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/general-info/fund-info", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","proj_name_th":"กองทุนสำรองเลี้ยงชีพ ทดสอบ","proj_name_en":"Test Provident Fund","proj_abbr_name":"TEST","fund_status":"Registered","init_date":"2020-01-01","regis_date":"2020-01-15","cancel_date":"","policy_desc":"Mixed","management_style":"AM","fund_class_name":"A","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	infos, _, err := c.GetPVDFundInfo(ctx, PVDProjOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(infos) != 1 {
		t.Fatalf("expected 1 info, got %d", len(infos))
	}
	if infos[0].ProjID != "PVD1234_2567" {
		t.Errorf("unexpected proj_id: %s", infos[0].ProjID)
	}
	if infos[0].FundStatus != "Registered" {
		t.Errorf("unexpected status: %s", infos[0].FundStatus)
	}
}

func TestGetPVDFundSpecs(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/general-info/fund-spec", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","spec_code":"EQ","spec_desc":"Equity Fund","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	specs, _, err := c.GetPVDFundSpecs(ctx, PVDProjOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(specs) != 1 {
		t.Fatalf("expected 1 spec, got %d", len(specs))
	}
	if specs[0].SpecCode != "EQ" {
		t.Errorf("unexpected spec code: %s", specs[0].SpecCode)
	}
}

func TestGetPVDFundMembers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-member", map[string]string{
		"proj_id":      "PVD1234_2567",
		"start_period": "202401",
		"end_period":   "202403",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"total_members":1000,"active_members":950,"retired_members":50,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	members, _, err := c.GetPVDFundMembers(ctx, PVDPeriodOptions{
		ProjID:      "PVD1234_2567",
		StartPeriod: 202401,
		EndPeriod:   202403,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(members) != 1 {
		t.Fatalf("expected 1 member record, got %d", len(members))
	}
	if members[0].TotalMembers != 1000 {
		t.Errorf("unexpected total members: %d", members[0].TotalMembers)
	}
}

func TestGetPVDFundAssets(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-asset", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"asset_seq":1,"asset_name":"Equity","asset_ratio":45.5,"market_value":455000000,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	assets, _, err := c.GetPVDFundAssets(ctx, PVDPeriodOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(assets) != 1 {
		t.Fatalf("expected 1 asset, got %d", len(assets))
	}
	if assets[0].AssetName != "Equity" {
		t.Errorf("unexpected asset name: %s", assets[0].AssetName)
	}
}

func TestGetPVDFundTransactions(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-transaction", map[string]string{
		"proj_id":      "PVD1234_2567",
		"start_period": "202401",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"trans_type":"Subscription","trans_value":10000000,"trans_count":100,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	transactions, _, err := c.GetPVDFundTransactions(ctx, PVDPeriodOptions{
		ProjID:      "PVD1234_2567",
		StartPeriod: 202401,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(transactions) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(transactions))
	}
	if transactions[0].TransType != "Subscription" {
		t.Errorf("unexpected trans type: %s", transactions[0].TransType)
	}
}

func TestGetPVDFundContributions(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-contribution", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"employee_contrib":5000000,"employer_contrib":5000000,"total_contrib":10000000,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	contributions, _, err := c.GetPVDFundContributions(ctx, PVDPeriodOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(contributions) != 1 {
		t.Fatalf("expected 1 contribution, got %d", len(contributions))
	}
	if contributions[0].TotalContrib != 10000000 {
		t.Errorf("unexpected total contrib: %f", contributions[0].TotalContrib)
	}
}

func TestGetPVDFundExpenses(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-expense", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"expense_type":"Management Fee","expense_value":1000000,"expense_ratio":0.1,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	expenses, _, err := c.GetPVDFundExpenses(ctx, PVDPeriodOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(expenses) != 1 {
		t.Fatalf("expected 1 expense, got %d", len(expenses))
	}
	if expenses[0].ExpenseType != "Management Fee" {
		t.Errorf("unexpected expense type: %s", expenses[0].ExpenseType)
	}
}

func TestGetPVDFundLiquidity(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-liquidity", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"cash_ratio":15.5,"current_ratio":1.2,"quick_ratio":1.1,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	liquidity, _, err := c.GetPVDFundLiquidity(ctx, PVDPeriodOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(liquidity) != 1 {
		t.Fatalf("expected 1 liquidity record, got %d", len(liquidity))
	}
	if liquidity[0].CashRatio != 15.5 {
		t.Errorf("unexpected cash ratio: %f", liquidity[0].CashRatio)
	}
}

func TestGetPVDFundPerformance(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-performance", map[string]string{
		"proj_id":      "PVD1234_2567",
		"start_period": "202401",
		"end_period":   "202403",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","period":202401,"performance_type":"Return","reference_period":"1Y","performance_value":8.5,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	performance, _, err := c.GetPVDFundPerformance(ctx, PVDPeriodOptions{
		ProjID:      "PVD1234_2567",
		StartPeriod: 202401,
		EndPeriod:   202403,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(performance) != 1 {
		t.Fatalf("expected 1 performance record, got %d", len(performance))
	}
	if performance[0].PerformanceValue != 8.5 {
		t.Errorf("unexpected performance value: %f", performance[0].PerformanceValue)
	}
}

func TestGetPVDFundBenchmarks(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-benchmark", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","period":202401,"benchmark":"SET Index","benchmark_value":5.2,"fund_value":8.5,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	benchmarks, _, err := c.GetPVDFundBenchmarks(ctx, PVDPeriodOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(benchmarks) != 1 {
		t.Fatalf("expected 1 benchmark, got %d", len(benchmarks))
	}
	if benchmarks[0].Benchmark != "SET Index" {
		t.Errorf("unexpected benchmark: %s", benchmarks[0].Benchmark)
	}
}

func TestGetPVDFundDividends(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-dividend", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","dividend_date":"2024-01-31","dividend_value":0.5,"book_close_date":"2024-01-25","last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	dividends, _, err := c.GetPVDFundDividends(ctx, PVDProjOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dividends) != 1 {
		t.Fatalf("expected 1 dividend, got %d", len(dividends))
	}
	if dividends[0].DividendValue != 0.5 {
		t.Errorf("unexpected dividend value: %f", dividends[0].DividendValue)
	}
}

func TestGetPVDFundPolicies(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-policy", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","policy_type":"Investment","policy_desc":"Mixed Equity-Bond","min_equity_ratio":20,"max_equity_ratio":80,"last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	policies, _, err := c.GetPVDFundPolicies(ctx, PVDProjOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(policies) != 1 {
		t.Fatalf("expected 1 policy, got %d", len(policies))
	}
	if policies[0].PolicyType != "Investment" {
		t.Errorf("unexpected policy type: %s", policies[0].PolicyType)
	}
}

func TestGetPVDFundFees(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-fee", map[string]string{
		"proj_id": "PVD1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","fund_class_name":"A","fee_type_desc":"Management Fee","rate":0.75,"rate_unit":"%","actual_value":0.75,"last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	fees, _, err := c.GetPVDFundFees(ctx, PVDProjOptions{ProjID: "PVD1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fees) != 1 {
		t.Fatalf("expected 1 fee, got %d", len(fees))
	}
	if fees[0].FeeTypeDesc != "Management Fee" {
		t.Errorf("unexpected fee type: %s", fees[0].FeeTypeDesc)
	}
}

func TestGetPVDFundCompliance(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/pvd/fund-compliance", map[string]string{
		"proj_id":      "PVD1234_2567",
		"start_period": "202401",
		"end_period":   "202403",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"PVD1234_2567","period":202401,"compliance_type":"Investment Limit","compliance_status":"Compliant","remark":"","last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	compliance, _, err := c.GetPVDFundCompliance(ctx, PVDPeriodOptions{
		ProjID:      "PVD1234_2567",
		StartPeriod: 202401,
		EndPeriod:   202403,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(compliance) != 1 {
		t.Fatalf("expected 1 compliance record, got %d", len(compliance))
	}
	if compliance[0].ComplianceStatus != "Compliant" {
		t.Errorf("unexpected compliance status: %s", compliance[0].ComplianceStatus)
	}
}
