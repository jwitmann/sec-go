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

// SubscriptionPeriodTranslation maps Thai subscription/redemption period
// descriptions to English.
var SubscriptionPeriodTranslation = map[string]string{
	"ทุกวันทำการ":                  "Every Business Day",
	"ทุกวันทำการของตลาดหลักทรัพย์": "Every Stock Exchange Business Day",
	"ทุกวันทำการของธนาคาร":         "Every Banking Day",
	"ทุกวันทำการของสถาบันการเงิน":  "Every Financial Institution Business Day",
	"วันทำการของตลาดหลักทรัพย์":    "Stock Exchange Business Day",
	"วันทำการของธนาคาร":            "Banking Day",
	"วันทำการของสถาบันการเงิน":     "Financial Institution Business Day",
	"วันทำการประจำสัปดาห์":         "Weekly Business Day",
	"ทุกวัน": "Every Day",
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
	"ตั๋วเงินคลัง (Treasury Bill)":    "Treasury Bill",
	"ตั๋วเงินคลัง":                    "Treasury Bill",
	"พันธบัตรรัฐบาล":                  "Government Bond",
	"เงินฝากธนาคาร":                   "Bank Deposit",
	"เงินฝากธนาคารประเภทกระแสรายวัน":  "Bank Deposit - Current Account",
	"เงินฝากธนาคารประเภทออมทรัพย์":    "Bank Deposit - Savings Account",
	"เงินฝากธนาคารประเภทประจำ":        "Bank Deposit - Fixed Deposit",
	"หุ้นสามัญ":                       "Common Stock",
	"หุ้นบุริมสิทธิ":                  "Preferred Stock",
	"หุ้นสามัญ ไม่ใช่บริษัทจดทะเบียน": "Common Stock - Unlisted",
	"หน่วยลงทุนในอสังหาริมทรัพย์":     "Property Fund Unit",
	"ใบสำคัญแสดงสิทธิ":                "Warrant",
	"ทรัพย์สินอื่นๆ":                  "Other Assets",
	"เงินสด": "Cash",
	"มูลค่าทรัพย์สินสุทธิ":                           "Net Asset Value",
	"รวมสินทรัพย์":                                   "Total Assets",
	"รวมหนี้สิน":                                     "Total Liabilities",
	"เงินลงทุนชั่วคราว":                              "Short-term Investment",
	"เงินให้สินเชื่อ":                                "Loan",
	"ลูกหนี้การค้า":                                  "Trade Receivable",
	"ลูกหนี้อื่น":                                    "Other Receivable",
	"สต็อกสินค้า":                                    "Inventory",
	"สินทรัพย์หมุนเวียนอื่น":                         "Other Current Assets",
	"ที่ดิน อาคาร และอุปกรณ์":                        "Property, Plant and Equipment",
	"สินทรัพย์ไม่หมุนเวียนอื่น":                      "Other Non-current Assets",
	"เงินลงทุนระยะยาว":                               "Long-term Investment",
	"สิทธิการเช่า":                                   "Leasehold Right",
	"สินทรัพย์ไม่มีตัวตน":                            "Intangible Assets",
	"ภาษีเงินได้รอการตัดบัญชี":                       "Deferred Tax Assets",
	"เงินให้กู้ระยะยาว":                              "Long-term Loan",
	"เงินกู้ระยะสั้น":                                "Short-term Loan",
	"เงินกู้ระยะยาว":                                 "Long-term Loan",
	"เจ้าหนี้การค้า":                                 "Trade Payable",
	"เจ้าหนี้อื่น":                                   "Other Payable",
	"ภาษีเงินได้ค้างจ่าย":                            "Income Tax Payable",
	"เงินกู้จากสถาบันการเงิน":                        "Loan from Financial Institution",
	"เงินกู้ระยะสั้นจากสถาบันการเงิน":                "Short-term Loan from Financial Institution",
	"เงินกู้ระยะยาวจากสถาบันการเงิน":                 "Long-term Loan from Financial Institution",
	"ตราสารหนี้":                                     "Debt Instruments",
	"ตราสารหนี้ภาครัฐ":                               "Government Debt Instruments",
	"ตราสารหนี้เอกชน":                                "Corporate Debt Instruments",
	"ตราสารหนี้ต่างประเทศ":                           "Foreign Debt Instruments",
	"ตราสารหนี้ที่มีสิทธิ":                           "Debt Instruments with Rights",
	"ตราสารหนี้อื่น":                                 "Other Debt Instruments",
	"หลักทรัพย์จดทะเบียน":                            "Listed Securities",
	"หลักทรัพย์ไม่จดทะเบียน":                         "Unlisted Securities",
	"เงินฝากและรับฝาก":                               "Deposits and Custody",
	"เงินฝากประจำ":                                   "Fixed Deposit",
	"เงินฝากออมทรัพย์":                               "Savings Deposit",
	"เงินฝากกระแสรายวัน":                             "Current Account Deposit",
	"เงินฝากต่างประเทศ":                              "Foreign Currency Deposit",
	"เงินฝากอื่น":                                    "Other Deposit",
	"สัญญาซื้อขายล่วงหน้า":                           "Forward Contract",
	"สัญญาซื้อขายล่วงหน้าอัตราแลกเปลี่ยน":            "Foreign Exchange Forward",
	"สัญญาซื้อขายล่วงหน้าดัชนี":                      "Index Forward",
	"สัญญาซื้อขายล่วงหน้าสินค้าโภคภัณฑ์":             "Commodity Forward",
	"สัญญาซื้อขายล่วงหน้าอัตราดอกเบี้ย":              "Interest Rate Forward",
	"สัญญาซื้อขายสินทรัพย์ล่วงหน้า":                  "Forward Contract",
	"สัญญาซื้อขายอนุพันธ์":                           "Derivative Contract",
	"สัญญาซื้อขายอนุพันธ์อัตราแลกเปลี่ยน":            "Foreign Exchange Derivative",
	"สัญญาซื้อขายอนุพันธ์ดัชนี":                      "Index Derivative",
	"สัญญาซื้อขายอนุพันธ์สินค้าโภคภัณฑ์":             "Commodity Derivative",
	"สัญญาซื้อขายอนุพันธ์อัตราดอกเบี้ย":              "Interest Rate Derivative",
	"สัญญาซื้อขายอนุพันธ์อื่น":                       "Other Derivative",
	"สัญญาซื้อขายอนุพันธ์เครดิต":                     "Credit Derivative",
	"สัญญาซื้อขายอนุพันธ์สิทธิ":                      "Equity Derivative",
	"สัญญาซื้อขายอนุพันธ์หนี้":                       "Debt Derivative",
	"สัญญาซื้อขายอนุพันธ์อัตราแลกเปลี่ยนและดอกเบี้ย": "Foreign Exchange and Interest Rate Derivative",
	"สัญญาซื้อขายอนุพันธ์สินค้าโภคภัณฑ์และอื่น":      "Commodity and Other Derivative",
	"สัญญาซื้อขายอนุพันธ์รวม":                        "Combined Derivative",
	"สัญญาซื้อขายอนุพันธ์อื่นๆ":                      "Other Derivatives",
	"สัญญาฟอร์เวิร์ด":                                "Forward Contract",
	"หน่วยลงทุนของกองทุนตราสารทุน":                   "Equity Fund Units",
	"หน่วยลงทุนของกองทุนตราสารหนี้":                  "Debt Fund Units",
	"หน่วยลงทุนของกองทุนรวม":                         "Mixed Fund Units",
	"หน่วยลงทุนของกองทุนหุ้น":                        "Stock Fund Units",
	"หน่วยลงทุนของกองทุนตราสารหนี้ระยะสั้น":          "Short-term Debt Fund Units",
	"หลักทรัพย์ค้ำประกัน":                            "Collateral Securities",
	"เงินค้ำประกัน":                                  "Cash Collateral",
	"หลักประกันอื่น":                                 "Other Collateral",
	"หลักทรัพย์กู้":                                  "Borrowed Securities",
	"เงินกู้หลักทรัพย์":                              "Securities Lending",
	"เงินกู้และหนี้สินอื่น":                          "Loans and Other Debts",
	"หนี้สินอื่น":                                    "Other Liabilities",
	"รายได้ค้างรับ":                                  "Accrued Income",
	"ค่าใช้จ่ายค้างจ่าย":                             "Accrued Expenses",
	"ภาษีมูลค่าเพิ่ม":                                "Value Added Tax",
	"ภาษีธุรกิจเฉพาะ":                                "Specific Business Tax",
	"ภาษีอากรอื่น":                                   "Other Taxes",
	"ค่าธรรมเนียมและค่าใช้จ่าย":                      "Fees and Expenses",
	"ค่าธรรมเนียมการจัดการ":                          "Management Fee",
	"ค่าธรรมเนียมผู้ดูแลผลประโยชน์":                  "Trustee Fee",
	"ค่าธรรมเนียมนายทะเบียน":                         "Registrar Fee",
	"ค่าธรรมเนียมการขาย":                             "Front-end Fee",
	"ค่าธรรมเนียมการรับซื้อคืน":                      "Back-end Fee",
	"ค่าธรรมเนียมการสับเปลี่ยน":                      "Switching Fee",
	"ค่าธรรมเนียมการโอน":                             "Transfer Fee",
	"ค่าใช้จ่ายอื่นๆ":                                "Other Expenses",
	"ค่าใช้จ่ายรวม":                                  "Total Expenses",
	"ค่าใช้จ่ายเฉลี่ย":                               "Average Expenses",
	"กำไรขั้นต้น":                                    "Gross Profit",
	"กำไรจากการดำเนินงาน":                            "Operating Profit",
	"กำไรก่อนหักภาษี":                                "Profit Before Tax",
	"กำไรสุทธิ":                                      "Net Profit",
	"รายได้รวม":                                      "Total Revenue",
	"รายได้จากการขาย":                                "Sales Revenue",
	"รายได้จากการให้บริการ":                          "Service Revenue",
	"รายได้ดอกเบี้ย":                                 "Interest Income",
	"รายได้ปันผล":                                    "Dividend Income",
	"รายได้จากการลงทุน":                              "Investment Income",
	"รายได้อื่น":                                     "Other Income",
	"ต้นทุนขาย":                                      "Cost of Goods Sold",
	"ค่าใช้จ่ายในการขาย":                             "Selling Expenses",
	"ค่าใช้จ่ายในการบริหาร":                          "Administrative Expenses",
	"ค่าใช้จ่ายอื่น":                                 "Other Expenses",
	"ต้นทุนและค่าใช้จ่ายรวม":                         "Total Costs and Expenses",
	" EBITDA":              "EBITDA",
	" EBIT":                "EBIT",
	" EBITDA Margin":       "EBITDA Margin",
	" EBIT Margin":         "EBIT Margin",
	" Net Profit Margin":   "Net Profit Margin",
	" Gross Profit Margin": "Gross Profit Margin",
	" ROA":                 "Return on Assets",
	" ROE":                 "Return on Equity",
	" Debt to Equity":      "Debt-to-Equity Ratio",
	" Current Ratio":       "Current Ratio",
	" Quick Ratio":         "Quick Ratio",
	" Interest Coverage":   "Interest Coverage Ratio",
	" Revenue Growth":      "Revenue Growth",
	" Profit Growth":       "Profit Growth",
	" Asset Growth":        "Asset Growth",
	" Equity Growth":       "Equity Growth",
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

// TranslateSubscriptionRedemptionPeriod translates the Thai period description
// on a FactsheetSubscriptionRedemptionPeriod to English when useEnglish is true.
// It mutates the provided period in place.
func TranslateSubscriptionRedemptionPeriod(period *FactsheetSubscriptionRedemptionPeriod, useEnglish bool) {
	if !useEnglish {
		return
	}
	if trans, ok := SubscriptionPeriodTranslation[period.Period]; ok {
		period.Period = trans
	}
	if period.RedempPeriodOth != "" {
		if trans, ok := SubscriptionPeriodTranslation[period.RedempPeriodOth]; ok {
			period.RedempPeriodOth = trans
		}
	}
}

// TranslateAllSubscriptionRedemptionPeriods translates every period in the
// slice when useEnglish is true.
func TranslateAllSubscriptionRedemptionPeriods(periods []FactsheetSubscriptionRedemptionPeriod, useEnglish bool) {
	for i := range periods {
		TranslateSubscriptionRedemptionPeriod(&periods[i], useEnglish)
	}
}

// TranslateAllMonthlyPortfolioAssetTypes translates every portfolio item in
// the slice when useEnglish is true.
func TranslateAllMonthlyPortfolioAssetTypes(items []MonthlyPortfolioAssetType, useEnglish bool) {
	for i := range items {
		TranslateMonthlyPortfolioAssetType(&items[i], useEnglish)
	}
}
