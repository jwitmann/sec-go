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

func TestTranslateThaiPeriod(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3 ปี 3 เดือน", "3 years 3 months"},
		{"1 เดือน 13 วัน", "1 month 13 days"},
		{"6 เดือน", "6 months"},
		{"15 วัน", "15 days"},
		{"1 ปี", "1 year"},
		{"1 เดือน", "1 month"},
		{"1 วัน", "1 day"},
		{"ไม่มี", "None"},
		{"-", "-"},
		{"", ""},
		{"3 years 3 months", "3 years 3 months"}, // already English
		{"1 month 13 days", "1 month 13 days"},   // already English
	}

	for _, tc := range tests {
		got := TranslateThaiPeriod(tc.input)
		if got != tc.expected {
			t.Errorf("TranslateThaiPeriod(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestConvertBuddhistDate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"03/04/2557", "2014-04-03"},
		{"15/05/2557", "2014-05-15"},
		{"01/01/2566", "2023-01-01"},
		{"31/12/2566", "2023-12-31"},
		{"2023-01-01", "2023-01-01"}, // already ISO
		{"-", "-"},
		{"", ""},
	}

	for _, tc := range tests {
		got := ConvertBuddhistDate(tc.input)
		if got != tc.expected {
			t.Errorf("ConvertBuddhistDate(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestTranslateFactsheetStatistics(t *testing.T) {
	stats := FactsheetStatistics{
		RecoveringPeriod:        "3 ปี 3 เดือน",
		PortfolioDurationPeriod: "1 เดือน 13 วัน",
		PortfolioTurnoverRatio:  "24.63",
	}

	TranslateFactsheetStatistics(&stats, false)
	if stats.RecoveringPeriod != "3 ปี 3 เดือน" {
		t.Errorf("useEnglish=false should not mutate, got %q", stats.RecoveringPeriod)
	}

	TranslateFactsheetStatistics(&stats, true)
	if stats.RecoveringPeriod != "3 years 3 months" {
		t.Errorf("recovering period = %q, want %q", stats.RecoveringPeriod, "3 years 3 months")
	}
	if stats.PortfolioDurationPeriod != "1 month 13 days" {
		t.Errorf("portfolio duration = %q, want %q", stats.PortfolioDurationPeriod, "1 month 13 days")
	}
	if stats.PortfolioTurnoverRatio != "24.63" {
		t.Errorf("turnover ratio should not change, got %q", stats.PortfolioTurnoverRatio)
	}
}

func TestTranslateFundIPO(t *testing.T) {
	ipo := FundIPO{
		FirstSellStartDate: "03/04/2557",
		FirstSellEndDate:   "15/05/2557",
	}

	TranslateFundIPO(&ipo, false)
	if ipo.FirstSellStartDate != "03/04/2557" {
		t.Errorf("useEnglish=false should not mutate, got %q", ipo.FirstSellStartDate)
	}

	TranslateFundIPO(&ipo, true)
	if ipo.FirstSellStartDate != "2014-04-03" {
		t.Errorf("first sell start = %q, want %q", ipo.FirstSellStartDate, "2014-04-03")
	}
	if ipo.FirstSellEndDate != "2014-05-15" {
		t.Errorf("first sell end = %q, want %q", ipo.FirstSellEndDate, "2014-05-15")
	}
}

func TestTranslateBenchmarkName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ดัชนี Hang Seng", "Hang Seng Index"},
		{"ดัชนี MSCI China TRN", "MSCI China TRN Index"},
		{"ผลการดำเนินงานของกองทุนรวมหลัก", "Performance of the Master Fund"},
		{"ดัชนีผลตอบแทนรวม SET 50", "SET 50 Total Return Index"},
		{"", ""},
		{"Already English", "Already English"},
	}

	for _, tc := range tests {
		got := TranslateBenchmarkName(tc.input)
		if got != tc.expected {
			t.Errorf("TranslateBenchmarkName(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestTranslateBenchmarkRemark(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"ปรับด้วยอัตราแลกเปลี่ยน เพื่อคำนวณผลตอบแทนเป็นสกุลเงินบาท ณ วันที่คำนวณผลตอบแทน",
			"Adjusted with exchange rate to calculate returns in Thai Baht as of the return calculation date",
		},
		{
			"ปรับด้วยต้นทุนการป้องกันความเสี่ยงด้านอัตราแลกเปลี่ยน เพื่อคำนวณผลตอบแทนเป็นสกุลเงินบาท ณ วันที่คำนวณผลตอบแทน",
			"Adjusted with foreign exchange hedging cost to calculate returns in Thai Baht as of the return calculation date",
		},
		{
			"ปรับด้วยต้นทุนการป้องกันความเสี่ยง 90% และปรับด้วยอัตราแลกเปลี่ยน 10%",
			"Adjusted with 90% hedging cost and 10% exchange rate",
		},
		{"", ""},
		{"Already English", "Already English"},
	}

	for _, tc := range tests {
		got := TranslateBenchmarkRemark(tc.input)
		if got != tc.expected {
			t.Errorf("TranslateBenchmarkRemark(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestTranslateFactsheetBenchmark(t *testing.T) {
	bench := FactsheetBenchmark{
		Benchmark: "ดัชนี Hang Seng",
		Remark:    "ปรับด้วยต้นทุนการป้องกันความเสี่ยง 90% และปรับด้วยอัตราแลกเปลี่ยน 10%",
	}

	TranslateFactsheetBenchmark(&bench, false)
	if bench.Benchmark != "ดัชนี Hang Seng" {
		t.Errorf("useEnglish=false should not mutate benchmark, got %q", bench.Benchmark)
	}

	TranslateFactsheetBenchmark(&bench, true)
	if bench.Benchmark != "Hang Seng Index" {
		t.Errorf("benchmark = %q, want %q", bench.Benchmark, "Hang Seng Index")
	}
	if bench.Remark != "Adjusted with 90% hedging cost and 10% exchange rate" {
		t.Errorf("remark = %q, want %q", bench.Remark, "Adjusted with 90% hedging cost and 10% exchange rate")
	}
}
