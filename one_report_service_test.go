package sec

import (
	"encoding/json"
	"testing"
)

func TestOneReportStructs(t *testing.T) {
	t.Run("SBOInfo", func(t *testing.T) {
		data := []byte(`[{
			"last_upd_date": "2023-10-26T04:50:50.0000000+00:00",
			"report_year": "2021",
			"unique_id": "C0000000013",
			"language": "T",
			"corp_name": "ธนาคารกรุงไทย จำกัด (มหาชน)",
			"symbol": "KTB",
			"address": "35 ถนนสุขุมวิท",
			"province": "กรุงเทพ",
			"zip_code": "10110",
			"business_type": "ธุรกิจธนาคาร",
			"registered_number": "0107537000882",
			"telephone": "0-2111-1111",
			"fax": "0-2255-9391",
			"website": "https://krungthai.com",
			"email": "test@example.com",
			"common_paidup_share": 13976061250,
			"preferred_paidup_share": 5500000
		}]`)

		var items []SBOInfo
		if err := json.Unmarshal(data, &items); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}
		item := items[0]
		if item.Symbol != "KTB" {
			t.Errorf("symbol: expected KTB, got %s", item.Symbol)
		}
		if item.ReportYear != "2021" {
			t.Errorf("report_year: expected 2021, got %s", item.ReportYear)
		}
		if item.CommonPaidupShare != 13976061250 {
			t.Errorf("common_paidup_share: expected 13976061250, got %f", item.CommonPaidupShare)
		}
	})

	t.Run("SBOProductIncome", func(t *testing.T) {
		data := []byte(`[{
			"last_upd_date": "2023-10-26T04:50:48.0000000+00:00",
			"report_year": "2021",
			"unique_id": "C0000000013",
			"business_income_code": "218",
			"sequence": 1,
			"business_income_desc": "รายได้รวม",
			"asof_year": 115785741,
			"asof_yesteryear": 122247423,
			"asof_year_before_yesteryear": 125657639,
			"asof_year_percent": 100,
			"asof_yesteryear_percent": 100,
			"asof_year_before_yesteryear_percent": 100
		}]`)

		var items []SBOProductIncome
		if err := json.Unmarshal(data, &items); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}
		item := items[0]
		if item.Sequence != 1 {
			t.Errorf("sequence: expected 1, got %d", item.Sequence)
		}
		if item.AsofYearPercent != 100 {
			t.Errorf("asof_year_percent: expected 100, got %f", item.AsofYearPercent)
		}
	})

	t.Run("SBORisk", func(t *testing.T) {
		data := []byte(`[{
			"last_upd_date": "2023-10-26T04:50:48.0000000+00:00",
			"report_year": "2021",
			"unique_id": "C0000000013",
			"risk_category": "01",
			"risk_code": "R001",
			"sequence": 1,
			"choice": "Y",
			"holder_risk": "N",
			"foreign_risk": "D"
		}]`)

		var items []SBORisk
		if err := json.Unmarshal(data, &items); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}
		item := items[0]
		if item.Choice != "Y" {
			t.Errorf("choice: expected Y, got %s", item.Choice)
		}
		if item.ForeignRisk != "D" {
			t.Errorf("foreign_risk: expected D, got %s", item.ForeignRisk)
		}
	})
}

func TestDigitalAssetStructs(t *testing.T) {
	t.Run("DigitalAssetIntermediary", func(t *testing.T) {
		data := []byte(`[{
			"unique_id": "DA001",
			"name_th": "Test ดิจิทัล",
			"name_en": "Test Digital",
			"lic_code": "L001",
			"lic_action_code": "A",
			"lic_efft_date": "2024-01-01",
			"lic_act_date": "2024-01-15",
			"lic_exp_date": "2025-01-01"
		}]`)

		var items []DigitalAssetIntermediary
		if err := json.Unmarshal(data, &items); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}
		item := items[0]
		if item.UniqueID != "DA001" {
			t.Errorf("unique_id: expected DA001, got %s", item.UniqueID)
		}
		if item.NameEN != "Test Digital" {
			t.Errorf("name_en: expected Test Digital, got %s", item.NameEN)
		}
	})

	t.Run("DigitalAssetIntermediaryRequest", func(t *testing.T) {
		req := DigitalAssetIntermediaryRequest{IntermediaryName: "Test"}
		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		expected := `{"IntermediaryName":"Test"}`
		if string(data) != expected {
			t.Errorf("expected %s, got %s", expected, string(data))
		}
	})
}
