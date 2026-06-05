package sec

// Language represents the preferred display language for bilingual fields.
type Language string

const (
	LanguageThai    Language = "th"
	LanguageEnglish Language = "en"
)

func normalizeLanguage(lang Language) Language {
	switch lang {
	case LanguageThai, "TH", "tha":
		return LanguageThai
	case LanguageEnglish, "EN", "eng":
		return LanguageEnglish
	default:
		return LanguageEnglish
	}
}

// Name returns the company name in the client's preferred language.
// Falls back to the other language if the preferred one is empty.
func (a AMC) Name(lang Language) string {
	return pickString(lang, a.CompNameTH, a.CompNameEN)
}

// Name returns the fund name in the client's preferred language.
// Falls back to the other language if the preferred one is empty.
func (p FundProfile) Name(lang Language) string {
	return pickString(lang, p.ProjNameTH, p.ProjNameEN)
}

// CompanyName returns the AMC name in the client's preferred language.
func (p FundProfile) CompanyName(lang Language) string {
	return pickString(lang, p.CompNameTH, p.CompNameEN)
}

// EntityName returns the entity name in the client's preferred language.
func (f FundInvolveParty) EntityName(lang Language) string {
	return pickString(lang, f.EntityNameTH, f.EntityNameEN)
}

func pickString(lang Language, thai, english string) string {
	lang = normalizeLanguage(lang)
	if lang == LanguageThai {
		if thai != "" {
			return thai
		}
		return english
	}
	if english != "" {
		return english
	}
	return thai
}

// FeeTypeTranslation maps Thai fee type descriptions to English.
// Extend this map for additional fee descriptions returned by the SEC API.
var FeeTypeTranslation = map[string]string{
	"ค่าธรรมเนียมการจัดการ":                                  "Management Fee",
	"ค่าธรรมเนียมผู้ดูแลผลประโยชน์":                          "Trustee Fee",
	"ค่าธรรมเนียมนายทะเบียนหน่วย":                            "Registrar Fee",
	"ค่าธรรมเนียมการขายหน่วยลงทุน (Front-end Fee)":           "Front-end Fee",
	"ค่าธรรมเนียมการรับซื้อคืนหน่วยลงทุน (Back-end Fee)":     "Back-end Fee",
	"ค่าธรรมเนียมการสับเปลี่ยนหน่วยลงทุนเข้า (SWITCHING IN)": "Switch-in Fee",
	"ค่าธรรมเนียมการสับเปลี่ยนหน่วยลงทุนออก (SWITCHING OUT)": "Switch-out Fee",
	"ค่าธรรมเนียมการโอนหน่วยลงทุน":                           "Transfer Fee",
	"ค่าใช้จ่ายอื่นๆ":                                        "Other Fee",
	"ค่าธรรมเนียมและค่าใช้จ่ายรวมทั้งหมด":                    "Total Expense Ratio",
}

// FeeUnitTranslation maps Thai fee unit descriptions to English.
var FeeUnitTranslation = map[string]string{
	"ต่อปี ของมูลค่าทรัพย์สินสุทธิของกองทุน": "per year of NAV",
	"% ต่อปี": "% per year",
	"บาท":     "baht",
}

// AssetNameTranslation maps common Thai asset type names to English.
var AssetNameTranslation = map[string]string{
	"หน่วยลงทุน":             "Investment Units",
	"เงินฝากธนาคาร และอื่นๆ": "Bank Deposits and Others",
	"ตราสารหนี้":             "Debt Instruments",
	"หุ้นสามัญ":              "Common Stocks",
	"ทรัพย์สินอื่นๆ":         "Other Assets",
	"เงินสด":                 "Cash",
}

// AssetLiabilityTranslation maps common Thai asset/liability category
// descriptions to English.
var AssetLiabilityTranslation = map[string]string{
	"ตั๋วเงินคลัง (Treasury Bill)": "Treasury Bill",
	"ตั๋วเงินคลัง":                 "Treasury Bill",
	"พันธบัตรรัฐบาล":               "Government Bond",
	"เงินฝากธนาคาร":                "Bank Deposit",
	"หุ้นสามัญ":                    "Common Stock",
}

// TranslateFee translates Thai fee descriptions on a MutualFundFee to English
// when useEnglish is true. It mutates the provided fee in place.
func TranslateFee(fee *MutualFundFee, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := FeeTypeTranslation[fee.FeeTypeDesc]; ok {
		fee.FeeTypeDesc = trans
	}
	if trans, ok := FeeUnitTranslation[fee.RateUnit]; ok {
		fee.RateUnit = trans
	}
}

// TranslateFactsheetFee translates Thai fee descriptions on a FactsheetFee to
// English when useEnglish is true. It mutates the provided fee in place.
func TranslateFactsheetFee(fee *FactsheetFee, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := FeeTypeTranslation[fee.FeeTypeDesc]; ok {
		fee.FeeTypeDesc = trans
	}
}

// TranslateAssetAllocation translates the Thai asset name on an AssetAllocation
// to English when useEnglish is true. It mutates the provided allocation in place.
func TranslateAssetAllocation(alloc *AssetAllocation, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := AssetNameTranslation[alloc.AssetName]; ok {
		alloc.AssetName = trans
	}
}

// TranslateTop5Holding translates the Thai asset name on a Top5Holding to
// English when useEnglish is true. It mutates the provided holding in place.
func TranslateTop5Holding(holding *Top5Holding, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := AssetNameTranslation[holding.AssetName]; ok {
		holding.AssetName = trans
	}
}

// TranslateQuarterlyPortfolio translates Thai asset/liability descriptions on
// a QuarterlyPortfolio to English when useEnglish is true.
func TranslateQuarterlyPortfolio(item *QuarterlyPortfolio, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := AssetLiabilityTranslation[item.AssetliabDesc]; ok {
		item.AssetliabDesc = trans
	}
}

// TranslateMonthlyPortfolioAssetType translates Thai asset/liability
// descriptions on a MonthlyPortfolioAssetType to English when useEnglish is true.
func TranslateMonthlyPortfolioAssetType(item *MonthlyPortfolioAssetType, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := AssetLiabilityTranslation[item.AssetliabDesc]; ok {
		item.AssetliabDesc = trans
	}
}

// TranslateAllFees translates every fee in the slice when useEnglish is true.
func TranslateAllFees(fees []MutualFundFee, useEnglish bool) {
	for i := range fees {
		TranslateFee(&fees[i], useEnglish)
	}
}

// TranslateAllFactsheetFees translates every factsheet fee in the slice when
// useEnglish is true.
func TranslateAllFactsheetFees(fees []FactsheetFee, useEnglish bool) {
	for i := range fees {
		TranslateFactsheetFee(&fees[i], useEnglish)
	}
}

// TranslateAllAssetAllocations translates every allocation in the slice when
// useEnglish is true.
func TranslateAllAssetAllocations(allocs []AssetAllocation, useEnglish bool) {
	for i := range allocs {
		TranslateAssetAllocation(&allocs[i], useEnglish)
	}
}

// TranslateAllTop5Holdings translates every holding in the slice when
// useEnglish is true.
func TranslateAllTop5Holdings(holdings []Top5Holding, useEnglish bool) {
	for i := range holdings {
		TranslateTop5Holding(&holdings[i], useEnglish)
	}
}

// TranslateAllQuarterlyPortfolios translates every portfolio item in the slice
// when useEnglish is true.
func TranslateAllQuarterlyPortfolios(items []QuarterlyPortfolio, useEnglish bool) {
	for i := range items {
		TranslateQuarterlyPortfolio(&items[i], useEnglish)
	}
}

// TranslateAllMonthlyPortfolioAssetTypes translates every portfolio item in
// the slice when useEnglish is true.
func TranslateAllMonthlyPortfolioAssetTypes(items []MonthlyPortfolioAssetType, useEnglish bool) {
	for i := range items {
		TranslateMonthlyPortfolioAssetType(&items[i], useEnglish)
	}
}
