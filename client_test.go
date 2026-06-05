package sec

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Fatal("expected error for empty API key")
	}

	c, err := NewClient("test-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.primaryKey != "test-key" {
		t.Errorf("expected primary key 'test-key', got %q", c.primaryKey)
	}
}

func TestClientFromEnv(t *testing.T) {
	os.Unsetenv("SEC_API_KEY")
	_, err := ClientFromEnv()
	if err == nil {
		t.Fatal("expected error when SEC_API_KEY not set")
	}

	os.Setenv("SEC_API_KEY", "env-key")
	defer os.Unsetenv("SEC_API_KEY")

	c, err := ClientFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.primaryKey != "env-key" {
		t.Errorf("expected key 'env-key', got %q", c.primaryKey)
	}
}

func TestClientKeySwitching(t *testing.T) {
	c, _ := NewClient("primary", WithSecondaryKey("secondary"))

	if c.currentKey() != "primary" {
		t.Error("expected primary key by default")
	}

	c.UseSecondaryKey()
	if c.currentKey() != "secondary" {
		t.Error("expected secondary key")
	}

	c.UsePrimaryKey()
	if c.currentKey() != "primary" {
		t.Error("expected primary key after switch")
	}
}

func TestClientAuthHeader(t *testing.T) {
	var receivedKey string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedKey = r.Header.Get("Ocp-Apim-Subscription-Key")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	c, _ := NewClient("test-key", WithBaseURL(server.URL))
	ctx := context.Background()
	_, _ = c.get(ctx, "/test")

	if receivedKey != "test-key" {
		t.Errorf("expected auth header 'test-key', got %q", receivedKey)
	}
}

func TestClientRateLimiting(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	start := time.Now()
	_, _ = c.get(ctx, "/test1")
	_, _ = c.get(ctx, "/test2")
	elapsed := time.Since(start)

	if requestCount != 2 {
		t.Errorf("expected 2 requests, got %d", requestCount)
	}
	if elapsed < 10*time.Millisecond {
		t.Error("rate limiting not enforced")
	}
}

func TestClientRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL), WithRetryDelay(10*time.Millisecond))
	ctx := context.Background()
	_, err := c.get(ctx, "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

func TestClientRetryExhausted(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL), WithRetryDelay(10*time.Millisecond), WithMaxRetries(2))
	ctx := context.Background()
	_, err := c.get(ctx, "/test")

	if err == nil {
		t.Fatal("expected error after retries exhausted")
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts (1 initial + 2 retries), got %d", attempts)
	}
}

func TestClientNoRetryOn4xx(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()
	_, err := c.get(ctx, "/test")

	if err == nil {
		t.Fatal("expected error")
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt (no retry on 4xx), got %d", attempts)
	}
}

func TestClientNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()
	_, err := c.get(ctx, "/test")

	if err == nil {
		t.Fatal("expected error for 204")
	}
	if !isNotFound(err) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func isNotFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == (ErrNotFound.Error() + ": /test")
}

func TestClientOptions(t *testing.T) {
	customHTTP := &http.Client{Timeout: 60 * time.Second}
	c, _ := NewClient(
		"key",
		WithHTTPClient(customHTTP),
		WithBaseURL("https://custom.api"),
		WithSecondaryKey("secondary"),
		WithMaxRetries(5),
		WithRetryDelay(200*time.Millisecond),
		WithTimeout(45*time.Second),
	)

	if c.httpClient != customHTTP {
		t.Error("custom HTTP client not set")
	}
	if c.baseURL != "https://custom.api" {
		t.Errorf("expected base URL 'https://custom.api', got %q", c.baseURL)
	}
	if c.secondaryKey != "secondary" {
		t.Error("secondary key not set")
	}
	if c.retryCfg.maxRetries != 5 {
		t.Errorf("expected max retries 5, got %d", c.retryCfg.maxRetries)
	}
	if c.retryCfg.baseDelay != 200*time.Millisecond {
		t.Errorf("expected retry delay 200ms, got %v", c.retryCfg.baseDelay)
	}
	if c.httpClient.Timeout != 45*time.Second {
		t.Errorf("expected timeout 45s, got %v", c.httpClient.Timeout)
	}
}

func TestClientPost(t *testing.T) {
	var receivedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"123"}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()
	payload := []byte(`{"name":"test"}`)
	resp, err := c.post(ctx, "/test", payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(receivedBody) != `{"name":"test"}` {
		t.Errorf("unexpected body: %s", string(receivedBody))
	}
	if string(resp) != `{"id":"123"}` {
		t.Errorf("unexpected response: %s", string(resp))
	}
}
