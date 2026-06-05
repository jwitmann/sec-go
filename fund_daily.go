package sec

import (
	"context"
	"net/url"
	"time"
)

type NAVOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	FundClassName string
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

type DividendHistoryOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	ClassAbbrName string
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
