# SEC-Go

Go client library for the Thailand SEC OpenAPI V2 (`api.sec.or.th`).

## Installation

```bash
go get github.com/jwitmann/sec-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jwitmann/sec-go"
)

func main() {
    client, err := sec.NewClient(os.Getenv("SEC_API_KEY"))
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    // Use client...
}
```

## Authentication

Pass your API key directly to the constructor:

```go
client, err := sec.NewClient("your-api-key")
```

Or use the environment variable helper:

```go
client, err := sec.ClientFromEnv() // reads SEC_API_KEY
```

For primary/secondary key support:

```go
client, err := sec.NewClient(
    "primary-key",
    sec.WithSecondaryKey("secondary-key"),
)

client.UseSecondaryKey() // switch to secondary key
client.UsePrimaryKey()   // switch back to primary key
```

## Configuration Options

```go
client, err := sec.NewClient(
    "your-api-key",
    sec.WithTimeout(60*time.Second),
    sec.WithMaxRetries(5),
    sec.WithRetryDelay(200*time.Millisecond),
    sec.WithBaseURL("https://custom.api.sec.or.th"),
)
```

## Rate Limiting

The client enforces a minimum 100ms delay between requests to comply with SEC's rate limits (3,000 calls per 300 seconds). The rate limiter is thread-safe and respects context cancellation.

## Error Handling

```go
var (
    sec.ErrRateLimited  // HTTP 429
    sec.ErrNotFound     // HTTP 204
    sec.ErrUnauthorized // HTTP 401
)
```

Retry behavior:
- Retries on: 429, 500, 502, 503, 504, network errors
- Does not retry on: 400, 401, 403, 404

## Testing

Run unit tests:
```bash
make test
```

Run all checks (format, lint, test):
```bash
make check
```

For integration tests (requires real API key in `config/sec-keys.json`):
```bash
make test-integration
```

## Project Structure

```
sec-go/
├── client.go          # Core HTTP client
├── options.go         # Functional options
├── error.go           # Error types
├── rate.go            # Rate limiter
├── retry.go           # Retry logic
├── client_test.go     # Unit tests
├── internal/
│   └── testutil/      # Test helpers
├── config/
│   └── sec-keys.example.json  # API key config template
└── Makefile           # Build tasks
```

## License

MIT
