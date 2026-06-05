package sec

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// SearchFunds searches across proj_id, proj_name_th, proj_name_en, and proj_abbr_name.
func (c *Client) SearchFunds(ctx context.Context, query string) ([]FundProfile, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]FundProfile, string, error) {
		return c.GetFundProfiles(ctx, ProfileOptions{
			ProjectInfo: query,
			PageSize:    100,
			Cursor:      cursor,
		})
	})
}

// GetFundsByCompany returns all fund profiles for a given AMC name or unique_id.
func (c *Client) GetFundsByCompany(ctx context.Context, companyName string) ([]FundProfile, error) {
	return FetchAllPages(ctx, func(ctx context.Context, cursor string) ([]FundProfile, string, error) {
		return c.GetFundProfiles(ctx, ProfileOptions{
			CompanyInfo: companyName,
			PageSize:    100,
			Cursor:      cursor,
		})
	})
}

// FindAMC searches AMCs by Thai name, English name, or unique_id.
func (c *Client) FindAMC(ctx context.Context, query string) (*AMC, error) {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil, fmt.Errorf("query is empty: %w", ErrNotFound)
	}

	cursor := ""
	for {
		amcs, next, err := c.ListAMCs(ctx, 100, cursor)
		if err != nil {
			return nil, err
		}

		for _, amc := range amcs {
			if strings.EqualFold(amc.UniqueID, query) {
				return &amc, nil
			}
			if strings.Contains(strings.ToLower(amc.CompNameTH), query) {
				return &amc, nil
			}
			if strings.Contains(strings.ToLower(amc.CompNameEN), query) {
				return &amc, nil
			}
		}

		if next == "" {
			break
		}
		cursor = next
	}

	return nil, fmt.Errorf("AMC not found: %s: %w", query, ErrNotFound)
}

// GetFundProfile returns the latest profile for a single fund.
func (c *Client) GetFundProfile(ctx context.Context, projID string) (*FundProfile, error) {
	profiles, _, err := c.GetFundProfiles(ctx, ProfileOptions{
		ProjectInfo: projID,
		PageSize:    1,
	})
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, fmt.Errorf("fund profile not found: %s: %w", projID, ErrNotFound)
	}
	return &profiles[0], nil
}

// GetFundLatestNAV returns the most recent NAV for a single fund.
func (c *Client) GetFundLatestNAV(ctx context.Context, projID string) (*DailyNAV, error) {
	navs, _, err := c.GetDailyNAV(ctx, NAVOptions{
		ProjID:   projID,
		PageSize: 1,
	})
	if err != nil {
		return nil, err
	}
	if len(navs) == 0 {
		return nil, fmt.Errorf("NAV not found: %s: %w", projID, ErrNotFound)
	}
	return &navs[0], nil
}

// GetFundRiskSpectrum returns the latest risk spectrum for a single fund.
func (c *Client) GetFundRiskSpectrum(ctx context.Context, projID string) (*RiskSpectrum, error) {
	items, _, err := c.GetRiskSpectrum(ctx, FactsheetOptions{
		ProjID:   projID,
		Latest:   true,
		PageSize: 1,
	})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("risk spectrum not found: %s: %w", projID, ErrNotFound)
	}
	return &items[0], nil
}

// GetFundFactsheetFees returns the latest factsheet fees for a single fund.
func (c *Client) GetFundFactsheetFees(ctx context.Context, projID string) ([]FactsheetFee, error) {
	fees, _, err := c.GetFactsheetFees(ctx, FactsheetOptions{
		ProjID:   projID,
		Latest:   true,
		PageSize: 100,
	})
	if err != nil {
		return nil, err
	}
	if len(fees) == 0 {
		return nil, fmt.Errorf("factsheet fees not found: %s: %w", projID, ErrNotFound)
	}
	return fees, nil
}

// GetFundAssetAllocation returns the latest asset allocation for a single fund.
func (c *Client) GetFundAssetAllocation(ctx context.Context, projID string) ([]AssetAllocation, error) {
	items, _, err := c.GetAssetAllocation(ctx, FactsheetOptions{
		ProjID:   projID,
		Latest:   true,
		PageSize: 100,
	})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("asset allocation not found: %s: %w", projID, ErrNotFound)
	}
	return items, nil
}

// GetFundTop5Holdings returns the latest top 5 holdings for a single fund.
func (c *Client) GetFundTop5Holdings(ctx context.Context, projID string) ([]Top5Holding, error) {
	items, _, err := c.GetTop5Holdings(ctx, FactsheetOptions{
		ProjID:   projID,
		Latest:   true,
		PageSize: 10,
	})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("top 5 holdings not found: %s: %w", projID, ErrNotFound)
	}
	return items, nil
}

// FundPortfolioView aggregates the latest portfolio-related data for a single fund.
type FundPortfolioView struct {
	ProjID                string
	AssetAllocation       []AssetAllocation
	Top5Holdings          []Top5Holding
	QuarterlyPortfolio    []QuarterlyPortfolio
	MonthlyAssetBreakdown []MonthlyPortfolioAssetType
}

// GetFundPortfolio fetches asset allocation, top 5 holdings, quarterly portfolio,
// and monthly asset breakdown concurrently for a single fund.
func (c *Client) GetFundPortfolio(ctx context.Context, projID string) (*FundPortfolioView, error) {
	if projID == "" {
		return nil, fmt.Errorf("proj_id is required")
	}

	view := &FundPortfolioView{ProjID: projID}

	type result struct {
		name string
		err  error
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	results := make(chan result, 4)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fetch := func(name string, fn func() error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := fn(); err != nil {
				select {
				case results <- result{name: name, err: err}:
				case <-ctx.Done():
				}
			}
		}()
	}

	fetch("asset allocation", func() error {
		items, _, err := c.GetAssetAllocation(ctx, FactsheetOptions{
			ProjID:   projID,
			Latest:   true,
			PageSize: 100,
		})
		if err != nil {
			return err
		}
		mu.Lock()
		view.AssetAllocation = items
		mu.Unlock()
		return nil
	})

	fetch("top 5 holdings", func() error {
		items, _, err := c.GetTop5Holdings(ctx, FactsheetOptions{
			ProjID:   projID,
			Latest:   true,
			PageSize: 10,
		})
		if err != nil {
			return err
		}
		mu.Lock()
		view.Top5Holdings = items
		mu.Unlock()
		return nil
	})

	fetch("quarterly portfolio", func() error {
		items, _, err := c.GetQuarterlyPortfolio(ctx, OutstandingOptions{
			ProjID:   projID,
			PageSize: 100,
		})
		if err != nil {
			return err
		}
		mu.Lock()
		view.QuarterlyPortfolio = items
		mu.Unlock()
		return nil
	})

	fetch("monthly asset breakdown", func() error {
		items, _, err := c.GetMonthlyPortfolioAssetType(ctx, OutstandingOptions{
			ProjID:   projID,
			PageSize: 100,
		})
		if err != nil {
			return err
		}
		mu.Lock()
		view.MonthlyAssetBreakdown = items
		mu.Unlock()
		return nil
	})

	go func() {
		wg.Wait()
		close(results)
	}()

	var errs []error
	for r := range results {
		errs = append(errs, fmt.Errorf("%s: %w", r.name, r.err))
	}
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return view, nil
}
