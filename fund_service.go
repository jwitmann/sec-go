package sec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func fetchPaginated[T any](ctx context.Context, c *Client, path string, op string) ([]T, string, error) {
	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	var response struct {
		PaginatedResponse
		Items []T `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal %s: %w", op, err)
	}

	return response.Items, response.NextCursor, nil
}

func setPagination(params url.Values, pageSize int, cursor string) {
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if cursor != "" {
		params.Set("next_cursor", cursor)
	}
}

func setDateRange(params url.Values, startKey, endKey string, startDate, endDate time.Time) {
	if !startDate.IsZero() {
		params.Set(startKey, startDate.Format("2006-01-02"))
	}
	if !endDate.IsZero() {
		params.Set(endKey, endDate.Format("2006-01-02"))
	}
}

func buildPath(base string, params url.Values) string {
	if len(params) == 0 {
		return base
	}
	return base + "?" + params.Encode()
}

func (c *Client) ListAMCs(ctx context.Context, pageSize int, cursor string) ([]AMC, string, error) {
	params := url.Values{}
	setPagination(params, pageSize, cursor)

	path := buildPath("/v2/fund/general-info/amcs", params)
	return fetchPaginated[AMC](ctx, c, path, "list AMCs")
}

func (c *Client) GetFundProfiles(ctx context.Context, opts ProfileOptions) ([]FundProfile, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}
	if opts.FundStatus != "" {
		params.Set("fund_status", opts.FundStatus)
	}
	if opts.ProjectInfo != "" {
		params.Set("project_info", opts.ProjectInfo)
	}
	if opts.CompanyInfo != "" {
		params.Set("company_info", opts.CompanyInfo)
	}

	path := buildPath("/v2/fund/general-info/profiles", params)
	return fetchPaginated[FundProfile](ctx, c, path, "get fund profiles")
}

type ProfileOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	FundClassName string
	FundStatus    string
	ProjectInfo   string
	CompanyInfo   string
}

func (c *Client) GetDailyNAV(ctx context.Context, opts NAVOptions) ([]DailyNAV, string, error) {
	path := buildPath("/v2/fund/daily-info/nav", buildNAVParams(opts))
	return fetchPaginated[DailyNAV](ctx, c, path, "get daily NAV")
}

func buildNAVParams(opts NAVOptions) url.Values {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	setDateRange(params, "start_nav_date", "end_nav_date", opts.StartDate, opts.EndDate)
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}
	return params
}

type NAVOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	FundClassName string
}

func (c *Client) GetMutualFundFees(ctx context.Context, opts FeeOptions) ([]MutualFundFee, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	path := buildPath("/v2/fund/general-info/mutual-fund-fees", params)
	return fetchPaginated[MutualFundFee](ctx, c, path, "get mutual fund fees")
}

func (c *Client) GetFundSpecifications(ctx context.Context, opts FeeOptions) ([]FundSpecification, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	path := buildPath("/v2/fund/general-info/specifications", params)
	return fetchPaginated[FundSpecification](ctx, c, path, "get fund specifications")
}

func (c *Client) GetFundInvolveParties(ctx context.Context, opts InvolvePartyOptions) ([]FundInvolveParty, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.EntityType != "" {
		params.Set("entity_type", opts.EntityType)
	}

	path := buildPath("/v2/fund/general-info/involve-parties", params)
	return fetchPaginated[FundInvolveParty](ctx, c, path, "get fund involve parties")
}

type InvolvePartyOptions struct {
	PageSize   int
	Cursor     string
	ProjID     string
	EntityType string
}

func (c *Client) GetFundFactsheetURLs(ctx context.Context, opts FeeOptions) ([]FundFactsheetURL, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	path := buildPath("/v2/fund/factsheet/urls", params)
	return fetchPaginated[FundFactsheetURL](ctx, c, path, "get fund factsheet URLs")
}

func (c *Client) GetFundIPOs(ctx context.Context, opts FactsheetOptions) ([]FundIPO, string, error) {
	path := "/v2/fund/factsheet/ipos" + buildFactsheetQuery(opts)
	return fetchPaginated[FundIPO](ctx, c, path, "get fund IPOs")
}

func (c *Client) GetFactsheetFees(ctx context.Context, opts FactsheetOptions) ([]FactsheetFee, string, error) {
	path := "/v2/fund/factsheet/fees" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetFee](ctx, c, path, "get factsheet fees")
}

func (c *Client) GetFactsheetPerformance(ctx context.Context, opts FactsheetOptions) ([]FactsheetPerformance, string, error) {
	path := "/v2/fund/factsheet/performance" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetPerformance](ctx, c, path, "get factsheet performance")
}

func (c *Client) GetFactsheetSubscriptionRedemptionMinimums(ctx context.Context, opts FactsheetOptions) ([]FactsheetSubscriptionRedemptionMinimum, string, error) {
	path := "/v2/fund/factsheet/subscription-redemption-minimums" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetSubscriptionRedemptionMinimum](ctx, c, path, "get subscription-redemption minimums")
}

func (c *Client) GetFactsheetSubscriptionRedemptionPeriods(ctx context.Context, opts FactsheetOptions) ([]FactsheetSubscriptionRedemptionPeriod, string, error) {
	path := "/v2/fund/factsheet/subscription-redemption-periods" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetSubscriptionRedemptionPeriod](ctx, c, path, "get subscription-redemption periods")
}

func (c *Client) GetFactsheetStatistics(ctx context.Context, opts FactsheetOptions) ([]FactsheetStatistics, string, error) {
	path := "/v2/fund/factsheet/statistics" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetStatistics](ctx, c, path, "get factsheet statistics")
}

func (c *Client) GetFactsheetDividendPolicy(ctx context.Context, opts FactsheetOptions) ([]FundDividendPolicy, string, error) {
	path := "/v2/fund/factsheet/dividend-policy" + buildFactsheetQuery(opts)
	return fetchPaginated[FundDividendPolicy](ctx, c, path, "get factsheet dividend policy")
}

func (c *Client) GetFactsheetBenchmarks(ctx context.Context, opts FactsheetOptions) ([]FactsheetBenchmark, string, error) {
	path := "/v2/fund/factsheet/benchmarks" + buildFactsheetQuery(opts)
	return fetchPaginated[FactsheetBenchmark](ctx, c, path, "get factsheet benchmarks")
}

type FeeOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	FundClassName string
}

type FactsheetOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	Latest        bool
	FundClassName string
}

func (c *Client) GetDividendHistory(ctx context.Context, opts DividendHistoryOptions) ([]DividendHistory, string, error) {
	path := buildPath("/v2/fund/daily-info/dividend-history", buildDividendParams(opts))
	return fetchPaginated[DividendHistory](ctx, c, path, "get dividend history")
}

func buildDividendParams(opts DividendHistoryOptions) url.Values {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	setDateRange(params, "start_dividend_date", "end_dividend_date", opts.StartDate, opts.EndDate)
	if opts.ClassAbbrName != "" {
		params.Set("class_abbr_name", opts.ClassAbbrName)
	}
	return params
}

type DividendHistoryOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	ClassAbbrName string
}

func (c *Client) GetAssetAllocation(ctx context.Context, opts FactsheetOptions) ([]AssetAllocation, string, error) {
	path := "/v2/fund/factsheet/asset-allocation" + buildFactsheetQuery(opts)
	return fetchPaginated[AssetAllocation](ctx, c, path, "get asset allocation")
}

func (c *Client) GetRiskSpectrum(ctx context.Context, opts FactsheetOptions) ([]RiskSpectrum, string, error) {
	path := "/v2/fund/factsheet/risk-spectrum" + buildFactsheetQuery(opts)
	return fetchPaginated[RiskSpectrum](ctx, c, path, "get risk spectrum")
}

func (c *Client) GetTop5Holdings(ctx context.Context, opts FactsheetOptions) ([]Top5Holding, string, error) {
	path := "/v2/fund/factsheet/top5-holdings" + buildFactsheetQuery(opts)
	return fetchPaginated[Top5Holding](ctx, c, path, "get top 5 holdings")
}

func (c *Client) GetQuarterlyPortfolio(ctx context.Context, opts OutstandingOptions) ([]QuarterlyPortfolio, string, error) {
	path := "/v2/fund/outstanding/portfolio" + buildOutstandingQuery(opts)
	return fetchPaginated[QuarterlyPortfolio](ctx, c, path, "get quarterly portfolio")
}

func (c *Client) GetMonthlyPortfolioAssetType(ctx context.Context, opts OutstandingOptions) ([]MonthlyPortfolioAssetType, string, error) {
	path := "/v2/fund/outstanding/portfolio-asset-type" + buildOutstandingQuery(opts)
	return fetchPaginated[MonthlyPortfolioAssetType](ctx, c, path, "get monthly portfolio asset type")
}

type OutstandingOptions struct {
	PageSize    int
	Cursor      string
	ProjID      string
	StartPeriod string
	EndPeriod   string
}

func buildFactsheetQuery(opts FactsheetOptions) string {
	params := url.Values{}
	if opts.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(opts.PageSize))
	}
	if opts.Cursor != "" {
		params.Set("next_cursor", opts.Cursor)
	}
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if !opts.StartDate.IsZero() {
		params.Set("start_date", opts.StartDate.Format("2006-01-02"))
	}
	if !opts.EndDate.IsZero() {
		params.Set("end_date", opts.EndDate.Format("2006-01-02"))
	}
	if opts.Latest {
		params.Set("latest", "true")
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

func buildOutstandingQuery(opts OutstandingOptions) string {
	params := url.Values{}
	if opts.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(opts.PageSize))
	}
	if opts.Cursor != "" {
		params.Set("next_cursor", opts.Cursor)
	}
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.StartPeriod != "" {
		params.Set("start_period", opts.StartPeriod)
	}
	if opts.EndPeriod != "" {
		params.Set("end_period", opts.EndPeriod)
	}

	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}
