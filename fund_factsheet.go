package sec

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

type FactsheetOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	StartDate     time.Time
	EndDate       time.Time
	Latest        bool
	FundClassName string
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
	ipos, cursor, err := fetchPaginated[FundIPO](ctx, c, path, "get fund IPOs")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateFundIPOs(ipos)
	return ipos, cursor, nil
}

func (c *Client) GetFactsheetFees(ctx context.Context, opts FactsheetOptions) ([]FactsheetFee, string, error) {
	path := "/v2/fund/factsheet/fees" + buildFactsheetQuery(opts)
	fees, cursor, err := fetchPaginated[FactsheetFee](ctx, c, path, "get factsheet fees")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateFactsheetFees(fees)
	return fees, cursor, nil
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
	periods, cursor, err := fetchPaginated[FactsheetSubscriptionRedemptionPeriod](ctx, c, path, "get subscription-redemption periods")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateSubscriptionRedemptionPeriods(periods)
	return periods, cursor, nil
}

func (c *Client) GetFactsheetStatistics(ctx context.Context, opts FactsheetOptions) ([]FactsheetStatistics, string, error) {
	path := "/v2/fund/factsheet/statistics" + buildFactsheetQuery(opts)
	stats, cursor, err := fetchPaginated[FactsheetStatistics](ctx, c, path, "get factsheet statistics")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateFactsheetStatistics(stats)
	return stats, cursor, nil
}

func (c *Client) GetFactsheetDividendPolicy(ctx context.Context, opts FactsheetOptions) ([]FundDividendPolicy, string, error) {
	path := "/v2/fund/factsheet/dividend-policy" + buildFactsheetQuery(opts)
	return fetchPaginated[FundDividendPolicy](ctx, c, path, "get factsheet dividend policy")
}

func (c *Client) GetFactsheetBenchmarks(ctx context.Context, opts FactsheetOptions) ([]FactsheetBenchmark, string, error) {
	path := "/v2/fund/factsheet/benchmarks" + buildFactsheetQuery(opts)
	benchmarks, cursor, err := fetchPaginated[FactsheetBenchmark](ctx, c, path, "get factsheet benchmarks")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateFactsheetBenchmarks(benchmarks)
	return benchmarks, cursor, nil
}

func (c *Client) GetAssetAllocation(ctx context.Context, opts FactsheetOptions) ([]AssetAllocation, string, error) {
	path := "/v2/fund/factsheet/asset-allocation" + buildFactsheetQuery(opts)
	allocs, cursor, err := fetchPaginated[AssetAllocation](ctx, c, path, "get asset allocation")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateAssetAllocations(allocs)
	return allocs, cursor, nil
}

func (c *Client) GetRiskSpectrum(ctx context.Context, opts FactsheetOptions) ([]RiskSpectrum, string, error) {
	path := "/v2/fund/factsheet/risk-spectrum" + buildFactsheetQuery(opts)
	return fetchPaginated[RiskSpectrum](ctx, c, path, "get risk spectrum")
}

func (c *Client) GetTop5Holdings(ctx context.Context, opts FactsheetOptions) ([]Top5Holding, string, error) {
	path := "/v2/fund/factsheet/top5-holdings" + buildFactsheetQuery(opts)
	holdings, cursor, err := fetchPaginated[Top5Holding](ctx, c, path, "get top 5 holdings")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateTop5Holdings(holdings)
	return holdings, cursor, nil
}
