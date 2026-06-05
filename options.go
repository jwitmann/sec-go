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

func WithLogger(l logger) Option {
	return func(c *Client) {
		c.logger = l
	}
}

func WithRequestHook(hook RequestHook) Option {
	return func(c *Client) {
		c.requestHook = hook
	}
}

func WithResponseHook(hook ResponseHook) Option {
	return func(c *Client) {
		c.responseHook = hook
	}
}

func WithLanguage(lang Language) Option {
	return func(c *Client) {
		c.language = normalizeLanguage(lang)
	}
}
