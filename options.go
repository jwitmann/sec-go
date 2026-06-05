package sec

import (
	"net/http"
	"time"
)

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithSecondaryKey(key string) Option {
	return func(c *Client) {
		c.secondaryKey = key
	}
}

func WithRateLimiter(rl RateLimiter) Option {
	return func(c *Client) {
		c.rateLimiter = rl
	}
}

func WithMaxRetries(n int) Option {
	return func(c *Client) {
		c.retryCfg.maxRetries = n
	}
}

func WithRetryDelay(d time.Duration) Option {
	return func(c *Client) {
		c.retryCfg.baseDelay = d
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

func WithCache(cache cacheClient, ttl time.Duration) Option {
	return func(c *Client) {
		c.cache = cache
		c.cacheEnabled = true
		c.cacheTTL = ttl
	}
}
