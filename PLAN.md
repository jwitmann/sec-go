# SEC-Go Development Plan

## Objective

Build a Go client library for the **Thailand SEC OpenAPI V2** (`api.sec.or.th`) — a complementary/fallback data source to Finnomena for Thai mutual fund data.

**Constraint:** V1 API is deprecated. All development targets V2 exclusively.

---

## Current Project Structure

```
sec-go/
├── client.go              # Core HTTP client, auth, retry, rate limiting
├── options.go             # Functional options
├── error.go               # Error types
├── rate.go                # Rate limiter
├── retry.go               # Retry logic
├── helpers.go             # Utility helpers
├── fund_service.go        # V2 Fund API methods (all 21 endpoints)
├── models.go              # V2 data models (includes flexible DateTime type)
├── pagination.go          # Pagination helper (FetchAllPages)
├── batch.go               # Concurrent batch operations
├── language.go            # Language preference + Thai↔English translation helpers
├── helpers_fund.go        # Convenience helpers: SearchFunds, GetFundsByCompany, etc.
│
├── client_test.go         # Client tests
├── fund_service_test.go   # Service method tests
├── pagination_test.go     # Pagination tests
├── batch_test.go          # Batch operation tests
├── language_test.go       # Language/translation tests
├── helpers_fund_test.go   # Convenience helper tests
├── integration_test.go    # Real API integration tests (build tag: integration)
│
├── internal/
│   ├── cache/             # Optional in-memory cache
│   └── testutil/          # Test helpers
│
├── config/
│   └── sec-keys.example.json  # API key config template
│
├── docs/
│   └── v2-schemas/        # Sample responses + API.md endpoint docs
│
├── examples/
│   ├── basic/             # List AMCs, get NAV
│   ├── batch/             # Fetch NAV range for multiple funds
│   └── thaifa/            # THAIFA fallback pattern
│
├── cmd/
│   ├── sec-cli/           # CLI tool for ad-hoc API queries
│   └── sec-lookup/        # Debug tool for fund lookups
│
├── go.mod
├── go.sum
├── Makefile
├── README.md
├── AGENTS.md
└── PLAN.md                # This document
```

---

## Phase 1: Generic Client Infrastructure ✅ COMPLETE

**Goal:** Build everything that doesn't depend on V2 schemas.

**Deliverables:**
- `client.go` — Core HTTP client with auth, retry, rate limiting
- `options.go` — Functional options (HTTP client, base URL, secondary key, rate limiter, retries, timeout, cache, logger, request hook)
- `error.go` — Custom error types (`APIError`, `ErrRateLimited`, `ErrNotFound`, `ErrUnauthorized`)
- `rate.go` — Thread-safe rate limiter with 60ms minimum delay
- `retry.go` — Exponential backoff with jitter, handles HTTP 421 + `Retry-After`
- `client_test.go` — Mock server tests for auth, rate limiting, retry, errors, options, POST
- `ClientFromEnv()` — Helper that reads `SEC_API_KEY` from environment

---

## Phase 2: V2 Schema Discovery ✅ COMPLETE

**Goal:** Discover actual V2 endpoints, request parameters, and response schemas.

**Key findings:**
- Base URL: `https://api.sec.or.th`
- Auth header: `Ocp-Apim-Subscription-Key`
- Correct NAV endpoint: `/v2/fund/daily-info/nav` (not `/v2/fund/daily-nav`)
- All working endpoints use cursor-based pagination with `message`, `page_size`, `next_cursor`, `items`
- Rate limit from SEC docs: 5,000 requests / 300 seconds, minimum 60ms between requests

**Deliverables:**
- Complete endpoint inventory in `docs/v2-schemas/API.md` (all 21 endpoints, sorted by category)
- Sample JSON responses in `docs/v2-schemas/` for discovered endpoints
- Test key infrastructure: `config/sec-keys.example.json` + `internal/testutil/keys.go`

---

## Phase 3: V2 Endpoint Implementation ✅ COMPLETE

**Goal:** Build typed Go methods for all 21 discovered V2 endpoints.

**Models added:** `AMC`, `FundProfile`, `FundSpecification`, `DailyNAV`, `MutualFundFee`, `FactsheetFee`, `FactsheetPerformance`, `DividendHistory`, `AssetAllocation`, `RiskSpectrum`, `Top5Holding`, `QuarterlyPortfolio`, `MonthlyPortfolioAssetType`, `FactsheetSubscriptionRedemptionMinimum`, `FactsheetSubscriptionRedemptionPeriod`, `FactsheetStatistics`, `FactsheetBenchmark`, `FundInvolveParty`, `FundFactsheetURL`, `FundIPO`, `FundDividendPolicy`.

**Service methods added (all 21 Fund API endpoints):**

General Info:
- `ListAMCs` — #1
- `GetFundProfiles` — #2
- `GetFundSpecifications` — #3
- `GetMutualFundFees` — #4
- `GetFundInvolveParties` — #5

Factsheet:
- `GetFundFactsheetURLs` — #6
- `GetFundIPOs` — #7
- `GetFactsheetBenchmarks` — #8
- `GetFactsheetSubscriptionRedemptionMinimums` — #9
- `GetFactsheetSubscriptionRedemptionPeriods` — #10
- `GetRiskSpectrum` — #11
- `GetFactsheetStatistics` — #12
- `GetFactsheetDividendPolicy` — #13
- `GetFactsheetFees` — #14
- `GetFactsheetPerformance` — #15
- `GetAssetAllocation` — #16
- `GetTop5Holdings` — #17

Outstanding:
- `GetQuarterlyPortfolio` — #18
- `GetMonthlyPortfolioAssetType` — #19

Daily Info:
- `GetDailyNAV` — #20
- `GetDividendHistory` — #21

**Infrastructure added:**
- Generic `FetchAllPages[T]` for cursor pagination
- `BatchGetNAVs` + `BatchGetFundProfiles` for concurrent fetching with progress callbacks
- Refactored `fund_service.go` using `fetchPaginated[T]`, `setPagination`, `setDateRange`, `buildPath` helpers
- Flexible `DateTime` type in `models.go` to handle SEC's inconsistent datetime formats (RFC3339, fractional seconds without timezone, plain dates)

**Convenience helpers (`helpers_fund.go`):**
- `SearchFunds(ctx, query)` — search across `proj_id`, names, and abbreviations
- `GetFundsByCompany(ctx, companyName)` — get all funds managed by an AMC
- `FindAMC(ctx, query)` — find an AMC by Thai/English name or `unique_id`
- `GetFundProfile(ctx, projID)` — single-fund profile lookup
- `GetFundLatestNAV(ctx, projID)` — latest NAV for one fund
- `GetFundRiskSpectrum(ctx, projID)` — latest risk spectrum
- `GetFundFactsheetFees(ctx, projID)` — latest factsheet fees
- `GetFundAssetAllocation(ctx, projID)` — latest asset allocation
- `GetFundTop5Holdings(ctx, projID)` — latest top 5 holdings
- `GetFundPortfolio(ctx, projID)` — unified portfolio view (concurrent fetch of 4 datasets)

**Language support (`language.go`):**
- `Language` type with `LanguageThai` / `LanguageEnglish`
- `WithLanguage(lang)` option and `Client.Language()` getter
- Bilingual helpers: `AMC.Name()`, `FundProfile.Name()`, `FundProfile.CompanyName()`, `FundInvolveParty.EntityName()`
- Thai→English translation maps and functions for fee types, units, asset names, and liability categories

**Tests added:**
- `fund_service_test.go` — mock server tests for all service methods + error cases
- `pagination_test.go` — pagination helper tests
- `batch_test.go` — batch operation tests (concurrency, progress, context cancellation)
- `language_test.go` — language normalization and translation tests
- `helpers_fund_test.go` — convenience helper tests (search, company lookup, portfolio view)

**All unit tests passing: 68 tests.**

---

## Phase 4: Polish, Documentation & Developer Tools ✅ COMPLETE

**Deliverables:**
- `README.md` — Installation, auth, all 21 endpoints, pagination, batch ops, convenience helpers, language support, examples, project structure
- `docs/v2-schemas/API.md` — Implementation-grade sorted endpoint inventory with parameters, response fields, example URLs, and linked sample JSON files
- `examples/basic/main.go` — list AMCs, get profiles, get NAV
- `examples/batch/main.go` — fetch NAV history for multiple funds concurrently
- `examples/thaifa/main.go` — THAIFA fallback integration pattern
- `cmd/sec-cli/main.go` — CLI tool for ad-hoc queries against the live API
- `cmd/sec-lookup/main.go` — Debug tool for fund lookups (e.g., TRAREEARTH investigation)
- `integration_test.go` — Integration tests gated by `//go:build integration`
- `Makefile` with `test`, `lint`, `format`, `check`, `test-integration` targets
- `AGENTS.md` with project guidelines

**Quality gates:**
- `golangci-lint` ✅
- `staticcheck` ✅
- `go vet` ✅
- `gocyclo -over 25` ✅
- `dupl -t 100` ✅ (tolerated clone groups in test setup patterns)
- `go test ./...` ✅ (68 tests)

---

## Phase 5: THAIFA Integration (Future)

**Not part of sec-go** — this is in the THAIFA repo.

**Planned integration pattern:**

```go
type PriceService struct {
    finnomenaClient *finnomena.Client
    secClient       *sec.Client
}

func (s *PriceService) GetPrices(fundID, shortCode string, from, to time.Time) (*models.BarsResponse, error) {
    bars, err := s.finnomenaClient.GetHistoricalPrices(shortCode, from, to)
    if err == nil && len(bars.Time) > 0 {
        return bars, nil
    }

    if s.secClient != nil {
        return s.fetchFromSEC(fundID, from, to)
    }

    return nil, err
}
```

**Open questions for THAIFA integration:**
- Mapping between Finnomena `fund_id`/`short_code` and SEC `proj_id`
- How to represent daily closing NAV as OHLCV bars (SEC provides close only)
- Cache invalidation strategy when combining Finnomena + SEC data

**Deliverable:** THAIFA uses sec-go as fallback. Out of scope for this repo.

---

## Key Design Decisions

| Decision | Rationale |
|----------|-----------|
| **V2 only** | V1 is dead. No backward compatibility needed. |
| **Context-first** | All methods accept `context.Context` for cancellation/timeouts. |
| **Functional options** | Flexible client configuration without breaking changes. |
| **Single package (`sec`)** | Simpler import path; all endpoints exposed on `*Client`. |
| **No enum types for strings** | SEC uses string codes (e.g., `fund_status: "RG"`). Keep as strings to avoid breakage if SEC adds new codes. Document known values in `docs/v2-schemas/API.md`. |
| **Rate limiting built-in** | Library handles SEC's limit transparently. |
| **No OHLCV emulation** | SEC provides daily closing NAV only. Don't fake data. |
| **Generic pagination** | `FetchAllPages[T]` works for any paginated endpoint. |
| **Batch operations separate** | `BatchGetNAVs` / `BatchGetFundProfiles` are opt-in concurrency helpers. |
| **Convenience helpers** | `SearchFunds`, `GetFundsByCompany`, `FindAMC`, `GetFundPortfolio` provide a simpler API on top of raw paginated endpoints. |
| **Flexible datetime parsing** | SEC returns inconsistent datetime formats (with/without timezone, fractional seconds). A custom `DateTime` type handles them transparently. |
| **Language support built-in** | Bilingual field helpers + Thai→English translation maps for Thai-only fields. |
| **Integration tests gated** | Real API tests behind `go:build integration` to avoid accidental usage/cost. |

---

## Next Steps

1. ✅ Phase 1 — Generic client
2. ✅ Phase 2 — Schema discovery
3. ✅ Phase 3 — Endpoint implementation (all 21 Fund API endpoints)
4. ✅ Phase 4 — Polish, docs, examples, CLI, integration tests
5. ⏳ Phase 5 — THAIFA integration (separate repo)

All sec-go code is complete and passing checks.
