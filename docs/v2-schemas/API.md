# SEC Open Data API - Complete Reference

> **82 endpoints** across 6 API categories.

This documentation is split by API category for easier navigation.

---

## Table of Contents

- [Common Standards](API.md#common-standards)
  - [Base URLs](API.md#base-urls)
  - [Authentication](API.md#authentication)
  - [Rate Limits](API.md#rate-limits)
  - [Pagination](API.md#pagination)
  - [Response Codes](API.md#response-codes)
  - [Parameter Types](API.md#parameter-types)
  - [Legend](API.md#legend)
- [Fund API (v2) - 21 endpoints](API-Fund.md)
- [Bond API (v2) - 6 endpoints](API-Bond.md)
- [PVD API (v1) - 15 endpoints](API-PVD.md)
- [License Check API (v1) - 16 endpoints](API-LicenseCheck.md)
- [One Report API (v1) - 23 endpoints](API-OneReport.md)
- [Digital Asset API (v1) - 1 endpoint](API-DigitalAsset.md)
- [Lookup Tables](API.md#lookup-tables)
- [Implementation Notes](API.md#implementation-notes)

---

## Common Standards

### Base URLs

| API | Base URL | Version |
|-----|----------|---------|
| Fund | `https://api.sec.or.th` | v2 |
| Bond | `https://api.sec.or.th` | v2 |
| PVD | `https://api.sec.or.th` | v1 |
| License Check | `https://api.sec.or.th` | v1 |
| One Report | `https://api.sec.or.th/onereport` | v1 |
| Digital Asset | `https://api.sec.or.th/DigitalAsset` | v1 |

### Authentication

All endpoints require the `Ocp-Apim-Subscription-Key` request header with your API subscription key.

```http
GET /v2/fund/general-info/amcs HTTP/1.1
Host: api.sec.or.th
Ocp-Apim-Subscription-Key: {your_api_key}
Accept: application/json
```

### Rate Limits

- **5,000 requests per 300 seconds**
- **Minimum 60ms delay between consecutive requests**
- HTTP **421 Misdirected Request** returned when rate limit is exceeded
- The `Retry-After` response header contains the number of seconds to wait

### Pagination

#### Paginated Endpoints (Fund, Bond, PVD, License Check)

These APIs return a paginated envelope:

```json
{
  "message": "success",
  "page_size": 100,
  "next_cursor": "xxxx-xxx-xxx",
  "items": [...]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `message` | string | Always `"success"` on 200 OK |
| `page_size` | int | Number of items in this page |
| `next_cursor` | string \| null | Cursor for the next page. Empty or null when there are no more pages. |
| `items` | array | Result items for this page |

**Pagination Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | No | 100 | Items per page (1-100) |
| `next_cursor` | string | No | `""` | Cursor from the previous response |

**Do not parse or construct cursor values manually.** Pass the exact `next_cursor` value from the previous response.

#### Direct Array Endpoints (One Report, Digital Asset)

These APIs return arrays directly (no pagination envelope):

```json
[
  { ... },
  { ... }
]
```

### Response Codes

| Status | Meaning | Notes |
|--------|---------|-------|
| 200 OK | Success | Response body contains data |
| 204 No Content | No results | Treat as "not found" for the given filters |
| 400 Bad Request | Invalid parameters | Check date formats, period formats |
| 401 Unauthorized | Missing or invalid API key | Verify `Ocp-Apim-Subscription-Key` header |
| 403 Forbidden | Subscription not allowed | Key may lack access to this API |
| 421 Misdirected Request | Rate limited | Read `Retry-After` header and wait |
| 500 Internal Server Error | SEC server error | Retry with exponential backoff |

**Error Response Body:** On non-2xx responses, the body is typically empty or contains a plain text message.

### Parameter Types

| Type | Format | Example |
|------|--------|---------|
| `date` | ISO 8601 (`YYYY-MM-DD`) | `2024-01-15` |
| `datetime` | ISO 8601 with time. SEC is inconsistent: may include timezone (`Z` or `+07:00`), fractional seconds (`2024-01-15T09:30:00.9`), or neither. | `2024-01-15T09:30:00`, `2024-01-15T09:30:00.9`, `2024-01-15T09:30:00+07:00` |
| `period` | `YYYYMM` | `202401` |
| `proj_id` | `{Type}{ID}_YYYY` | `M0000_2552`, `PRINCIPALi9` |

**Note on datetime parsing:** Use a lenient datetime parser. SEC does not consistently apply a single datetime format across endpoints or records.

### Legend

- **R** = Required
- **O** = Optional
- Fields marked **nullable** may be returned as `null`
- Fields marked **?** are observed in responses but may not always be present

---

## Lookup Tables

### Fund Status Values

| Status | Description |
|--------|-------------|
| `Registered` | Officially registered and active |
| `IPO` | Initial Public Offering period |
| `Expired` | Offering period ended |
| `Canceled` | Project canceled |
| `Liquidated` | Fund liquidated |

### Management Style Values

| Code | Description |
|------|-------------|
| `AM` | Active Management |
| `AN` | Thai fund follows master fund (Active) |
| `PM` | Passive Management/Index Tracking |
| `PN` | Thai fund follows master fund (Passive) |
| `IM` | Inverse Management |
| `IN` | Thai fund follows master fund (Inverse) |
| `LM` | Leveraged Management |
| `LN` | Thai fund follows master fund (Leveraged) |
| `BH` | Buy-and-Hold |
| `SM` | Index Tracking with occasional alpha |
| `OT` | Others |

### Retail Type Values

| Code | Description |
|------|-------------|
| `A` | Non-retail investors |
| `B` | High net worth investors |
| `F` | Liquidity support for corporate bonds |
| `G` | Government policy response |
| `H` | Non-retail and high net worth |
| `N` | Institutional investors |
| `R` | General investors |
| `V` | Provident funds |
| `X` | Institutional and ultra-high net worth |

### Risk Spectrum Values

| Code | Typical Description |
|------|---------------------|
| `RS1` | Lowest risk |
| `RS2` | Low risk |
| `RS3` | Moderately low risk |
| `RS4` | Moderate risk |
| `RS5` | Moderately high risk |
| `RS6` | High risk |
| `RS7` | Very high risk |
| `RS8` | Highest risk |
| `RS81` | Complex/sophisticated investment |

### Prospectus Type Values

| Code | Description |
|------|-------------|
| `IPO` | Initial Public Offering factsheet |
| `Monthly` | Monthly factsheet update |
| `SignificantFactsheet` | Significant change factsheet |

---

## Implementation Notes

1. **Always send `Ocp-Apim-Subscription-Key`** as a request header. Query parameter auth is not supported.

2. **Respect rate limits.** If you receive HTTP 421, read the `Retry-After` header and wait before retrying. The Go client handles this automatically.

3. **Date parameters must be exact.** Use `YYYY-MM-DD` for dates and `YYYYMM` for periods. Malformed dates usually return HTTP 400.

4. **Pagination cursor encoding is opaque.** Do not parse or construct cursor values manually. Pass the exact `next_cursor` value from the previous response.

5. **`latest=true` is the most common query pattern** for factsheet endpoints when you only need current fund information.

6. **Empty `items` arrays are valid.** A 200 OK with `items: []` means the query succeeded but no records matched.

7. **Many fields are returned as strings containing numbers.** The Go client uses `string` types for fields like `portfolio_turnover_ratio`, `sharpe_ratio`, etc., because SEC returns them as strings rather than JSON numbers.

8. **One Report and Digital Asset use different base URLs.** One Report uses `https://api.sec.or.th/onereport` and Digital Asset uses `https://api.sec.or.th/DigitalAsset`. Do not prepend `/v1` or `/v2` to these paths.

9. **One Report schemas are still being finalized.** Many endpoints return only `last_upd_date`, `report_year`, and `unique_id` as confirmed fields. Additional fields will be documented as SEC publishes them.

---

## Version Migration Notice

- **Fund & Bond APIs**: v1 → v2 (deadline: April 30, 2026)
- **Other APIs** (PVD, License Check, One Report, Digital Asset): legacy → v1 (deadline: June 30, 2026)

All endpoints documented in this file use the latest stable version as of the documentation date.
