package sec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	defaultBaseURL    = "https://api.sec.or.th"
	defaultTimeout    = 30 * time.Second
	defaultMaxRetries = 3
	defaultRetryDelay = 100 * time.Millisecond
)

type cacheClient interface {
	Get(key string) ([]byte, bool)
	Set(key string, data []byte, ttl time.Duration)
}

type Client struct {
	httpClient   *http.Client
	baseURL      string
	primaryKey   string
	secondaryKey string
	useSecondary bool
	mu           sync.RWMutex
	rateLimiter  RateLimiter
	retryCfg     retryConfig
	cache        cacheClient
	cacheEnabled bool
	cacheTTL     time.Duration
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required: %w", ErrUnauthorized)
	}

	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL:     defaultBaseURL,
		primaryKey:  apiKey,
		rateLimiter: NewRateLimiter(),
		retryCfg:    defaultRetryConfig(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func ClientFromEnv(opts ...Option) (*Client, error) {
	key := getEnv("SEC_API_KEY", "")
	if key == "" {
		return nil, fmt.Errorf("SEC_API_KEY environment variable not set: %w", ErrUnauthorized)
	}
	return NewClient(key, opts...)
}

func (c *Client) UsePrimaryKey() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.useSecondary = false
}

func (c *Client) UseSecondaryKey() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.useSecondary = true
}

func (c *Client) currentKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.useSecondary && c.secondaryKey != "" {
		return c.secondaryKey
	}
	return c.primaryKey
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter: %w", err)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", c.currentKey())
	req.Header.Set("Content-Type", "application/json")

	return doWithRetry(ctx, c.retryCfg, func() (*http.Response, error) {
		return c.httpClient.Do(req)
	})
}

func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	return c.get(ctx, path)
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	if c.cacheEnabled && c.cache != nil {
		if data, ok := c.cache.Get(path); ok {
			return data, nil
		}
	}

	resp, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, path)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if c.cacheEnabled && c.cache != nil {
		c.cache.Set(path, data, c.cacheTTL)
	}

	return data, nil
}

func (c *Client) Post(ctx context.Context, path string, payload []byte) ([]byte, error) {
	return c.post(ctx, path, payload)
}

func (c *Client) post(ctx context.Context, path string, payload []byte) ([]byte, error) {
	resp, err := c.do(ctx, http.MethodPost, path, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, path)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return data, nil
}
