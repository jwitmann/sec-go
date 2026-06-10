package sec

import (
	"net/http"
	"testing"
)

func TestListBondIssuers(t *testing.T) {
	c, ctx := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/bond/general-info/amcs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","page_size":10,"next_cursor":"","items":[{"unique_id":"C0000000021","comp_name_th":"บลจ. กรุงศรี","comp_name_en":"Krungthai Asset Management","last_upd_date":"2024-01-15T00:00:00Z"}]}`))
	})

	issuers, cursor, err := c.ListBondIssuers(ctx, BondIssuerOptions{PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(issuers) != 1 {
		t.Fatalf("expected 1 issuer, got %d", len(issuers))
	}
	if issuers[0].CompNameEN != "Krungthai Asset Management" {
		t.Errorf("unexpected name: %s", issuers[0].CompNameEN)
	}
	if cursor != "" {
		t.Errorf("expected empty cursor, got %q", cursor)
	}
}

func TestGetBondFeatures(t *testing.T) {
	c, ctx := newQueryServer(t, "/v2/bond/general-info/features", map[string]string{
		"proj_id": "B1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"B1234_2567","proj_name_th":"บอนด์ทดสอบ","proj_name_en":"Test Bond","proj_abbr_name":"TEST","issue_date":"2024-01-15","maturity_date":"2029-01-15","coupon_rate":3.5,"face_value":1000,"issue_value":1000000000,"bond_type":"Corporate","bond_sub_type":"Senior","secured_flag":"Y","guarantee_flag":"N","collateral_desc":"","purpose_th":"","purpose_en":"","remark_th":"","remark_en":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	features, _, err := c.GetBondFeatures(ctx, BondFeatureOptions{ProjID: "B1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(features))
	}
	if features[0].ProjID != "B1234_2567" {
		t.Errorf("unexpected proj_id: %s", features[0].ProjID)
	}
	if features[0].CouponRate != 3.5 {
		t.Errorf("unexpected coupon rate: %f", features[0].CouponRate)
	}
}

func TestGetBondCreditRatings(t *testing.T) {
	c, ctx := newQueryServer(t, "/v2/bond/credit-rating", map[string]string{
		"proj_id":      "B1234_2567",
		"start_period": "202401",
		"end_period":   "202403",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"B1234_2567","period":202401,"rating_agency":"TRIS","rating":"AAA","outlook":"Stable","as_of_date":"2024-01-31","last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	ratings, _, err := c.GetBondCreditRatings(ctx, BondPeriodOptions{
		ProjID:      "B1234_2567",
		StartPeriod: 202401,
		EndPeriod:   202403,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ratings) != 1 {
		t.Fatalf("expected 1 rating, got %d", len(ratings))
	}
	if ratings[0].Rating != "AAA" {
		t.Errorf("unexpected rating: %s", ratings[0].Rating)
	}
	if ratings[0].RatingAgency != "TRIS" {
		t.Errorf("unexpected agency: %s", ratings[0].RatingAgency)
	}
}

func TestGetBondOutstanding(t *testing.T) {
	c, ctx := newQueryServer(t, "/v2/bond/outstanding", map[string]string{
		"proj_id":      "B1234_2567",
		"start_period": "202401",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"B1234_2567","period":202401,"outstanding_qty":1000000,"outstanding_val":1000000000,"held_by_investor":950000,"held_by_issuer":50000,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	outstanding, _, err := c.GetBondOutstanding(ctx, BondPeriodOptions{
		ProjID:      "B1234_2567",
		StartPeriod: 202401,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(outstanding) != 1 {
		t.Fatalf("expected 1 outstanding record, got %d", len(outstanding))
	}
	if outstanding[0].OutstandingVal != 1000000000 {
		t.Errorf("unexpected outstanding val: %f", outstanding[0].OutstandingVal)
	}
}

func TestGetBondRelatedParties(t *testing.T) {
	c, ctx := newQueryServer(t, "/v2/bond/related-party", map[string]string{
		"proj_id": "B1234_2567",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"B1234_2567","period":202401,"party_type":"Issuer","party_name_th":"บริษัท ทดสอบ จำกัด","party_name_en":"Test Company Ltd.","last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	parties, _, err := c.GetBondRelatedParties(ctx, BondPeriodOptions{ProjID: "B1234_2567"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(parties) != 1 {
		t.Fatalf("expected 1 party, got %d", len(parties))
	}
	if parties[0].PartyType != "Issuer" {
		t.Errorf("unexpected party type: %s", parties[0].PartyType)
	}
}

func TestGetBondInvestorHoldings(t *testing.T) {
	c, ctx := newQueryServer(t, "/v2/bond/investor-holding", map[string]string{
		"proj_id":      "B1234_2567",
		"start_period": "202401",
		"end_period":   "202401",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"proj_id":"B1234_2567","period":202401,"investor_type":"Institutional","holding_qty":500000,"holding_val":500000000,"holding_pct":50,"last_upd_date":"2024-02-01T00:00:00Z"}]}`)

	holdings, _, err := c.GetBondInvestorHoldings(ctx, BondPeriodOptions{
		ProjID:      "B1234_2567",
		StartPeriod: 202401,
		EndPeriod:   202401,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(holdings) != 1 {
		t.Fatalf("expected 1 holding, got %d", len(holdings))
	}
	if holdings[0].InvestorType != "Institutional" {
		t.Errorf("unexpected investor type: %s", holdings[0].InvestorType)
	}
	if holdings[0].HoldingPct != 50 {
		t.Errorf("unexpected holding pct: %f", holdings[0].HoldingPct)
	}
}
