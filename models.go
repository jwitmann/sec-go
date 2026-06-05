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
