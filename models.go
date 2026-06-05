package sec

import "time"

type PaginatedResponse struct {
	Message    string `json:"message"`
	PageSize   int    `json:"page_size"`
	NextCursor string `json:"next_cursor"`
}

type AMC struct {
	UniqueID    string    `json:"unique_id"`
	CompNameTH  string    `json:"comp_name_th"`
	CompNameEN  string    `json:"comp_name_en"`
	LastUpdDate time.Time `json:"last_upd_date"`
}

type FundProfile struct {
	UniqueID                     string    `json:"unique_id"`
	CompNameTH                   string    `json:"comp_name_th"`
	CompNameEN                   string    `json:"comp_name_en"`
	ProjID                       string    `json:"proj_id"`
	RegisID                      string    `json:"regis_id"`
	InitDate                     string    `json:"init_date"`
	RegisDate                    string    `json:"regis_date"`
	CancelDate                   string    `json:"cancel_date"`
	ProjNameTH                   string    `json:"proj_name_th"`
	ProjNameEN                   string    `json:"proj_name_en"`
	ProjAbbrName                 string    `json:"proj_abbr_name"`
	FundStatus                   string    `json:"fund_status"`
	InvestCountryFlag            string    `json:"invest_country_flag"`
	ProjRetailType               string    `json:"proj_retail_type"`
	ProjTermFlag                 string    `json:"proj_term_flag"`
	ProjTermDay                  string    `json:"proj_term_day"`
	ProjTermMonth                string    `json:"proj_term_month"`
	ProjTermYear                 string    `json:"proj_term_year"`
	PolicyDesc                   string    `json:"policy_desc"`
	InvestmentPolicyDesc         string    `json:"investment_policy_desc"`
	ManagementStyle              string    `json:"management_style"`
	FeederFundMasterFund         string    `json:"feederfund_master_fund"`
	FeederFundCountry            string    `json:"feederfund_country"`
	ExchangeRateProtectionPolicy string    `json:"exchange_rate_protection_policy"`
	FundClassName                string    `json:"fund_class_name"`
	FundClassDetail              string    `json:"fund_class_detail"`
	FundClassDescription         string    `json:"fund_class_description"`
	FundClassTaxIncentiveType    string    `json:"fund_class_tax_incentive_type"`
	FundClassISINCode            string    `json:"fund_class_isin_code"`
	LastUpdDate                  time.Time `json:"last_upd_date"`
}

type DailyNAV struct {
	ProjID        string    `json:"proj_id"`
	UniqueID      string    `json:"unique_id"`
	FundClassName string    `json:"fund_class_name"`
	NavDate       string    `json:"nav_date"`
	NetAsset      float64   `json:"net_asset"`
	LastVal       float64   `json:"last_val"`
	SellPrice     float64   `json:"sell_price"`
	BuyPrice      float64   `json:"buy_price"`
	SellSwapPrice float64   `json:"sell_swap_price"`
	BuySwapPrice  float64   `json:"buy_swap_price"`
	LastUpdDate   time.Time `json:"last_upd_date"`
}

type MutualFundFee struct {
	ProjID        string    `json:"proj_id"`
	FundClassName string    `json:"fund_class_name"`
	FeeTypeDesc   string    `json:"fee_type_desc"`
	Rate          float64   `json:"rate"`
	RateUnit      string    `json:"rate_unit"`
	FeeOtherDesc  string    `json:"fee_other_desc"`
	LastUpdDate   time.Time `json:"last_upd_date"`
}

type FactsheetFee struct {
	ProjID         string    `json:"proj_id"`
	FundClassName  string    `json:"fund_class_name"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ProspectusType string    `json:"prospectus_type"`
	FeeTypeDesc    string    `json:"fee_type_desc"`
	Rate           float64   `json:"rate"`
	ActualValue    float64   `json:"actual_value"`
	FeeOtherDesc   string    `json:"fee_other_desc"`
	LastUpdDate    time.Time `json:"last_upd_date"`
}

type FactsheetPerformance struct {
	ProjID              string    `json:"proj_id"`
	FundClassName       string    `json:"fund_class_name"`
	StartDate           string    `json:"start_date"`
	EndDate             string    `json:"end_date"`
	ProspectusType      string    `json:"prospectus_type"`
	PerformanceTypeDesc string    `json:"performance_type_desc"`
	ReferencePeriod     string    `json:"reference_period"`
	PerformanceValue    float64   `json:"performance_value"`
	LastUpdDate         time.Time `json:"last_upd_date"`
}

type DividendHistory struct {
	ProjID        string    `json:"proj_id"`
	UniqueID      string    `json:"unique_id"`
	ClassAbbrName string    `json:"class_abbr_name"`
	BookCloseDate string    `json:"book_close_date"`
	DividendDate  string    `json:"dividend_date"`
	DividendValue float64   `json:"dividend_value"`
	LastUpdDate   time.Time `json:"last_upd_date"`
}

type AssetAllocation struct {
	ProjID         string    `json:"proj_id"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ProspectusType string    `json:"prospectus_type"`
	AssetSeq       int       `json:"asset_seq"`
	AssetName      string    `json:"asset_name"`
	AssetRatio     float64   `json:"asset_ratio"`
	LastUpdDate    time.Time `json:"last_upd_date"`
}

type RiskSpectrum struct {
	ProjID           string    `json:"proj_id"`
	StartDate        string    `json:"start_date"`
	EndDate          string    `json:"end_date"`
	ProspectusType   string    `json:"prospectus_type"`
	RiskSpectrum     string    `json:"risk_spectrum"`
	RiskSpectrumDesc string    `json:"risk_spectrum_desc"`
	LastUpdDate      time.Time `json:"last_upd_date"`
}

type Top5Holding struct {
	ProjID         string    `json:"proj_id"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ProspectusType string    `json:"prospectus_type"`
	AssetSeq       int       `json:"asset_seq"`
	AssetName      string    `json:"asset_name"`
	AssetRatio     float64   `json:"asset_ratio"`
	LastUpdDate    time.Time `json:"last_upd_date"`
}

type QuarterlyPortfolio struct {
	ProjID         string  `json:"proj_id"`
	Period         int     `json:"period"`
	AsOfDate       string  `json:"as_of_date"`
	AssetliabID    string  `json:"assetliab_id"`
	AssetliabDesc  string  `json:"assetliab_desc"`
	IssueCode      string  `json:"issue_code"`
	ISINCode       string  `json:"isin_code"`
	Issuer         string  `json:"issuer"`
	AssetliabValue float64 `json:"assetliab_value"`
	PercentNAV     float64 `json:"percent_nav"`
	LastUpdDate    string  `json:"last_upd_date"`
}

type MonthlyPortfolioAssetType struct {
	ProjID        string  `json:"proj_id"`
	Period        int     `json:"period"`
	AssetliabCode string  `json:"assetliab_code"`
	AssetliabDesc string  `json:"assetliab_desc"`
	MarketValue   float64 `json:"market_value"`
	PercentNAV    float64 `json:"percent_nav"`
}

type FactsheetSubscriptionRedemptionMinimum struct {
	ProjID             string    `json:"proj_id"`
	FundClassName      string    `json:"fund_class_name"`
	StartDate          string    `json:"start_date"`
	EndDate            string    `json:"end_date"`
	ProspectusType     string    `json:"prospectus_type"`
	MinimumSubIPO      float64   `json:"minimum_sub_ipo"`
	MinimumSubIPOCur   string    `json:"minimum_sub_ipo_cur"`
	MinimumSub         float64   `json:"minimum_sub"`
	MinimumSubCur      string    `json:"minimum_sub_cur"`
	MinimumSubUnit     string    `json:"minimum_sub_unit"`
	MinimumRedempt     float64   `json:"minimum_redempt"`
	MinimumRedemptCur  string    `json:"minimum_redempt_cur"`
	MinimumRedemptUnit string    `json:"minimum_redempt_unit"`
	LowbalVal          float64   `json:"lowbal_val"`
	LowbalValCur       string    `json:"lowbal_val_cur"`
	LowbalUnit         string    `json:"lowbal_unit"`
	LastUpdDate        time.Time `json:"last_upd_date"`
}

type FactsheetStatistics struct {
	ProjID                  string    `json:"proj_id"`
	FundClassName           string    `json:"fund_class_name"`
	StartDate               string    `json:"start_date"`
	EndDate                 string    `json:"end_date"`
	ProspectusType          string    `json:"prospectus_type"`
	PortfolioTurnoverRatio  string    `json:"portfolio_turnover_ratio"`
	RecoveringPeriod        string    `json:"recovering_period"`
	PortfolioDurationPeriod string    `json:"portfolio_duration_period"`
	MaximumDrawdown         string    `json:"maximum_drawdown"`
	SharpeRatio             string    `json:"sharpe_ratio"`
	Beta                    string    `json:"beta"`
	Alpha                   string    `json:"alpha"`
	FXHedging               string    `json:"fx_hedging"`
	TrackingError           string    `json:"tracking_error"`
	YieldToMaturity         string    `json:"yield_to_maturity"`
	LastUpdDate             time.Time `json:"last_upd_date"`
}

type FactsheetSubscriptionRedemptionPeriod struct {
	ProjID           string    `json:"proj_id"`
	FundClassName    string    `json:"fund_class_name"`
	StartDate        string    `json:"start_date"`
	EndDate          string    `json:"end_date"`
	ProspectusType   string    `json:"prospectus_type"`
	Type             string    `json:"type"`
	Period           string    `json:"period"`
	RedempPeriodOth  string    `json:"redemp_period_oth"`
	SettlementPeriod string    `json:"settlement_period"`
	LastUpdDate      time.Time `json:"last_upd_date"`
}

type FundDividendPolicy struct {
	ProjID         string    `json:"proj_id"`
	FundClassName  string    `json:"fund_class_name"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ProspectusType string    `json:"prospectus_type"`
	DividendPolicy string    `json:"dividend_policy"`
	LastUpdDate    time.Time `json:"last_upd_date"`
}

type FundSpecification struct {
	ProjID        string    `json:"proj_id"`
	FundClassName string    `json:"fund_class_name"`
	SpecCode      string    `json:"spec_code"`
	SpecDesc      string    `json:"spec_desc"`
	LastUpdDate   time.Time `json:"last_upd_date"`
}

type FactsheetBenchmark struct {
	ProjID         string    `json:"proj_id"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ProspectusType string    `json:"prospectus_type"`
	GroupSeq       int       `json:"group_seq"`
	Benchmark      string    `json:"benchmark"`
	Remark         string    `json:"benchmark_remark"`
	LastUpdDate    time.Time `json:"last_upd_date"`
}

type FundInvolveParty struct {
	ProjID       string    `json:"proj_id"`
	EntityType   string    `json:"entity_type"`
	EntityNameTH string    `json:"entity_name_th"`
	EntityNameEN string    `json:"entity_name_en"`
	Address      string    `json:"address"`
	LastUpdDate  time.Time `json:"last_upd_date"`
}

type FundFactsheetURL struct {
	ProjID          string    `json:"proj_id"`
	FundClassName   string    `json:"fund_class_name"`
	ProspectusType  string    `json:"prospectus_type"`
	AMCURLFactsheet string    `json:"amc_url_factsheet"`
	PDFFactsheet    string    `json:"pdf_factsheet"`
	AsOfDate        string    `json:"as_of_date"`
	LastUpdDate     time.Time `json:"last_upd_date"`
}

type FundIPO struct {
	ProjID             string    `json:"proj_id"`
	StartDate          string    `json:"start_date"`
	EndDate            string    `json:"end_date"`
	ProspectusType     string    `json:"prospectus_type"`
	FirstSellStartDate string    `json:"first_sell_start_date"`
	FirstSellEndDate   string    `json:"first_sell_end_date"`
	LastUpdDate        time.Time `json:"last_upd_date"`
}
