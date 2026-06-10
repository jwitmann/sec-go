package sec

import (
	"context"
	"net/url"
)

// LicenseCheckResult represents a generic license check result.
// Fields vary by license type, so we use a flexible structure.
type LicenseCheckResult struct {
	LicenseID     string   `json:"license_id"`
	LicenseNo     string   `json:"license_no"`
	LicenseType   string   `json:"license_type"`
	EntityNameTH  string   `json:"entity_name_th"`
	EntityNameEN  string   `json:"entity_name_en"`
	LicenseStatus string   `json:"license_status"`
	IssueDate     string   `json:"issue_date"`
	ExpireDate    string   `json:"expire_date"`
	LicenseDetail string   `json:"license_detail"`
	Remark        string   `json:"remark"`
	LastUpdDate   DateTime `json:"last_upd_date"`
}

// LicenseCheckOptions defines query options for license check endpoints.
type LicenseCheckOptions struct {
	PageSize int
	Cursor   string
	// Optional filters
	LicenseNo     string
	EntityName    string
	LicenseStatus string
}

// BondSaleRepOptions defines query options for bond sale representatives.
type BondSaleRepOptions struct {
	PageSize  int
	Cursor    string
	LicenseNo string
	RepName   string
}

// SecuritiesCompanyOptions defines query options for securities companies.
type SecuritiesCompanyOptions struct {
	PageSize    int
	Cursor      string
	LicenseNo   string
	CompanyName string
}

// checkLicense is a helper to call license check endpoints.
func (c *Client) checkLicense(ctx context.Context, path string, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.LicenseNo != "" {
		params.Set("license_no", opts.LicenseNo)
	}
	if opts.EntityName != "" {
		params.Set("entity_name", opts.EntityName)
	}
	if opts.LicenseStatus != "" {
		params.Set("license_status", opts.LicenseStatus)
	}
	fullPath := buildPath(path, params)
	return fetchPaginated[LicenseCheckResult](ctx, c, fullPath, "license check")
}

// CheckBondSaleReps checks bond sale representative licenses.
func (c *Client) CheckBondSaleReps(ctx context.Context, opts BondSaleRepOptions) ([]LicenseCheckResult, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.LicenseNo != "" {
		params.Set("license_no", opts.LicenseNo)
	}
	if opts.RepName != "" {
		params.Set("rep_name", opts.RepName)
	}
	path := buildPath("/v1/license-check/bond-sale-rep", params)
	return fetchPaginated[LicenseCheckResult](ctx, c, path, "check bond sale reps")
}

// CheckSecuritiesCompanies checks securities company licenses.
func (c *Client) CheckSecuritiesCompanies(ctx context.Context, opts SecuritiesCompanyOptions) ([]LicenseCheckResult, string, error) {
	params := url.Values{}
	setPagination(params, opts.PageSize, opts.Cursor)
	if opts.LicenseNo != "" {
		params.Set("license_no", opts.LicenseNo)
	}
	if opts.CompanyName != "" {
		params.Set("company_name", opts.CompanyName)
	}
	path := buildPath("/v1/license-check/securities-company", params)
	return fetchPaginated[LicenseCheckResult](ctx, c, path, "check securities companies")
}

// CheckDerivativesCompanies checks derivatives company licenses.
func (c *Client) CheckDerivativesCompanies(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/derivatives-company", opts)
}

// CheckSecuritiesBrokers checks securities broker licenses.
func (c *Client) CheckSecuritiesBrokers(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/securities-broker", opts)
}

// CheckDerivativesBrokers checks derivatives broker licenses.
func (c *Client) CheckDerivativesBrokers(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/derivatives-broker", opts)
}

// CheckInvestmentAdvisors checks investment advisor licenses.
func (c *Client) CheckInvestmentAdvisors(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/investment-advisor", opts)
}

// CheckSecuritiesFundManagers checks securities fund manager licenses.
func (c *Client) CheckSecuritiesFundManagers(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/securities-fund-manager", opts)
}

// CheckFundSupervisors checks fund supervisor licenses.
func (c *Client) CheckFundSupervisors(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/fund-supervisor", opts)
}

// CheckAuditors checks auditor licenses.
func (c *Client) CheckAuditors(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/auditor", opts)
}

// CheckCreditRatingCompanies checks credit rating company licenses.
func (c *Client) CheckCreditRatingCompanies(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/credit-rating-company", opts)
}

// CheckPrivateFunds checks private fund licenses.
func (c *Client) CheckPrivateFunds(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/private-fund", opts)
}

// CheckDerivativesFundManagers checks derivatives fund manager licenses.
func (c *Client) CheckDerivativesFundManagers(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/derivatives-fund-manager", opts)
}

// CheckSecuritiesBorrowingLending checks securities borrowing & lending licenses.
func (c *Client) CheckSecuritiesBorrowingLending(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/securities-borrowing-lending", opts)
}

// CheckFinancialAdvisors checks financial advisor licenses.
func (c *Client) CheckFinancialAdvisors(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/financial-advisor", opts)
}

// CheckAcquirers checks acquirer licenses.
func (c *Client) CheckAcquirers(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/acquirer", opts)
}

// CheckVentureCapitals checks venture capital licenses.
func (c *Client) CheckVentureCapitals(ctx context.Context, opts LicenseCheckOptions) ([]LicenseCheckResult, string, error) {
	return c.checkLicense(ctx, "/v1/license-check/venture-capital", opts)
}
