package sec

import (
	"testing"
)

func TestCheckBondSaleReps(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/bond-sale-rep", map[string]string{
		"license_no": "BR1234",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"BR1234","license_no":"BR1234","license_type":"Bond Sale Rep","entity_name_th":"นายทดสอบ","entity_name_en":"Mr. Test","license_status":"Active","issue_date":"2024-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckBondSaleReps(ctx, BondSaleRepOptions{LicenseNo: "BR1234"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].LicenseNo != "BR1234" {
		t.Errorf("unexpected license no: %s", results[0].LicenseNo)
	}
	if results[0].LicenseStatus != "Active" {
		t.Errorf("unexpected status: %s", results[0].LicenseStatus)
	}
}

func TestCheckSecuritiesCompanies(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/securities-company", map[string]string{
		"company_name": "Test Company",
	}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"SC1234","license_no":"SC1234","license_type":"Securities Company","entity_name_th":"บริษัท ทดสอบ","entity_name_en":"Test Company","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckSecuritiesCompanies(ctx, SecuritiesCompanyOptions{CompanyName: "Test Company"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].EntityNameEN != "Test Company" {
		t.Errorf("unexpected entity name: %s", results[0].EntityNameEN)
	}
}

func TestCheckDerivativesCompanies(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/derivatives-company", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"DC1234","license_no":"DC1234","license_type":"Derivatives Company","entity_name_th":"บริษัท อนุพันธ์","entity_name_en":"Derivatives Company","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckDerivativesCompanies(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].LicenseType != "Derivatives Company" {
		t.Errorf("unexpected license type: %s", results[0].LicenseType)
	}
}

func TestCheckSecuritiesBrokers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/securities-broker", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"SB1234","license_no":"SB1234","license_type":"Securities Broker","entity_name_th":"นายหน้าทดสอบ","entity_name_en":"Test Broker","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckSecuritiesBrokers(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckDerivativesBrokers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/derivatives-broker", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"DB1234","license_no":"DB1234","license_type":"Derivatives Broker","entity_name_th":"นายหน้าอนุพันธ์","entity_name_en":"Derivatives Broker","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckDerivativesBrokers(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckInvestmentAdvisors(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/investment-advisor", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"IA1234","license_no":"IA1234","license_type":"Investment Advisor","entity_name_th":"ที่ปรึกษาทดสอบ","entity_name_en":"Test Advisor","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckInvestmentAdvisors(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckSecuritiesFundManagers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/securities-fund-manager", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"SFM1234","license_no":"SFM1234","license_type":"Securities Fund Manager","entity_name_th":"ผู้จัดการกองทุนทดสอบ","entity_name_en":"Test Fund Manager","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckSecuritiesFundManagers(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckFundSupervisors(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/fund-supervisor", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"FS1234","license_no":"FS1234","license_type":"Fund Supervisor","entity_name_th":"ผู้ดูแลกองทุนทดสอบ","entity_name_en":"Test Supervisor","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckFundSupervisors(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckAuditors(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/auditor", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"AU1234","license_no":"AU1234","license_type":"Auditor","entity_name_th":"ผู้สอบบัญชีทดสอบ","entity_name_en":"Test Auditor","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckAuditors(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckCreditRatingCompanies(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/credit-rating-company", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"CR1234","license_no":"CR1234","license_type":"Credit Rating Company","entity_name_th":"บริษัท จัดอันดับความน่าเชื่อถือ","entity_name_en":"Credit Rating Company","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckCreditRatingCompanies(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckPrivateFunds(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/private-fund", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"PF1234","license_no":"PF1234","license_type":"Private Fund","entity_name_th":"กองทุนส่วนบุคคลทดสอบ","entity_name_en":"Test Private Fund","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckPrivateFunds(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckDerivativesFundManagers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/derivatives-fund-manager", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"DFM1234","license_no":"DFM1234","license_type":"Derivatives Fund Manager","entity_name_th":"ผู้จัดการกองทุนอนุพันธ์","entity_name_en":"Derivatives Fund Manager","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckDerivativesFundManagers(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckSecuritiesBorrowingLending(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/securities-borrowing-lending", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"SBL1234","license_no":"SBL1234","license_type":"Securities Borrowing & Lending","entity_name_th":"บริษัท ให้กู้หลักทรัพย์","entity_name_en":"Securities Borrowing & Lending","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckSecuritiesBorrowingLending(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckFinancialAdvisors(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/financial-advisor", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"FA1234","license_no":"FA1234","license_type":"Financial Advisor","entity_name_th":"ที่ปรึกษาทางการเงิน","entity_name_en":"Financial Advisor","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckFinancialAdvisors(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckAcquirers(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/acquirer", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"AQ1234","license_no":"AQ1234","license_type":"Acquirer","entity_name_th":"ผู้ควบรวมกิจการ","entity_name_en":"Acquirer","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckAcquirers(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestCheckVentureCapitals(t *testing.T) {
	c, ctx := newQueryServer(t, "/v1/license-check/venture-capital", map[string]string{}, `{"message":"ok","page_size":10,"next_cursor":"","items":[{"license_id":"VC1234","license_no":"VC1234","license_type":"Venture Capital","entity_name_th":"บริษัท ร่วมทุนทดสอบ","entity_name_en":"Test Venture Capital","license_status":"Active","issue_date":"2020-01-01","expire_date":"2025-01-01","license_detail":"","remark":"","last_upd_date":"2024-01-15T00:00:00Z"}]}`)

	results, _, err := c.CheckVentureCapitals(ctx, LicenseCheckOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}
