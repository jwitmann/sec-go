package sec

import (
	"context"
	"fmt"
	"strings"
)

// FetchAllPVDs returns all provident funds across every page.
func (c *Client) FetchAllPVDs(ctx context.Context) ([]PVDListItem, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]PVDListItem, string, error) {
		return c.ListPVDs(ctx, PVDListOptions{PageSize: 100, Cursor: cursor})
	})
}

// FindPVD searches provident funds by Thai name, English name, or unique_id.
func (c *Client) FindPVD(ctx context.Context, query string) (*PVDListItem, error) {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil, fmt.Errorf("query is empty: %w", ErrNotFound)
	}

	cursor := ""
	for {
		pvds, next, err := c.ListPVDs(ctx, PVDListOptions{PageSize: 100, Cursor: cursor})
		if err != nil {
			return nil, err
		}

		for _, pvd := range pvds {
			if strings.EqualFold(pvd.UniqueID, query) {
				return &pvd, nil
			}
			if strings.Contains(strings.ToLower(pvd.CompNameTH), query) {
				return &pvd, nil
			}
			if strings.Contains(strings.ToLower(pvd.CompNameEN), query) {
				return &pvd, nil
			}
		}

		if next == "" {
			break
		}
		cursor = next
	}

	return nil, fmt.Errorf("PVD not found: %s: %w", query, ErrNotFound)
}

// FetchAllPVDFundInfo returns all PVD fund info across every page.
func (c *Client) FetchAllPVDFundInfo(ctx context.Context, opts PVDProjOptions) ([]PVDFundInfo, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]PVDFundInfo, string, error) {
		opts.PageSize = 100
		opts.Cursor = cursor
		return c.GetPVDFundInfo(ctx, opts)
	})
}
