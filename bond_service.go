package sec

import (
	"context"
	"net/url"
	"strconv"
)

// BondIssuer represents a bond issuer/AMC from the getAmcList endpoint.
type BondIssuer struct {
	UniqueID    string   `json:"unique_id"`
	CompNameTH  string   `json:"comp_name_th"`
	CompNameEN  string   `json:"comp_name_en"`
	LastUpdDate DateTime `json:"last_upd_date"`
}

// BondFeature represents bond specifications and features.
type BondFeature struct {
	ProjID         string   `json:"proj_id"`
	ProjNameTH     string   `json:"proj_name_th"`
	ProjNameEN     string   `json:"proj_name_en"`
	ProjAbbrName   string   `json:"proj_abbr_name"`
	IssueDate      string   `json:"issue_date"`
	MaturityDate   string   `json:"maturity_date"`
	CouponRate     float64  `json:"coupon_rate"`
	FaceValue      float64  `json:"face_value"`
	IssueValue     float64  `json:"issue_value"`
	BondType       string   `json:"bond_type"`
	BondSubType    string   `json:"bond_sub_type"`
	SecuredFlag    string   `json:"secured_flag"`
	GuaranteeFlag  string   `json:"guarantee_flag"`
	CollateralDesc string   `json:"collateral_desc"`
	PurposeTH      string   `json:"purpose_th"`
	PurposeEN      string   `json:"purpose_en"`
	RemarkTH       string   `json:"remark_th"`
	RemarkEN       string   `json:"remark_en"`
	LastUpdDate    DateTime `json:"last_upd_date"`
}

// BondCreditRating represents credit ratings for bonds by time period.
type BondCreditRating struct {
	ProjID       string   `json:"proj_id"`
	Period       int      `json:"period"`
	RatingAgency string   `json:"rating_agency"`
	Rating       string   `json:"rating"`
	Outlook      string   `json:"outlook"`
	AsOfDate     string   `json:"as_of_date"`
	LastUpdDate  DateTime `json:"last_upd_date"`
}

// BondOutstanding represents outstanding values by time period.
type BondOutstanding struct {
	ProjID         string  `json:"proj_id"`
	Period         int     `json:"period"`
	OutstandingQty float64 `json:"outstanding_qty"`
	OutstandingVal float64 `json:"outstanding_val"`
	HeldByInvestor float64 `json:"held_by_investor"`
	HeldByIssuer   float64 `json:"held_by_issuer"`
	LastUpdDate    string  `json:"last_upd_date"`
}

// BondRelatedParty represents related parties by time period.
type BondRelatedParty struct {
	ProjID      string   `json:"proj_id"`
	Period      int      `json:"period"`
	PartyType   string   `json:"party_type"`
	PartyNameTH string   `json:"party_name_th"`
	PartyNameEN string   `json:"party_name_en"`
	LastUpdDate DateTime `json:"last_upd_date"`
}

// BondInvestorHolding represents investor holdings by type.
type BondInvestorHolding struct {
	ProjID       string  `json:"proj_id"`
	Period       int     `json:"period"`
	InvestorType string  `json:"investor_type"`
	HoldingQty   float64 `json:"holding_qty"`
	HoldingVal   float64 `json:"holding_val"`
	HoldingPct   float64 `json:"holding_pct"`
	LastUpdDate  string  `json:"last_upd_date"`
}

// BondIssuerOptions defines query options for ListBondIssuers.
type BondIssuerOptions struct {
	PageSize int
	Cursor   string
}

// BondFeatureOptions defines query options for GetBondFeatures.
type BondFeatureOptions struct {
	PageSize int
	Cursor   string
	ProjID   string
}

// BondPeriodOptions defines query options for bond endpoints that filter by period.
type BondPeriodOptions struct {
	PageSize    int
	Cursor      string
	ProjID      string
	StartPeriod int
	EndPeriod   int
}

// ListBondIssuers returns the list of bond issuers (AMCs).
func (c *Client) ListBondIssuers(ctx context.Context, opts BondIssuerOptions) ([]BondIssuer, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	path := buildPath("/v2/bond/general-info/amcs", params)
	return fetchPaginated[BondIssuer](ctx, c, path, "list bond issuers")
}

// GetBondFeatures returns bond specifications and features.
func (c *Client) GetBondFeatures(ctx context.Context, opts BondFeatureOptions) ([]BondFeature, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	path := buildPath("/v2/bond/general-info/features", params)
	return fetchPaginated[BondFeature](ctx, c, path, "get bond features")
}

// GetBondCreditRatings returns credit ratings for bonds by time period.
func (c *Client) GetBondCreditRatings(ctx context.Context, opts BondPeriodOptions) ([]BondCreditRating, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.StartPeriod > 0 {
		params.Set("start_period", strconv.Itoa(opts.StartPeriod))
	}
	if opts.EndPeriod > 0 {
		params.Set("end_period", strconv.Itoa(opts.EndPeriod))
	}
	path := buildPath("/v2/bond/credit-rating", params)
	return fetchPaginated[BondCreditRating](ctx, c, path, "get bond credit ratings")
}

// GetBondOutstanding returns outstanding values by time period.
func (c *Client) GetBondOutstanding(ctx context.Context, opts BondPeriodOptions) ([]BondOutstanding, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.StartPeriod > 0 {
		params.Set("start_period", strconv.Itoa(opts.StartPeriod))
	}
	if opts.EndPeriod > 0 {
		params.Set("end_period", strconv.Itoa(opts.EndPeriod))
	}
	path := buildPath("/v2/bond/outstanding", params)
	return fetchPaginated[BondOutstanding](ctx, c, path, "get bond outstanding")
}

// GetBondRelatedParties returns related parties by time period.
func (c *Client) GetBondRelatedParties(ctx context.Context, opts BondPeriodOptions) ([]BondRelatedParty, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.StartPeriod > 0 {
		params.Set("start_period", strconv.Itoa(opts.StartPeriod))
	}
	if opts.EndPeriod > 0 {
		params.Set("end_period", strconv.Itoa(opts.EndPeriod))
	}
	path := buildPath("/v2/bond/related-party", params)
	return fetchPaginated[BondRelatedParty](ctx, c, path, "get bond related parties")
}

// GetBondInvestorHoldings returns investor holdings by type.
func (c *Client) GetBondInvestorHoldings(ctx context.Context, opts BondPeriodOptions) ([]BondInvestorHolding, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.ProjID != "" {
		params.Set("proj_id", opts.ProjID)
	}
	if opts.StartPeriod > 0 {
		params.Set("start_period", strconv.Itoa(opts.StartPeriod))
	}
	if opts.EndPeriod > 0 {
		params.Set("end_period", strconv.Itoa(opts.EndPeriod))
	}
	path := buildPath("/v2/bond/investor-holding", params)
	return fetchPaginated[BondInvestorHolding](ctx, c, path, "get bond investor holdings")
}

// setPagination and buildPath are defined in fund_service.go.
