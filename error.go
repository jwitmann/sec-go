package sec

import (
	"errors"
	"fmt"
)

var (
	ErrRateLimited  = errors.New("rate limited")
	ErrNotFound     = errors.New("data not found")
	ErrUnauthorized = errors.New("invalid API key")
)

type APIError struct {
	StatusCode int
	Message    string
	RawBody    []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

func IsRetryable(statusCode int) bool {
	switch statusCode {
	case 421, 429, 500, 502, 503, 504:
		return true
	default:
		return false
	}
}
