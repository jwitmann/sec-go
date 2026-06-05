package sec

import (
	"context"
	"sync"
	"time"
)

const defaultMinDelay = 100 * time.Millisecond

type RateLimiter interface {
	Wait(ctx context.Context) error
}

type simpleRateLimiter struct {
	mu       sync.Mutex
	lastReq  time.Time
	minDelay time.Duration
}

func NewRateLimiter() RateLimiter {
	return &simpleRateLimiter{
		minDelay: defaultMinDelay,
	}
}

func (r *simpleRateLimiter) Wait(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	elapsed := time.Since(r.lastReq)
	if elapsed < r.minDelay {
		delay := r.minDelay - elapsed
		timer := time.NewTimer(delay)
		defer timer.Stop()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
		}
	}

	r.lastReq = time.Now()
	return nil
}
