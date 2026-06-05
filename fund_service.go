package sec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func (c *Client) ListAMCs(ctx context.Context, pageSize int, cursor string) ([]AMC, string, error) {
	params := url.Values{}
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if cursor != "" {
		params.Set("next_cursor", cursor)
	}

	path := "/v2/fund/general-info/amcs"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("list AMCs: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []AMC `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal AMCs: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetFundProfiles(ctx context.Context, opts ProfileOptions) ([]FundProfile, string, error) {
	params := url.Values{}
	if opts.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(opts.PageSize))
	}
	if opts.Cursor != "" {
		params.Set("next_cursor", opts.Cursor)
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

	path := "/v2/fund/general-info/profiles"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get fund profiles: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []FundProfile `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal profiles: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

type ProfileOptions struct {
	PageSize      int
	Cursor        string
	FundClassName string
	FundStatus    string
	ProjectInfo   string
	CompanyInfo   string
}

func (c *Client) GetDailyNAV(ctx context.Context, opts NAVOptions) ([]DailyNAV, string, error) {
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
		params.Set("start_nav_date", opts.StartDate.Format("2006-01-02"))
	}
	if !opts.EndDate.IsZero() {
		params.Set("end_nav_date", opts.EndDate.Format("2006-01-02"))
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	path := "/v2/fund/daily-info/nav"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get daily NAV: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []DailyNAV `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal NAV: %w", err)
	}

	return response.Items, response.NextCursor, nil
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
	if opts.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(opts.PageSize))
	}
	if opts.Cursor != "" {
		params.Set("next_cursor", opts.Cursor)
	}
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.FundClassName != "" {
		params.Set("fund_class_name", opts.FundClassName)
	}

	path := "/v2/fund/general-info/mutual-fund-fees"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get mutual fund fees: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []MutualFundFee `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal fees: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetFactsheetFees(ctx context.Context, opts FactsheetOptions) ([]FactsheetFee, string, error) {
	path := "/v2/fund/factsheet/fees" + buildFactsheetQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get factsheet fees: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []FactsheetFee `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal factsheet fees: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetFactsheetPerformance(ctx context.Context, opts FactsheetOptions) ([]FactsheetPerformance, string, error) {
	path := "/v2/fund/factsheet/performance" + buildFactsheetQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get factsheet performance: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []FactsheetPerformance `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal performance: %w", err)
	}

	return response.Items, response.NextCursor, nil
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
		params.Set("start_dividend_date", opts.StartDate.Format("2006-01-02"))
	}
	if !opts.EndDate.IsZero() {
		params.Set("end_dividend_date", opts.EndDate.Format("2006-01-02"))
	}
	if opts.ClassAbbrName != "" {
		params.Set("class_abbr_name", opts.ClassAbbrName)
	}

	path := "/v2/fund/daily-info/dividend-history"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get dividend history: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []DividendHistory `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal dividend history: %w", err)
	}

	return response.Items, response.NextCursor, nil
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

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get asset allocation: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []AssetAllocation `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal asset allocation: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetRiskSpectrum(ctx context.Context, opts FactsheetOptions) ([]RiskSpectrum, string, error) {
	path := "/v2/fund/factsheet/risk-spectrum" + buildFactsheetQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get risk spectrum: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []RiskSpectrum `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal risk spectrum: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetTop5Holdings(ctx context.Context, opts FactsheetOptions) ([]Top5Holding, string, error) {
	path := "/v2/fund/factsheet/top5-holdings" + buildFactsheetQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get top 5 holdings: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []Top5Holding `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal top 5 holdings: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetQuarterlyPortfolio(ctx context.Context, opts OutstandingOptions) ([]QuarterlyPortfolio, string, error) {
	path := "/v2/fund/outstanding/portfolio" + buildOutstandingQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get quarterly portfolio: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []QuarterlyPortfolio `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal quarterly portfolio: %w", err)
	}

	return response.Items, response.NextCursor, nil
}

func (c *Client) GetMonthlyPortfolioAssetType(ctx context.Context, opts OutstandingOptions) ([]MonthlyPortfolioAssetType, string, error) {
	path := "/v2/fund/outstanding/portfolio-asset-type" + buildOutstandingQuery(opts)

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("get monthly portfolio asset type: %w", err)
	}

	var response struct {
		PaginatedResponse
		Items []MonthlyPortfolioAssetType `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal monthly portfolio asset type: %w", err)
	}

	return response.Items, response.NextCursor, nil
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
