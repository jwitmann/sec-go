package sec

import (
	"context"
	"net/url"
)

// PVDListItem represents a provident fund from the getPvdList endpoint.
type PVDListItem struct {
	UniqueID    string   `json:"unique_id"`
	CompNameTH  string   `json:"comp_name_th"`
	CompNameEN  string   `json:"comp_name_en"`
	LastUpdDate DateTime `json:"last_upd_date"`
}

// PVDFundInfo represents provident fund information.
type PVDFundInfo struct {
	ProjID          string   `json:"proj_id"`
	ProjNameTH      string   `json:"proj_name_th"`
	ProjNameEN      string   `json:"proj_name_en"`
	ProjAbbrName    string   `json:"proj_abbr_name"`
	FundStatus      string   `json:"fund_status"`
	InitDate        string   `json:"init_date"`
	RegisDate       string   `json:"regis_date"`
	CancelDate      string   `json:"cancel_date"`
	PolicyDesc      string   `json:"policy_desc"`
	ManagementStyle string   `json:"management_style"`
	FundClassName   string   `json:"fund_class_name"`
	LastUpdDate     DateTime `json:"last_upd_date"`
}

// PVDFundSpec represents provident fund specifications.
type PVDFundSpec struct {
	ProjID        string   `json:"proj_id"`
	FundClassName string   `json:"fund_class_name"`
	SpecCode      string   `json:"spec_code"`
	SpecDesc      string   `json:"spec_desc"`
	LastUpdDate   DateTime `json:"last_upd_date"`
}

// PVDFundMember represents provident fund member statistics.
type PVDFundMember struct {
	ProjID         string `json:"proj_id"`
	Period         int    `json:"period"`
	TotalMembers   int    `json:"total_members"`
	ActiveMembers  int    `json:"active_members"`
	RetiredMembers int    `json:"retired_members"`
	LastUpdDate    string `json:"last_upd_date"`
}

// PVDFundAsset represents provident fund asset allocation.
type PVDFundAsset struct {
	ProjID      string  `json:"proj_id"`
	Period      int     `json:"period"`
	AssetSeq    int     `json:"asset_seq"`
	AssetName   string  `json:"asset_name"`
	AssetRatio  float64 `json:"asset_ratio"`
	MarketValue float64 `json:"market_value"`
	LastUpdDate string  `json:"last_upd_date"`
}

// PVDFundTransaction represents provident fund transaction data.
type PVDFundTransaction struct {
	ProjID      string  `json:"proj_id"`
	Period      int     `json:"period"`
	TransType   string  `json:"trans_type"`
	TransValue  float64 `json:"trans_value"`
	TransCount  int     `json:"trans_count"`
	LastUpdDate string  `json:"last_upd_date"`
}

// PVDFundContribution represents provident fund contribution records.
type PVDFundContribution struct {
	ProjID          string  `json:"proj_id"`
	Period          int     `json:"period"`
	EmployeeContrib float64 `json:"employee_contrib"`
	EmployerContrib float64 `json:"employer_contrib"`
	TotalContrib    float64 `json:"total_contrib"`
	LastUpdDate     string  `json:"last_upd_date"`
}

// PVDFundExpense represents provident fund expense data.
type PVDFundExpense struct {
	ProjID       string  `json:"proj_id"`
	Period       int     `json:"period"`
	ExpenseType  string  `json:"expense_type"`
	ExpenseValue float64 `json:"expense_value"`
	ExpenseRatio float64 `json:"expense_ratio"`
	LastUpdDate  string  `json:"last_upd_date"`
}

// PVDFundLiquidity represents provident fund liquidity metrics.
type PVDFundLiquidity struct {
	ProjID       string  `json:"proj_id"`
	Period       int     `json:"period"`
	CashRatio    float64 `json:"cash_ratio"`
	CurrentRatio float64 `json:"current_ratio"`
	QuickRatio   float64 `json:"quick_ratio"`
	LastUpdDate  string  `json:"last_upd_date"`
}

// PVDFundPerformance represents provident fund performance data.
type PVDFundPerformance struct {
	ProjID           string   `json:"proj_id"`
	FundClassName    string   `json:"fund_class_name"`
	Period           int      `json:"period"`
	PerformanceType  string   `json:"performance_type"`
	ReferencePeriod  string   `json:"reference_period"`
	PerformanceValue float64  `json:"performance_value"`
	LastUpdDate      DateTime `json:"last_upd_date"`
}

// PVDFundBenchmark represents provident fund benchmark comparisons.
type PVDFundBenchmark struct {
	ProjID         string   `json:"proj_id"`
	FundClassName  string   `json:"fund_class_name"`
	Period         int      `json:"period"`
	Benchmark      string   `json:"benchmark"`
	BenchmarkValue float64  `json:"benchmark_value"`
	FundValue      float64  `json:"fund_value"`
	LastUpdDate    DateTime `json:"last_upd_date"`
}

// PVDFundDividend represents provident fund dividend history.
type PVDFundDividend struct {
	ProjID        string   `json:"proj_id"`
	FundClassName string   `json:"fund_class_name"`
	DividendDate  string   `json:"dividend_date"`
	DividendValue float64  `json:"dividend_value"`
	BookCloseDate string   `json:"book_close_date"`
	LastUpdDate   DateTime `json:"last_upd_date"`
}

// PVDFundPolicy represents provident fund investment policies.
type PVDFundPolicy struct {
	ProjID         string   `json:"proj_id"`
	FundClassName  string   `json:"fund_class_name"`
	PolicyType     string   `json:"policy_type"`
	PolicyDesc     string   `json:"policy_desc"`
	MinEquityRatio float64  `json:"min_equity_ratio"`
	MaxEquityRatio float64  `json:"max_equity_ratio"`
	LastUpdDate    DateTime `json:"last_upd_date"`
}

// PVDFundFee represents provident fund fee structures.
type PVDFundFee struct {
	ProjID        string   `json:"proj_id"`
	FundClassName string   `json:"fund_class_name"`
	FeeTypeDesc   string   `json:"fee_type_desc"`
	Rate          float64  `json:"rate"`
	RateUnit      string   `json:"rate_unit"`
	ActualValue   float64  `json:"actual_value"`
	LastUpdDate   DateTime `json:"last_upd_date"`
}

// PVDFundCompliance represents provident fund compliance data.
type PVDFundCompliance struct {
	ProjID           string   `json:"proj_id"`
	Period           int      `json:"period"`
	ComplianceType   string   `json:"compliance_type"`
	ComplianceStatus string   `json:"compliance_status"`
	Remark           string   `json:"remark"`
	LastUpdDate      DateTime `json:"last_upd_date"`
}

// PVDListOptions defines query options for ListPVDs.
type PVDListOptions struct {
	PageSize int
	Cursor   string
}

// PVDProjOptions defines query options for PVD endpoints that filter by proj_id.
type PVDProjOptions struct {
	PageSize int
	Cursor   string
	ProjID   string
}

// PVDPeriodOptions defines query options for PVD endpoints that filter by period.
type PVDPeriodOptions struct {
	PageSize    int
	Cursor      string
	ProjID      string
	StartPeriod int
	EndPeriod   int
}

// ListPVDs returns the list of provident funds.
func (c *Client) ListPVDs(ctx context.Context, opts PVDListOptions) ([]PVDListItem, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	path := buildPath("/v1/pvd/general-info/list", params)
	return fetchPaginated[PVDListItem](ctx, c, path, "list PVDs")
}

// GetPVDFundInfo returns provident fund information.
func (c *Client) GetPVDFundInfo(ctx context.Context, opts PVDProjOptions) ([]PVDFundInfo, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v1/pvd/general-info/fund-info", params)
	return fetchPaginated[PVDFundInfo](ctx, c, path, "get PVD fund info")
}

// GetPVDFundSpecs returns provident fund specifications.
func (c *Client) GetPVDFundSpecs(ctx context.Context, opts PVDProjOptions) ([]PVDFundSpec, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v1/pvd/general-info/fund-spec", params)
	return fetchPaginated[PVDFundSpec](ctx, c, path, "get PVD fund specs")
}

// GetPVDFundMembers returns provident fund member statistics.
func (c *Client) GetPVDFundMembers(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundMember, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-member", params)
	return fetchPaginated[PVDFundMember](ctx, c, path, "get PVD fund members")
}

// GetPVDFundAssets returns provident fund asset allocation.
func (c *Client) GetPVDFundAssets(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundAsset, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-asset", params)
	return fetchPaginated[PVDFundAsset](ctx, c, path, "get PVD fund assets")
}

// GetPVDFundTransactions returns provident fund transaction data.
func (c *Client) GetPVDFundTransactions(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundTransaction, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-transaction", params)
	return fetchPaginated[PVDFundTransaction](ctx, c, path, "get PVD fund transactions")
}

// GetPVDFundContributions returns provident fund contribution records.
func (c *Client) GetPVDFundContributions(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundContribution, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-contribution", params)
	return fetchPaginated[PVDFundContribution](ctx, c, path, "get PVD fund contributions")
}

// GetPVDFundExpenses returns provident fund expense data.
func (c *Client) GetPVDFundExpenses(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundExpense, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-expense", params)
	return fetchPaginated[PVDFundExpense](ctx, c, path, "get PVD fund expenses")
}

// GetPVDFundLiquidity returns provident fund liquidity metrics.
func (c *Client) GetPVDFundLiquidity(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundLiquidity, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-liquidity", params)
	return fetchPaginated[PVDFundLiquidity](ctx, c, path, "get PVD fund liquidity")
}

// GetPVDFundPerformance returns provident fund performance data.
func (c *Client) GetPVDFundPerformance(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundPerformance, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-performance", params)
	return fetchPaginated[PVDFundPerformance](ctx, c, path, "get PVD fund performance")
}

// GetPVDFundBenchmarks returns provident fund benchmark comparisons.
func (c *Client) GetPVDFundBenchmarks(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundBenchmark, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-benchmark", params)
	return fetchPaginated[PVDFundBenchmark](ctx, c, path, "get PVD fund benchmarks")
}

// GetPVDFundDividends returns provident fund dividend history.
func (c *Client) GetPVDFundDividends(ctx context.Context, opts PVDProjOptions) ([]PVDFundDividend, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v1/pvd/fund-dividend", params)
	return fetchPaginated[PVDFundDividend](ctx, c, path, "get PVD fund dividends")
}

// GetPVDFundPolicies returns provident fund investment policies.
func (c *Client) GetPVDFundPolicies(ctx context.Context, opts PVDProjOptions) ([]PVDFundPolicy, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v1/pvd/fund-policy", params)
	return fetchPaginated[PVDFundPolicy](ctx, c, path, "get PVD fund policies")
}

// GetPVDFundFees returns provident fund fee structures.
func (c *Client) GetPVDFundFees(ctx context.Context, opts PVDProjOptions) ([]PVDFundFee, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v1/pvd/fund-fee", params)
	return fetchPaginated[PVDFundFee](ctx, c, path, "get PVD fund fees")
}

// GetPVDFundCompliance returns provident fund compliance data.
func (c *Client) GetPVDFundCompliance(ctx context.Context, opts PVDPeriodOptions) ([]PVDFundCompliance, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	setPeriodRange(params, opts.ProjID, opts.StartPeriod, opts.EndPeriod)
	path := buildPath("/v1/pvd/fund-compliance", params)
	return fetchPaginated[PVDFundCompliance](ctx, c, path, "get PVD fund compliance")
}
