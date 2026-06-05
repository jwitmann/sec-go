package sec

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BatchNAVResult struct {
	ProjID string
	NAVs   []DailyNAV
	Err    error
}

type BatchNAVOptions struct {
	Concurrency int
	Progress    func(completed, total int)
}

func BatchGetNAVs(ctx context.Context, c *Client, projIDs []string, startDate, endDate time.Time, opts BatchNAVOptions) []BatchNAVResult {
	if opts.Concurrency <= 0 {
		opts.Concurrency = 4
	}

	total := len(projIDs)
	results := make([]BatchNAVResult, total)
	work := make(chan int, total)
	var completed int
	var mu sync.Mutex

	for i := range projIDs {
		work <- i
	}
	close(work)

	var wg sync.WaitGroup
	for i := 0; i < opts.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range work {
				select {
				case <-ctx.Done():
					results[idx] = BatchNAVResult{
						ProjID: projIDs[idx],
						Err:    ctx.Err(),
					}
					continue
				default:
				}

				navs, _, err := c.GetDailyNAV(ctx, NAVOptions{
					ProjID:    projIDs[idx],
					StartDate: startDate,
					EndDate:   endDate,
				})
				results[idx] = BatchNAVResult{
					ProjID: projIDs[idx],
					NAVs:   navs,
					Err:    err,
				}

				mu.Lock()
				completed++
				if opts.Progress != nil {
					opts.Progress(completed, total)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return results
}

type BatchFundProfileResult struct {
	ProjID  string
	Profile *FundProfile
	Err     error
}

type BatchProfileOptions struct {
	Concurrency int
	Progress    func(completed, total int)
}

func BatchGetFundProfiles(ctx context.Context, c *Client, projIDs []string, opts BatchProfileOptions) []BatchFundProfileResult {
	if opts.Concurrency <= 0 {
		opts.Concurrency = 4
	}

	total := len(projIDs)
	results := make([]BatchFundProfileResult, total)
	work := make(chan int, total)
	var completed int
	var mu sync.Mutex

	for i := range projIDs {
		work <- i
	}
	close(work)

	var wg sync.WaitGroup
	for i := 0; i < opts.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range work {
				select {
				case <-ctx.Done():
					results[idx] = BatchFundProfileResult{
						ProjID: projIDs[idx],
						Err:    ctx.Err(),
					}
					continue
				default:
				}

				profiles, _, err := c.GetFundProfiles(ctx, ProfileOptions{
					ProjectInfo: projIDs[idx],
				})
				if err != nil {
					results[idx] = BatchFundProfileResult{
						ProjID: projIDs[idx],
						Err:    fmt.Errorf("fetch profile: %w", err),
					}
					mu.Lock()
					completed++
					if opts.Progress != nil {
						opts.Progress(completed, total)
					}
					mu.Unlock()
					continue
				}

				var profile *FundProfile
				if len(profiles) > 0 {
					profile = &profiles[0]
				}

				results[idx] = BatchFundProfileResult{
					ProjID:  projIDs[idx],
					Profile: profile,
					Err:     nil,
				}

				mu.Lock()
				completed++
				if opts.Progress != nil {
					opts.Progress(completed, total)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return results
}
