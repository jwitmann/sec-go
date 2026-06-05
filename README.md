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

    sec "github.com/jwitmann/sec-go"
)

func main() {
    client, err := sec.NewClient(os.Getenv("SEC_API_KEY"))
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    amcs, _, err := client.ListAMCs(ctx, 10, "")
    if err != nil {
        log.Fatal(err)
    }
    for _, amc := range amcs {
        fmt.Println(amc.CompNameEN)
    }
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
    sec.WithCache(cache, 5*time.Minute),
)
```

## Supported Endpoints

### General Info
- `ListAMCs` — `/v2/fund/general-info/amcs`
- `GetFundProfiles` — `/v2/fund/general-info/profiles`
- `GetFundSpecifications` — `/v2/fund/general-info/specifications`
- `GetMutualFundFees` — `/v2/fund/general-info/mutual-fund-fees`
- `GetFundInvolveParties` — `/v2/fund/general-info/involve-parties`

### Daily Info
- `GetDailyNAV` — `/v2/fund/daily-info/nav`
- `GetDividendHistory` — `/v2/fund/daily-info/dividend-history`

### Factsheet
- `GetFactsheetFees` — `/v2/fund/factsheet/fees`
- `GetFactsheetPerformance` — `/v2/fund/factsheet/performance`
- `GetFactsheetSubscriptionRedemptionMinimums` — `/v2/fund/factsheet/subscription-redemption-minimums`
- `GetFactsheetSubscriptionRedemptionPeriods` — `/v2/fund/factsheet/subscription-redemption-periods`
- `GetFactsheetStatistics` — `/v2/fund/factsheet/statistics`
- `GetFactsheetDividendPolicy` — `/v2/fund/factsheet/dividend-policy`
- `GetFactsheetBenchmarks` — `/v2/fund/factsheet/benchmarks`
- `GetFundFactsheetURLs` — `/v2/fund/factsheet/urls`
- `GetFundIPOs` — `/v2/fund/factsheet/ipos`
- `GetAssetAllocation` — `/v2/fund/factsheet/asset-allocation`
- `GetRiskSpectrum` — `/v2/fund/factsheet/risk-spectrum`
- `GetTop5Holdings` — `/v2/fund/factsheet/top5-holdings`

### Outstanding
- `GetQuarterlyPortfolio` — `/v2/fund/outstanding/portfolio`
- `GetMonthlyPortfolioAssetType` — `/v2/fund/outstanding/portfolio-asset-type`

All endpoints return paginated results: `([]T, nextCursor, error)`. Use `FetchAllPages` to automatically traverse cursors.

## Pagination

```go
navs, err := sec.FetchAllPages(func(ctx context.Context, cursor string) ([]sec.DailyNAV, string, error) {
    return client.GetDailyNAV(ctx, sec.NAVOptions{
        ProjID:    "PRINCIPALi9",
        StartDate: start,
        EndDate:   end,
        Cursor:    cursor,
    })
})
```

## Batch Operations

Fetch NAV history for multiple funds concurrently with built-in rate limiting:

```go
results := sec.BatchGetNAVs(ctx, client, projIDs, startDate, endDate, sec.BatchNAVOptions{
    Concurrency: 4,
    Progress: func(completed, total int) {
        fmt.Printf("Progress: %d/%d\n", completed, total)
    },
})

for _, r := range results {
    if r.Err != nil {
        log.Printf("%s failed: %v", r.ProjID, r.Err)
        continue
    }
    fmt.Printf("%s: %d NAV records\n", r.ProjID, len(r.NAVs))
}
```

## Rate Limiting

The client enforces a minimum 16ms delay between requests to comply with SEC's rate limits (5,000 calls per 300 seconds). The rate limiter is thread-safe and respects context cancellation.

## Error Handling

```go
var (
    sec.ErrRateLimited  // HTTP 429
    sec.ErrNotFound     // HTTP 204
    sec.ErrUnauthorized // HTTP 401 / missing API key
)
```

Retry behavior:
- Retries on: 429, 500, 502, 503, 504, network errors
- Does not retry on: 400, 401, 403, 404
- Special handling for HTTP 421 with `Retry-After` header

## Examples

See `examples/`:
- `examples/basic/` — list AMCs and get NAV
- `examples/batch/` — fetch NAV history for multiple funds
- `examples/thaifa/` — THAIFA fallback integration pattern

## Testing

Run unit tests:
```bash
make test
```

Run all checks (format, lint, test, duplicate code):
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
├── client.go              # Core HTTP client
├── options.go             # Functional options
├── error.go               # Error types
├── rate.go                # Rate limiter
├── retry.go               # Retry logic
├── fund_service.go        # Fund API service methods
├── models.go              # V2 response models
├── pagination.go          # Pagination helpers
├── batch.go               # Batch/concurrent operations
├── client_test.go         # Client unit tests
├── fund_service_test.go   # Service method tests
├── pagination_test.go     # Pagination tests
├── batch_test.go          # Batch operation tests
├── internal/
│   ├── cache/             # In-memory cache
│   └── testutil/          # Test helpers
├── config/
│   └── sec-keys.example.json  # API key config template
├── docs/
│   └── v2-schemas/        # Sample responses + API.md
├── examples/              # Usage examples
└── Makefile               # Build tasks
```

## License

MIT
