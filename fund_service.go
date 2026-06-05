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
