package sec

import (
	"context"
	"net/url"
)

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
	fees, cursor, err := fetchPaginated[MutualFundFee](ctx, c, path, "get mutual fund fees")
	if err != nil {
		return nil, "", err
	}
	c.autoTranslateFees(fees)
	return fees, cursor, nil
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

type FeeOptions struct {
	PageSize      int
	Cursor        string
	ProjID        string
	FundClassName string
}
