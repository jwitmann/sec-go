# SEC-Go Development Plan

## Objective

Build a Go client library for the **Thailand SEC OpenAPI V2** (`api.sec.or.th`) — a complementary/fallback data source to Finnomena for Thai mutual fund data.

**Constraint:** V1 API is deprecated. All development targets V2 exclusively.

---

## Project Structure

```
sec-go/
├── client.go              # Core HTTP client, auth, retry, rate limiting
├── models.go              # V2 data models (discovered in Phase 2)
├── options.go             # Functional options for client configuration
├── error.go               # Custom error types
├── rate.go                # Rate limiter (token bucket or simple delay)
├── retry.go               # Exponential backoff logic
│
├── fund/                  # Fund API endpoints
│   ├── general.go         # /v2/fund/general-info/* endpoints
│   ├── factsheet.go       # Fund factsheet endpoints
│   ├── daily.go           # Daily NAV endpoints
│   └── performance.go     # Performance data endpoints
│
├── internal/
│   └── testutil/          # Mock server helpers for tests
│
├── docs/
│   ├── PLAN.md            # This document
│   └── v2-schemas/        # Sample V2 responses (populated in Phase 2)
│
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

---

## Phase 1: Generic Client Infrastructure

**Goal:** Build everything that doesn't depend on V2 schemas. This is safe to implement now.

### 1.1 Client (`client.go`)

```go
type Client struct {
    httpClient *http.Client
    baseURL    string
    apiKey     string
    rateLimiter RateLimiter
    maxRetries int
    retryDelay time.Duration
}
```

- **Base URL:** `https://api.sec.or.th`
- **Auth header:** `Ocp-Apim-Subscription-Key: {key}`
- **Timeout:** Configurable (default 30s)
- **User-Agent:** `sec-go/{version}`

### 1.2 Functional Options (`options.go`)

```go
type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option
func WithBaseURL(url string) Option
func WithRateLimiter(rl RateLimiter) Option
func WithMaxRetries(n int) Option
func WithRetryDelay(d time.Duration) Option
```

### 1.3 Rate Limiting (`rate.go`)

SEC rate limit: **3,000 calls per 300 seconds** (10 req/sec).

Implement a simple rate limiter:
- Minimum 100ms delay between requests
- Optional: Token bucket for burst handling
- Thread-safe

### 1.4 Retry Logic (`retry.go`)

Exponential backoff with jitter:
- Retry on: 429, 500, 502, 503, 504, network errors
- Do NOT retry on: 400, 401, 403, 404, 204
- Base delay: 100ms
- Max delay: 5s
- Max retries: 3 (configurable)

### 1.5 Error Handling (`error.go`)

```go
type APIError struct {
    StatusCode int
    Message    string
    RawBody    []byte
}

var (
    ErrRateLimited    = errors.New("rate limited")
    ErrNotFound       = errors.New("data not found (204)")
    ErrUnauthorized   = errors.New("invalid API key")
)
```

### 1.6 Testing

- Mock HTTP server for all client tests
- Test auth header injection
- Test rate limiting
- Test retry behavior
- Test error parsing

**Deliverable:** `client.go`, `options.go`, `rate.go`, `retry.go`, `error.go` + tests — all passing.

---

## Phase 2: V2 Schema Discovery

**Goal:** Discover actual V2 endpoints, request parameters, and response schemas.

**Prerequisite:** User must have V2 API key and test endpoints.

### 2.1 Discovery Script

Create a small Go program or shell script to systematically probe V2:

```bash
# Step 1: List AMCs
curl -H "Ocp-Apim-Subscription-Key: $V2_KEY" \
  https://api.sec.or.th/v2/fund/general-info/amcs

# Step 2: Get fund profiles
curl -H "Ocp-Apim-Subscription-Key: $V2_KEY" \
  "https://api.sec.or.th/v2/fund/general-info/profiles?company_info=UNIQUE_ID"

# Step 3: Probe for NAV endpoint
curl -H "Ocp-Apim-Subscription-Key: $V2_KEY" \
  "https://api.sec.or.th/v2/fund/daily-nav/{proj_id}?nav_date=2024-01-15"

# Step 4: Probe for dividend endpoint
curl -H "Ocp-Apim-Subscription-Key: $V2_KEY" \
  "https://api.sec.or.th/v2/fund/daily-nav/{proj_id}/dividend"
```

### 2.2 Endpoint Inventory

Document all working endpoints:

| Endpoint | Method | Auth | Parameters | Status |
|----------|--------|------|------------|--------|
| `/v2/fund/general-info/amcs` | GET | V2 Key | None | ? |
| `/v2/fund/general-info/profiles` | GET | V2 Key | `company_info` | ? |
| `/v2/fund/...` | ? | ? | ? | ? |

### 2.3 Response Samples

Save raw JSON responses to `docs/v2-schemas/`:
- `amcs-response.json`
- `profiles-response.json`
- `nav-response.json`
- `dividend-response.json`
- etc.

These become the source of truth for Go struct generation.

### 2.4 OpenAPI Spec Hunt

Check if V2 portal provides:
- Swagger UI with spec download
- Postman collection export
- Any API documentation download

**Deliverable:** Complete endpoint inventory + sample responses in `docs/v2-schemas/`

---

## Phase 3: V2 Endpoint Implementation

**Goal:** Build typed Go methods for all discovered V2 endpoints.

### 3.1 Model Generation (`models.go`)

From sample responses, create Go structs:

```go
type AMC struct {
    UniqueID     string `json:"unique_id"`
    NameTH       string `json:"name_th"`
    NameEN       string `json:"name_en"`
    LastUpdDate  string `json:"last_upd_date"`
}

type FundProfile struct {
    ProjID         string `json:"proj_id"`
    ProjNameTH     string `json:"proj_name_th"`
    ProjNameEN     string `json:"proj_name_en"`
    ProjAbbrName   string `json:"proj_abbr_name"`
    FundClassName  string `json:"fund_class_name"`
    AMCUniqueID    string `json:"amc_unique_id"`
    FundStatus     string `json:"fund_status"`
    // ... other fields from actual response
}

type DailyNAV struct {
    NavDate       string  `json:"nav_date"`
    LastVal       float64 `json:"last_val"`
    PreviousVal   float64 `json:"previous_val"`
    NetAsset      float64 `json:"net_asset"`
    ClassAbbrName string  `json:"class_abbr_name"`
    ProjID        string  `json:"proj_id"`
    // ... amc_info, etc.
}
```

### 3.2 Service Methods

```go
// fund/general.go
func (c *Client) ListAMCs(ctx context.Context) ([]AMC, error)
func (c *Client) GetFundProfiles(ctx context.Context, companyInfo string) ([]FundProfile, error)

// fund/daily.go
func (c *Client) GetDailyNAV(ctx context.Context, projID, navDate string) (*DailyNAV, error)
func (c *Client) GetNAVRange(ctx context.Context, projID string, from, to time.Time) ([]DailyNAV, error)

// fund/factsheet.go
func (c *Client) GetFundFactsheet(ctx context.Context, projID string) (*FundFactsheet, error)

// fund/performance.go
func (c *Client) GetFundPerformance(ctx context.Context, projID string) (*Performance, error)
```

### 3.3 Batch Operations

Implement concurrent fetching with rate limiting:
- `GetNAVRange()` fetches multiple dates concurrently
- Respects rate limits
- Optional progress callback
- Skips weekends/holidays (user-configurable)

### 3.4 Testing

- Mock server tests for each endpoint
- Integration tests (gated by `-integration` flag)
- Batch operation tests

**Deliverable:** All endpoint methods + tests passing.

---

## Phase 4: Polish & Documentation

### 4.1 README

- Installation: `go get github.com/jerome/something`
- Quick start example
- Authentication setup
- Rate limiting explanation
- Available endpoints
- Error handling

### 4.2 Examples

Create `examples/` directory:
- `basic/main.go` — list AMCs, get NAV
- `batch/main.go` — fetch NAV range for multiple funds
- `thaifa/main.go` — integration pattern

### 4.3 Go Best Practices

- Context support on all methods
- Proper error wrapping
- No `interface{}` — all typed
- JSON tags match API exactly
- Zero-value structs are valid

---

## Phase 5: THAIFA Integration (Future)

**Not part of sec-go** — this is in the THAIFA repo.

### 5.1 PriceService Changes

```go
type PriceService struct {
    finnomenaClient *finnomena.Client
    secClient       *sec.Client      // NEW
    // ...
}
```

### 5.2 Fallback Logic

```go
func (s *PriceService) GetPrices(fundID, shortCode string, from, to time.Time) (*models.BarsResponse, error) {
    // Try Finnomena first
    bars, err := s.finnomenaClient.GetHistoricalPrices(shortCode, from, to)
    if err == nil && len(bars.Time) > 0 {
        return bars, nil
    }

    // Fallback to SEC
    if s.secClient != nil {
        return s.fetchFromSEC(fundID, from, to)
    }

    return nil, err
}
```

### 5.3 Mapping

Need mapping between Finnomena `fund_id`/`short_code` and SEC `proj_id`.

**Deliverable:** THAIFA uses sec-go as fallback. Out of scope for this repo.

---

## Key Design Decisions

| Decision | Rationale |
|----------|-----------|
| **V2 only** | V1 is dead. No backward compatibility needed. |
| **Context-first** | All methods accept `context.Context` for cancellation/timeouts. |
| **Functional options** | Flexible client configuration without breaking changes. |
| **Separate `fund/` package** | Organize endpoints by domain; keep root package clean. |
| **No enum types for strings** | SEC uses string codes (e.g., `fund_status: "RG"`). Keep as strings to avoid breakage if SEC adds new codes. Document known values in comments. |
| **Rate limiting built-in** | Library handles SEC's 10 req/sec limit transparently. |
| **No OHLCV emulation** | SEC provides daily closing NAV only. Don't fake data. |

---

## Open Questions (Resolve in Phase 2)

1. **V2 base path confirmation** — Is it `/v2/fund/...` or something else?
2. **Auth mechanism** — Still `Ocp-Apim-Subscription-Key` header?
3. **Endpoint paths** — What are the exact V2 endpoint paths?
4. **Response schemas** — What fields does each endpoint return?
5. **Pagination** — Does V2 paginate large lists?
6. **Date format** — Still `YYYY-MM-DD` for NAV dates?
7. **Share class handling** — How is `fund_class_name` used in V2 vs V1's `class_abbr_name`?

---

## Next Steps

1. **Initialize repo** — Create `sec-go` repo, `go mod init`
2. **Phase 1** — Build generic client (no V2 knowledge needed)
3. **Phase 2** — User tests V2 endpoints, shares responses
4. **Phase 3** — Implement V2 endpoints from sample responses
5. **Phase 4** — Polish, document, examples
6. **Phase 5** — THAIFA integration (separate repo)

---

## Files to Create (Phase 1)

```
sec-go/
├── go.mod
├── client.go
├── options.go
├── error.go
├── rate.go
├── retry.go
├── client_test.go
└── README.md (stub)
```

All Phase 1 code is V2-agnostic and safe to write now.
