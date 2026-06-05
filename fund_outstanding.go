package sec

import (
	"context"
	"net/url"
	"strconv"
)

type OutstandingOptions struct {
	PageSize    int
	Cursor      string
	ProjID      string
	StartPeriod string
	EndPeriod   string
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

func (c *Client) GetQuarterlyPortfolio(ctx context.Context, opts OutstandingOptions) ([]QuarterlyPortfolio, string, error) {
	path := "/v2/fund/outstanding/portfolio" + buildOutstandingQuery(opts)
	items, cursor, err := fetchPaginated[QuarterlyPortfolio](ctx, c, path, "get quarterly portfolio")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateQuarterlyPortfolios(items)
	return items, cursor, nil
}

func (c *Client) GetMonthlyPortfolioAssetType(ctx context.Context, opts OutstandingOptions) ([]MonthlyPortfolioAssetType, string, error) {
	path := "/v2/fund/outstanding/portfolio-asset-type" + buildOutstandingQuery(opts)
	items, cursor, err := fetchPaginated[MonthlyPortfolioAssetType](ctx, c, path, "get monthly portfolio asset type")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateMonthlyPortfolioAssetTypes(items)
	return items, cursor, nil
}
