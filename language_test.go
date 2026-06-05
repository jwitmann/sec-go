package sec

import (
	"testing"
)

func TestNormalizeLanguage(t *testing.T) {
	tests := []struct {
		input    Language
		expected Language
	}{
		{"th", LanguageThai},
		{"TH", LanguageThai},
		{"tha", LanguageThai},
		{"en", LanguageEnglish},
		{"EN", LanguageEnglish},
		{"eng", LanguageEnglish},
		{"", LanguageEnglish},
		{"fr", LanguageEnglish},
	}

	for _, tc := range tests {
		got := normalizeLanguage(tc.input)
		if got != tc.expected {
			t.Errorf("normalizeLanguage(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestAMCName(t *testing.T) {
	a := AMC{CompNameTH: "บลจ. กรุงศรี", CompNameEN: "Krungthai Asset Management"}

	if got := a.Name(LanguageEnglish); got != "Krungthai Asset Management" {
		t.Errorf("English name = %q, want %q", got, "Krungthai Asset Management")
	}
	if got := a.Name(LanguageThai); got != "บลจ. กรุงศรี" {
		t.Errorf("Thai name = %q, want %q", got, "บลจ. กรุงศรี")
	}
}

func TestAMCNameFallback(t *testing.T) {
	a := AMC{CompNameTH: "", CompNameEN: "English Only"}
	if got := a.Name(LanguageThai); got != "English Only" {
		t.Errorf("Thai fallback = %q, want %q", got, "English Only")
	}

	a = AMC{CompNameTH: "ไทยอย่างเดียว", CompNameEN: ""}
	if got := a.Name(LanguageEnglish); got != "ไทยอย่างเดียว" {
		t.Errorf("English fallback = %q, want %q", got, "ไทยอย่างเดียว")
	}
}

func TestFundProfileName(t *testing.T) {
	p := FundProfile{ProjNameTH: "กองทุนไทย", ProjNameEN: "Thai Fund"}

	if got := p.Name(LanguageEnglish); got != "Thai Fund" {
		t.Errorf("English name = %q, want %q", got, "Thai Fund")
	}
	if got := p.Name(LanguageThai); got != "กองทุนไทย" {
		t.Errorf("Thai name = %q, want %q", got, "กองทุนไทย")
	}
}

func TestFundProfileCompanyName(t *testing.T) {
	p := FundProfile{CompNameTH: "บลจ. กสิกร", CompNameEN: "Kasikorn Asset Management"}

	if got := p.CompanyName(LanguageEnglish); got != "Kasikorn Asset Management" {
		t.Errorf("English company = %q, want %q", got, "Kasikorn Asset Management")
	}
}

func TestFundInvolvePartyEntityName(t *testing.T) {
	f := FundInvolveParty{EntityNameTH: "บริษัท ทดสอบ", EntityNameEN: "Test Company"}

	if got := f.EntityName(LanguageThai); got != "บริษัท ทดสอบ" {
		t.Errorf("Thai entity = %q, want %q", got, "บริษัท ทดสอบ")
	}
}

func TestTranslateFee(t *testing.T) {
	fee := MutualFundFee{FeeTypeDesc: "ค่าธรรมเนียมการจัดการ", RateUnit: "% ต่อปี"}

	TranslateFee(&fee, false)
	if fee.FeeTypeDesc != "ค่าธรรมเนียมการจัดการ" {
		t.Errorf("useEnglish=false should not mutate, got %q", fee.FeeTypeDesc)
	}

	TranslateFee(&fee, true)
	if fee.FeeTypeDesc != "Management Fee" {
		t.Errorf("fee type = %q, want %q", fee.FeeTypeDesc, "Management Fee")
	}
	if fee.RateUnit != "% per year" {
		t.Errorf("rate unit = %q, want %q", fee.RateUnit, "% per year")
	}
}

func TestTranslateFeeUnknown(t *testing.T) {
	fee := MutualFundFee{FeeTypeDesc: "ค่าธรรมเนียมที่ไม่รู้จัก", RateUnit: "หน่วยที่ไม่รู้จัก"}
	original := fee.FeeTypeDesc

	TranslateFee(&fee, true)
	if fee.FeeTypeDesc != original {
		t.Errorf("unknown fee should not be mutated, got %q", fee.FeeTypeDesc)
	}
}

func TestTranslateFactsheetFee(t *testing.T) {
	fee := FactsheetFee{FeeTypeDesc: "ค่าธรรมเนียมผู้ดูแลผลประโยชน์"}
	TranslateFactsheetFee(&fee, true)
	if fee.FeeTypeDesc != "Trustee Fee" {
		t.Errorf("factsheet fee = %q, want %q", fee.FeeTypeDesc, "Trustee Fee")
	}
}

func TestTranslateAssetAllocation(t *testing.T) {
	alloc := AssetAllocation{AssetName: "หน่วยลงทุน"}
	TranslateAssetAllocation(&alloc, true)
	if alloc.AssetName != "Investment Units" {
		t.Errorf("asset name = %q, want %q", alloc.AssetName, "Investment Units")
	}
}

func TestTranslateTop5Holding(t *testing.T) {
	holding := Top5Holding{AssetName: "หุ้นสามัญ"}
	TranslateTop5Holding(&holding, true)
	if holding.AssetName != "Common Stocks" {
		t.Errorf("asset name = %q, want %q", holding.AssetName, "Common Stocks")
	}
}

func TestTranslateQuarterlyPortfolio(t *testing.T) {
	item := QuarterlyPortfolio{AssetliabDesc: "ตั๋วเงินคลัง (Treasury Bill)"}
	TranslateQuarterlyPortfolio(&item, true)
	if item.AssetliabDesc != "Treasury Bill" {
		t.Errorf("assetliab desc = %q, want %q", item.AssetliabDesc, "Treasury Bill")
	}
}

func TestTranslateMonthlyPortfolioAssetType(t *testing.T) {
	item := MonthlyPortfolioAssetType{AssetliabDesc: "พันธบัตรรัฐบาล"}
	TranslateMonthlyPortfolioAssetType(&item, true)
	if item.AssetliabDesc != "Government Bond" {
		t.Errorf("assetliab desc = %q, want %q", item.AssetliabDesc, "Government Bond")
	}
}

func TestTranslateAllFees(t *testing.T) {
	fees := []MutualFundFee{
		{FeeTypeDesc: "ค่าธรรมเนียมการจัดการ"},
		{FeeTypeDesc: "ค่าธรรมเนียมผู้ดูแลผลประโยชน์"},
	}
	TranslateAllFees(fees, true)

	if fees[0].FeeTypeDesc != "Management Fee" {
		t.Errorf("fee[0] = %q, want Management Fee", fees[0].FeeTypeDesc)
	}
	if fees[1].FeeTypeDesc != "Trustee Fee" {
		t.Errorf("fee[1] = %q, want Trustee Fee", fees[1].FeeTypeDesc)
	}
}

func TestClientLanguageOption(t *testing.T) {
	c, err := NewClient("key", WithLanguage(LanguageThai))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Language() != LanguageThai {
		t.Errorf("language = %q, want %q", c.Language(), LanguageThai)
	}

	c2, _ := NewClient("key", WithLanguage("TH"))
	if c2.Language() != LanguageThai {
		t.Errorf("normalized language = %q, want %q", c2.Language(), LanguageThai)
	}
}
