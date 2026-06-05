# SEC-Go

[![Go Reference](https://pkg.go.dev/badge/github.com/jwitmann/sec-go.svg)](https://pkg.go.dev/github.com/jwitmann/sec-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jwitmann/sec-go)](https://goreportcard.com/report/github.com/jwitmann/sec-go)

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
    sec.WithLogger(log.New(os.Stdout, "[sec] ", log.LstdFlags)),
    sec.WithRequestHook(func(req *http.Request) {
        // e.g., add tracing headers
    }),
    sec.WithResponseHook(func(req *http.Request, resp *http.Response, err error) {
        // e.g., record metrics
    }),
    sec.WithLanguage(sec.LanguageThai),
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

## Convenience Helpers

### Search by Company or Name

```go
// Get all funds managed by an AMC
profiles, err := client.GetFundsByCompany(ctx, "Krungthai Asset Management")

// Search across fund IDs, Thai names, English names, and abbreviations
profiles, err := client.SearchFunds(ctx, "alpha")

// Find an AMC by Thai/English name or unique ID
amc, err := client.FindAMC(ctx, "กรุงศรี")
```

### Single-Fund Lookups

```go
profile, err := client.GetFundProfile(ctx, "KT-Alpha")
nav, err := client.GetFundLatestNAV(ctx, "KT-Alpha")
rs, err := client.GetFundRiskSpectrum(ctx, "KT-Alpha")
fees, err := client.GetFundFactsheetFees(ctx, "KT-Alpha")
allocation, err := client.GetFundAssetAllocation(ctx, "KT-Alpha")
holdings, err := client.GetFundTop5Holdings(ctx, "KT-Alpha")
```

### Unified Portfolio View

Fetches asset allocation, top 5 holdings, quarterly portfolio, and monthly asset breakdown concurrently:

```go
portfolio, err := client.GetFundPortfolio(ctx, "KT-Alpha")
fmt.Println("Asset allocation:", portfolio.AssetAllocation)
fmt.Println("Top 5 holdings:", portfolio.Top5Holdings)
fmt.Println("Quarterly portfolio:", portfolio.QuarterlyPortfolio)
fmt.Println("Monthly asset breakdown:", portfolio.MonthlyAssetBreakdown)
```

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

## DateTime Handling

The SEC API returns datetime values in inconsistent formats (e.g., `2026-06-05T15:15:20.9` with fractional seconds but no timezone). The library uses a custom `sec.DateTime` type that transparently parses RFC3339, fractional-second timestamps, and plain dates (`YYYY-MM-DD`). You can use `.Time()` to get a standard `time.Time` value:

```go
fmt.Println(amc.LastUpdDate.Time())
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

## Language Support

The SEC API returns bilingual fields (e.g., `proj_name_th` + `proj_name_en`) and some Thai-only fields. Use `WithLanguage` to set a client-wide preference, then use helper methods to pick the right value:

```go
client, err := sec.NewClient("key", sec.WithLanguage(sec.LanguageThai))

profile := profiles[0]
fmt.Println(profile.Name(client.Language()))      // Thai or English fund name
fmt.Println(profile.CompanyName(client.Language())) // Thai or English AMC name
```

For Thai-only fields, use translation helpers (similar to finnomena-go):

```go
fees, _, _ := client.GetMutualFundFees(ctx, sec.FeeOptions{})
sec.TranslateAllFees(fees, true) // true = use English

for _, fee := range fees {
    fmt.Println(fee.FeeTypeDesc) // "Management Fee" instead of "ค่าธรรมเนียมการจัดการ"
}
```

Supported translations:
- `TranslateFee` / `TranslateAllFees` — fee types and units
- `TranslateFactsheetFee` / `TranslateAllFactsheetFees`
- `TranslateAssetAllocation` / `TranslateAllAssetAllocations`
- `TranslateTop5Holding` / `TranslateAllTop5Holdings`
- `TranslateQuarterlyPortfolio` / `TranslateAllQuarterlyPortfolios`
- `TranslateMonthlyPortfolioAssetType` / `TranslateAllMonthlyPortfolioAssetTypes`

Extend the public `FeeTypeTranslation`, `FeeUnitTranslation`, `AssetNameTranslation`, and `AssetLiabilityTranslation` maps to add more translations as you discover them.

## CLI Tool

A command-line tool is included for ad-hoc API queries:

```bash
go run ./cmd/sec-cli amcs
go run ./cmd/sec-cli profiles --company-info C0000000021
go run ./cmd/sec-cli nav --proj-id M0000_2552 --start 2024-01-01 --end 2024-01-31
go run ./cmd/sec-cli benchmarks --proj-id M0000_2552 --latest
go run ./cmd/sec-cli factsheet-urls --proj-id M0000_2552
```

The CLI reads API keys from `config/sec-keys.json` or falls back to the `SEC_API_KEY` environment variable.

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
├── fund_service.go        # Fund API service methods (all 21 endpoints)
├── models.go              # V2 response models (includes flexible DateTime parsing)
├── pagination.go          # Pagination helpers
├── batch.go               # Batch/concurrent operations
├── language.go            # Language preference + Thai↔English translation helpers
├── helpers_fund.go        # Convenience helpers: SearchFunds, GetFundsByCompany, etc.
├── client_test.go         # Client unit tests
├── fund_service_test.go   # Service method tests
├── pagination_test.go     # Pagination tests
├── batch_test.go          # Batch operation tests
├── language_test.go       # Language/translation tests
├── helpers_fund_test.go   # Convenience helper tests
├── integration_test.go    # Real API integration tests
├── internal/
│   ├── cache/             # In-memory cache
│   └── testutil/          # Test helpers
├── config/
│   └── sec-keys.example.json  # API key config template
├── cmd/
│   ├── sec-cli/           # CLI tool for ad-hoc queries
│   └── sec-lookup/        # Debug tool for fund lookups
├── docs/
│   └── v2-schemas/        # Sample responses + API.md
├── examples/              # Usage examples
└── Makefile               # Build tasks
```

## License

MIT
