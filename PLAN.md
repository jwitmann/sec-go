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
├── fund_service.go        # V2 Fund API methods
├── models.go              # V2 data models
├── pagination.go          # Pagination helper (FetchAllPages)
├── batch.go               # Concurrent batch operations
│
├── client_test.go         # Client tests
├── fund_service_test.go   # Service method tests
├── pagination_test.go     # Pagination tests
├── batch_test.go          # Batch operation tests
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
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── PLAN.md                # This document
```

---

## Phase 1: Generic Client Infrastructure ✅ COMPLETE

**Goal:** Build everything that doesn't depend on V2 schemas.

**Deliverables:**
- `client.go` — Core HTTP client with auth, retry, rate limiting
- `options.go` — Functional options (HTTP client, base URL, secondary key, rate limiter, retries, timeout, cache)
- `error.go` — Custom error types (`APIError`, `ErrRateLimited`, `ErrNotFound`, `ErrUnauthorized`)
- `rate.go` — Thread-safe rate limiter with 16ms minimum delay
- `retry.go` — Exponential backoff with jitter, handles HTTP 421 + `Retry-After`
- `client_test.go` — Mock server tests for auth, rate limiting, retry, errors, options, POST

---

## Phase 2: V2 Schema Discovery ✅ COMPLETE

**Goal:** Discover actual V2 endpoints, request parameters, and response schemas.

**Key findings:**
- Base URL: `https://api.sec.or.th`
- Auth header: `Ocp-Apim-Subscription-Key`
- Correct NAV endpoint: `/v2/fund/daily-info/nav` (not `/v2/fund/daily-nav`)
- All working endpoints use cursor-based pagination with `message`, `page_size`, `next_cursor`, `items`
- Rate limit from SEC docs: 5,000 requests / 300 seconds, minimum 16ms between requests

**Deliverables:**
- Complete endpoint inventory in `docs/v2-schemas/API.md`
- Sample JSON responses in `docs/v2-schemas/` for all working endpoints
- Test key infrastructure: `config/sec-keys.example.json` + `internal/testutil/keys.go`

---

## Phase 3: V2 Endpoint Implementation ✅ COMPLETE

**Goal:** Build typed Go methods for all discovered V2 endpoints.

**Models added:** `AMC`, `FundProfile`, `DailyNAV`, `MutualFundFee`, `FactsheetFee`, `FactsheetPerformance`, `DividendHistory`, `AssetAllocation`, `RiskSpectrum`, `Top5Holding`, `QuarterlyPortfolio`, `MonthlyPortfolioAssetType`.

**Service methods added:**
- `ListAMCs`
- `GetFundProfiles`
- `GetDailyNAV`
- `GetMutualFundFees`
- `GetFactsheetFees`
- `GetFactsheetPerformance`
- `GetDividendHistory`
- `GetAssetAllocation`
- `GetRiskSpectrum`
- `GetTop5Holdings`
- `GetQuarterlyPortfolio`
- `GetMonthlyPortfolioAssetType`

**Infrastructure added:**
- Generic `FetchAllPages[T]` for cursor pagination
- `BatchGetNAVs` + `BatchGetFundProfiles` for concurrent fetching with progress callbacks
- Refactored `fund_service.go` using `fetchPaginated[T]`, `setPagination`, `setDateRange`, `buildPath` helpers
- Zero clone groups detected by `dupl`

**Tests added:**
- `fund_service_test.go` — mock server tests for 8+ service methods + error cases
- `pagination_test.go` — pagination helper tests
- `batch_test.go` — batch operation tests (concurrency, progress, context cancellation)

**All unit tests passing: 28 tests.**

---

## Phase 4: Polish & Documentation ✅ COMPLETE

**Deliverables:**
- `README.md` — Installation, auth, endpoints, pagination, batch ops, examples, project structure
- `examples/basic/main.go` — list AMCs, get profiles, get NAV
- `examples/batch/main.go` — fetch NAV history for multiple funds concurrently
- `examples/thaifa/main.go` — THAIFA fallback integration pattern
- `Makefile` with `test`, `lint`, `format`, `check` targets
- `AGENTS.md` with project guidelines

**Quality gates:**
- `golangci-lint` ✅
- `staticcheck` ✅
- `go vet` ✅
- `gocyclo -over 25` ✅
- `dupl -t 100` ✅ (0 clone groups)
- `go test ./...` ✅ (28 tests)

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

---

## Next Steps

1. ✅ Phase 1 — Generic client
2. ✅ Phase 2 — Schema discovery
3. ✅ Phase 3 — Endpoint implementation
4. ✅ Phase 4 — Polish, docs, examples
5. ⏳ Phase 5 — THAIFA integration (separate repo)

All sec-go code is complete and passing checks.
