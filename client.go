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

type (
	RequestHook  func(req *http.Request)
	ResponseHook func(req *http.Request, resp *http.Response, err error)
)

type logger interface {
	Printf(format string, v ...any)
}

type noopLogger struct{}

func (noopLogger) Printf(format string, v ...any) {}

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
	logger       logger
	requestHook  RequestHook
	responseHook ResponseHook
	language     Language
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
		logger:      noopLogger{},
		language:    LanguageEnglish,
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

func (c *Client) Language() Language {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.language
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

	c.logger.Printf("[sec-go] %s %s", method, url)
	if c.requestHook != nil {
		c.requestHook(req)
	}

	resp, err := doWithRetry(ctx, c.retryCfg, func() (*http.Response, error) {
		return c.httpClient.Do(req)
	})

	if c.responseHook != nil {
		c.responseHook(req, resp, err)
	}

	if err != nil {
		c.logger.Printf("[sec-go] %s %s error: %v", method, url, err)
	} else {
		c.logger.Printf("[sec-go] %s %s -> %d", method, url, resp.StatusCode)
	}

	return resp, err
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

// auto-translate helpers: when client language is English, translate Thai
// descriptions in-place so callers receive English results transparently.

func (c *Client) autoTranslateFees(fees []MutualFundFee) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllFees(fees, true)
}

func (c *Client) autoTranslateFactsheetFees(fees []FactsheetFee) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllFactsheetFees(fees, true)
}

func (c *Client) autoTranslateAssetAllocations(allocs []AssetAllocation) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllAssetAllocations(allocs, true)
}

func (c *Client) autoTranslateTop5Holdings(holdings []Top5Holding) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllTop5Holdings(holdings, true)
}

func (c *Client) autoTranslateQuarterlyPortfolios(items []QuarterlyPortfolio) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllQuarterlyPortfolios(items, true)
}

func (c *Client) autoTranslateMonthlyPortfolioAssetTypes(items []MonthlyPortfolioAssetType) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllMonthlyPortfolioAssetTypes(items, true)
}

func (c *Client) autoTranslateSubscriptionRedemptionPeriods(periods []FactsheetSubscriptionRedemptionPeriod) {
	if c.language != LanguageEnglish {
		return
	}
	TranslateAllSubscriptionRedemptionPeriods(periods, true)
}
