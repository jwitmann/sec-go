package sec

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultBaseDelay = 100 * time.Millisecond
	defaultMaxDelay  = 5 * time.Second
)

type retryConfig struct {
	baseDelay  time.Duration
	maxDelay   time.Duration
	maxRetries int
}

func defaultRetryConfig() retryConfig {
	return retryConfig{
		baseDelay:  defaultBaseDelay,
		maxDelay:   defaultMaxDelay,
		maxRetries: defaultMaxRetries,
	}
}

func (r *retryConfig) sleep(ctx context.Context, attempt int) error {
	if attempt <= 0 {
		return nil
	}

	delay := r.baseDelay * time.Duration(1<<uint(attempt-1))
	if delay > r.maxDelay {
		delay = r.maxDelay
	}

	jitter := time.Duration(rand.Int63n(int64(delay) / 2))
	delay = delay/2 + jitter

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func shouldRetry(err error, statusCode int) bool {
	if err != nil {
		return true
	}
	return IsRetryable(statusCode)
}

func doWithRetry(ctx context.Context, rc retryConfig, do func() (*http.Response, error)) (*http.Response, error) {
	var lastErr error
	var lastStatusCode int

	for attempt := 0; attempt <= rc.maxRetries; attempt++ {
		if attempt > 0 {
			if err := rc.sleep(ctx, attempt); err != nil {
				return nil, err
			}
		}

		resp, err := do()
		if err != nil {
			lastErr = err
			if !shouldRetry(err, 0) {
				return nil, err
			}
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		lastStatusCode = resp.StatusCode
		lastErr = &APIError{StatusCode: resp.StatusCode, Message: resp.Status}

		if !shouldRetry(nil, resp.StatusCode) {
			return nil, lastErr
		}

		resp.Body.Close()
	}

	if lastErr != nil {
		return nil, lastErr
	}

	return nil, &APIError{StatusCode: lastStatusCode, Message: http.StatusText(lastStatusCode)}
}
