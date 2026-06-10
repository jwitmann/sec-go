package sec

import (
	"context"
	"fmt"
	"strings"
)

// FetchAllBondIssuers returns all bond issuers across every page.
func (c *Client) FetchAllBondIssuers(ctx context.Context) ([]BondIssuer, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]BondIssuer, string, error) {
		return c.ListBondIssuers(ctx, BondIssuerOptions{PageSize: 100, Cursor: cursor})
	})
}

// FindBondIssuer searches bond issuers by Thai name, English name, or unique_id.
func (c *Client) FindBondIssuer(ctx context.Context, query string) (*BondIssuer, error) {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil, fmt.Errorf("query is empty: %w", ErrNotFound)
	}

	cursor := ""
	for {
		issuers, next, err := c.ListBondIssuers(ctx, BondIssuerOptions{PageSize: 100, Cursor: cursor})
		if err != nil {
			return nil, err
		}

		for _, issuer := range issuers {
			if strings.EqualFold(issuer.UniqueID, query) {
				return &issuer, nil
			}
			if strings.Contains(strings.ToLower(issuer.CompNameTH), query) {
				return &issuer, nil
			}
			if strings.Contains(strings.ToLower(issuer.CompNameEN), query) {
				return &issuer, nil
			}
		}

		if next == "" {
			break
		}
		cursor = next
	}

	return nil, fmt.Errorf("bond issuer not found: %s: %w", query, ErrNotFound)
}

// FetchAllBondFeatures returns all bond features across every page.
func (c *Client) FetchAllBondFeatures(ctx context.Context, opts BondFeatureOptions) ([]BondFeature, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]BondFeature, string, error) {
		opts.PageSize = 100
		opts.Cursor = cursor
		return c.GetBondFeatures(ctx, opts)
	})
}
