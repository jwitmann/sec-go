# SEC Open API V2 - Fund Endpoints

## Base URL

```
https://api.sec.or.th
```

## Authentication

All endpoints require the `Ocp-Apim-Subscription-Key` request header with your API subscription key.

```http
GET /v2/fund/general-info/amcs HTTP/1.1
Host: api.sec.or.th
Ocp-Apim-Subscription-Key: {your_api_key}
Accept: application/json
```

## Rate Limits

- **5,000 requests per 300 seconds**
- **Minimum 16ms delay between consecutive requests**
- HTTP **421 Misdirected Request** returned when rate limit is exceeded
- The `Retry-After` response header contains the number of seconds to wait

## Pagination

All list endpoints use cursor-based pagination. The response envelope is always:

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

### Pagination Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | No | 100 | Items per page (1-100) |
| `next_cursor` | string | No | `""` | Cursor from the previous response |

---

## Common Response Codes

| Status | Meaning | Notes |
|--------|---------|-------|
| 200 OK | Success | Response body contains paginated envelope |
| 204 No Content | No results | Treat as "not found" for the given filters |
| 400 Bad Request | Invalid parameters | Check date formats, period formats |
| 401 Unauthorized | Missing or invalid API key | Verify `Ocp-Apim-Subscription-Key` header |
| 403 Forbidden | Subscription not allowed | Key may lack access to this API |
| 421 Misdirected Request | Rate limited | Read `Retry-After` header and wait |
| 500 Internal Server Error | SEC server error | Retry with exponential backoff |

### Error Response Body

On non-2xx responses, the body is typically empty or contains a plain text message. The Go client surfaces these as `APIError` with the HTTP status code.

---

## Common Parameter Types

| Type | Format | Example |
|------|--------|---------|
| `date` | ISO 8601 (`YYYY-MM-DD`) | `2024-01-15` |
| `datetime` | ISO 8601 with time. SEC is inconsistent: may include timezone (`Z` or `+07:00`), fractional seconds (`2024-01-15T09:30:00.9`), or neither. | `2024-01-15T09:30:00`, `2024-01-15T09:30:00.9`, `2024-01-15T09:30:00+07:00` |
| `period` | `YYYYMM` | `202401` |
| `proj_id` | `{Type}{ID}_YYYY` | `M0000_2552`, `PRINCIPALi9` |

**Note on datetime parsing:** If you are implementing a client, use a lenient datetime parser. SEC does not consistently apply a single datetime format across endpoints or records.

---

## Legend

- **R** = Required
- **O** = Optional
- Fields marked **nullable** may be returned as `null`
- Fields marked **?** are observed in responses but may not always be present

---

## General Information Endpoints

### 1. List Asset Management Companies (AMCs)

```http
GET /v2/fund/general-info/amcs?page_size={page_size}&next_cursor={next_cursor}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |

**Response Items (`AMC`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `unique_id` | string | No | AMC identifier (e.g., `C0000000021`) |
| `comp_name_th` | string | No | Company name in Thai |
| `comp_name_en` | string | No | Company name in English |
| `last_upd_date` | datetime | No | Last update timestamp |

**Example:**
```http
GET /v2/fund/general-info/amcs?page_size=10
```

**Sample response:** [`amcs.json`](amcs.json)

---

### 2. Fund Profiles

```http
GET /v2/fund/general-info/profiles?page_size={page_size}&next_cursor={next_cursor}&fund_class_name={fund_class_name}&fund_status={fund_status}&project_info={project_info}&company_info={company_info}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `fund_class_name` | string | O | — | Filter by fund class |
| `fund_status` | string | O | — | Filter by status: `IPO`, `Registered`, `Expired`, `Canceled`, `Liquidated` |
| `project_info` | string | O | — | Search across `proj_id`, `proj_name_th`, `proj_name_en`, `proj_abbr_name` |
| `company_info` | string | O | — | Search across `unique_id`, `comp_name_th`, `comp_name_en` |

**Response Items (`FundProfile`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `unique_id` | string | No | AMC identifier |
| `proj_id` | string | No | Project number (`{Type}{ID}_YYYY`) |
| `regis_id` | string | Yes | Fund registration number |
| `proj_name_th` | string | No | Fund name in Thai |
| `proj_name_en` | string | Yes | Fund name in English |
| `proj_abbr_name` | string | No | Fund abbreviation |
| `fund_status` | string | No | `Registered`, `IPO`, `Expired`, `Canceled`, `Liquidated` |
| `init_date` | date | Yes | Establishment date |
| `regis_date` | date | Yes | Registration date |
| `cancel_date` | date | Yes | Cancellation date |
| `policy_desc` | string | Yes | Fund policy type |
| `management_style` | string | Yes | `AM`, `AN`, `PM`, `PN`, `IM`, `IN`, `LM`, `LN`, `BH`, `SM`, `OT` |
| `fund_class_name` | string | Yes | Fund class abbreviation |
| `fund_class_isin_code` | string | Yes | ISIN code |

**Example:**
```http
GET /v2/fund/general-info/profiles?fund_status=Registered&page_size=10
```

**Sample response:** [`profiles.json`](profiles.json)

---

### 3. Fund Specifications

```http
GET /v2/fund/general-info/specifications?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&fund_class_name={fund_class_name}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |
| `fund_class_name` | string | O | — | Fund class name |

**Response Items (`FundSpecification`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `spec_code` | string | No | Special characteristic code |
| `spec_desc` | string | No | Fund type description |

**Example:**
```http
GET /v2/fund/general-info/specifications?proj_id=M0000_2552&page_size=10
```

---

### 4. Mutual Fund Fees

```http
GET /v2/fund/general-info/mutual-fund-fees?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&fund_class_name={fund_class_name}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |
| `fund_class_name` | string | O | — | Fund class name |

**Response Items (`MutualFundFee`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `fee_type_desc` | string | No | Fee type (`Management Fee`, `Trustee Fee`, etc.) |
| `rate` | float | Yes | Fee rate |
| `rate_unit` | string | Yes | Rate unit |
| `fee_other_desc` | string | Yes | Additional remarks |

**Example:**
```http
GET /v2/fund/general-info/mutual-fund-fees?proj_id=M0000_2552&page_size=10
```

---

### 5. Fund Involve Parties

```http
GET /v2/fund/general-info/involve-parties?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&entity_type={entity_type}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |
| `entity_type` | string | O | — | Entity type code |

**Response Items (`FundInvolveParty`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `entity_type` | string | No | Entity type code |
| `entity_name_th` | string | Yes | Name in Thai |
| `entity_name_en` | string | Yes | Name in English |
| `address` | string | Yes | Address |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Entity Type Values:**

| Code | Description |
|------|-------------|
| `A` | Auditor |
| `U` | Underwriter |
| `S` | Selling Agent |
| `R` | Registrar |
| `V` | Mutual Fund Supervisor / Trustee |
| `M` | Investment Solicitor / Advisor |
| `O` | Outsource Company (Investment Management) |
| `P` | Private Equity / Professional Investor |
| `K` | Market Maker |
| `N` | Financial Advisor |
| `F` | Fund Manager |

**Example:**
```http
GET /v2/fund/general-info/involve-parties?proj_id=M0000_2552&entity_type=F&page_size=10
```

---

## Factsheet Endpoints

Factsheet endpoints share common filtering parameters:

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Filter by project ID |
| `start_date` | date | O | — | Factsheet effective start date (`YYYY-MM-DD`) |
| `end_date` | date | O | — | Factsheet effective end date (`YYYY-MM-DD`) |
| `latest` | boolean | O | `false` | Return only the latest factsheet record |
| `fund_class_name` | string | O | — | Filter by fund class |

All factsheet items include `start_date`, `end_date`, and `prospectus_type`. `end_date` is `null` when the factsheet is currently effective.

### 6. Fund Factsheet URLs

```http
GET /v2/fund/factsheet/urls?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&fund_class_name={fund_class_name}
```

**Response Items (`FundFactsheetURL`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `prospectus_type` | string | No | Factsheet frequency |
| `amc_url_factsheet` | string | Yes | URL on AMC website |
| `pdf_factsheet` | string | Yes | PDF URL hosted by SEC |
| `as_of_date` | date | Yes | Reference date |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/urls?proj_id=M0000_2552&page_size=10
```

---

### 7. IPO Offering Period

```http
GET /v2/fund/factsheet/ipos?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}
```

**Response Items (`FundIPO`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `first_sell_start_date` | string | Yes | IPO start date |
| `first_sell_end_date` | string | Yes | IPO end date |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/ipos?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 8. Fund Benchmark

```http
GET /v2/fund/factsheet/benchmarks?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}
```

**Response Items (`FactsheetBenchmark`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `group_seq` | int | No | Group sequence number |
| `benchmark` | string | No | Benchmark index |
| `benchmark_remark` | string | Yes | Remark |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/benchmarks?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 9. Subscription/Redemption Minimums

```http
GET /v2/fund/factsheet/subscription-redemption-minimums?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FactsheetSubscriptionRedemptionMinimum`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `minimum_sub_ipo` | float | Yes | Minimum initial subscription amount |
| `minimum_sub_ipo_cur` | string | Yes | Currency for minimum initial subscription |
| `minimum_sub` | float | Yes | Minimum subsequent subscription amount |
| `minimum_sub_cur` | string | Yes | Currency for minimum subsequent subscription |
| `minimum_sub_unit` | string | Yes | Minimum subscription units |
| `minimum_redempt` | float | Yes | Minimum redemption amount |
| `minimum_redempt_cur` | string | Yes | Currency for minimum redemption |
| `minimum_redempt_unit` | string | Yes | Minimum redemption units |
| `lowbal_val` | float | Yes | Minimum remaining balance amount |
| `lowbal_val_cur` | string | Yes | Currency for minimum balance |
| `lowbal_unit` | string | Yes | Minimum remaining unit balance |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/subscription-redemption-minimums?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 10. Subscription/Redemption Periods

```http
GET /v2/fund/factsheet/subscription-redemption-periods?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FactsheetSubscriptionRedemptionPeriod`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `type` | string | No | Transaction type: `subscription` or `redemption` |
| `period` | string | No | Subscription/redemption period |
| `redemp_period_oth` | string | Yes | Additional description if period is `Other` |
| `settlement_period` | string | Yes | Settlement period for redemption payment |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/subscription-redemption-periods?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 11. Fund Risk Spectrum

```http
GET /v2/fund/factsheet/risk-spectrum?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}
```

**Response Items (`RiskSpectrum`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `risk_spectrum` | string | No | Risk level: `RS1`-`RS8`, `RS81` |
| `risk_spectrum_desc` | string | No | Risk level description |

**Example:**
```http
GET /v2/fund/factsheet/risk-spectrum?proj_id=M0000_2552&latest=true&page_size=10
```

**Sample response:** [`risk-spectrum.json`](risk-spectrum.json)

---

### 12. Fund Statistics

```http
GET /v2/fund/factsheet/statistics?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FactsheetStatistics`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `portfolio_turnover_ratio` | string | Yes | Portfolio Turnover Ratio |
| `recovering_period` | string | Yes | Recovering Period |
| `portfolio_duration_period` | string | Yes | Portfolio Duration for bond funds |
| `maximum_drawdown` | string | Yes | Maximum drawdown over past 5 years |
| `sharpe_ratio` | string | Yes | Sharpe Ratio (equity funds) |
| `beta` | string | Yes | Beta (equity funds) |
| `alpha` | string | Yes | Alpha (equity funds) |
| `fx_hedging` | string | Yes | FX Hedging |
| `tracking_error` | string | Yes | Tracking Error |
| `yield_to_maturity` | string | Yes | Yield to Maturity |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/statistics?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 13. Fund Dividend Policy

```http
GET /v2/fund/factsheet/dividend-policy?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FundDividendPolicy`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `dividend_policy` | string | No | Dividend payment policy |
| `last_upd_date` | datetime | Yes | Last update timestamp |

**Example:**
```http
GET /v2/fund/factsheet/dividend-policy?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 14. Fund Fees (Factsheet)

```http
GET /v2/fund/factsheet/fees?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FactsheetFee`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `fee_type_desc` | string | No | `Front-end Fee`, `Back-end Fee`, etc. |
| `rate` | float | Yes | Specified rate |
| `actual_value` | float | Yes | Actual charged rate |

**Example:**
```http
GET /v2/fund/factsheet/fees?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 15. Historical Fund Performance

```http
GET /v2/fund/factsheet/performance?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`FactsheetPerformance`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `performance_type_desc` | string | No | Performance type description |
| `reference_period` | string | No | Reference period and look-back year |
| `performance_value` | float | Yes | Historical performance value |

**Example:**
```http
GET /v2/fund/factsheet/performance?proj_id=M0000_2552&latest=true&page_size=10
```

---

### 16. Asset Allocation

```http
GET /v2/fund/factsheet/asset-allocation?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}&fund_class_name={fund_class_name}
```

**Response Items (`AssetAllocation`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `asset_seq` | int | No | Asset sequence/order |
| `asset_name` | string | No | Asset type name (Thai) |
| `asset_ratio` | float | Yes | Allocation percentage |

**Example:**
```http
GET /v2/fund/factsheet/asset-allocation?proj_id=M0000_2552&latest=true&page_size=10
```

**Sample response:** [`asset-allocation.json`](asset-allocation.json)

---

### 17. Top 5 Holdings

```http
GET /v2/fund/factsheet/top5-holdings?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_date={start_date}&end_date={end_date}&latest={latest}
```

**Response Items (`Top5Holding`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `start_date` | date | No | Effective start date |
| `end_date` | date | Yes | Effective end date (`null` if current) |
| `prospectus_type` | string | No | `IPO`, `Monthly`, `SignificantFactsheet` |
| `asset_seq` | int | No | Holding rank (1-5) |
| `asset_name` | string | No | Asset/security name |
| `asset_ratio` | float | Yes | %NAV allocation |

**Example:**
```http
GET /v2/fund/factsheet/top5-holdings?proj_id=M0000_2552&latest=true&page_size=10
```

**Sample response:** [`top5-holdings.json`](top5-holdings.json)

---

## Outstanding Endpoints

### 18. Quarterly Fund Portfolio

```http
GET /v2/fund/outstanding/portfolio?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Filter by project ID |
| `start_period` | period | O | — | Starting period (`YYYYMM`) |
| `end_period` | period | O | — | Ending period (`YYYYMM`) |

**Response Items (`QuarterlyPortfolio`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `as_of_date` | date | No | Portfolio as-of date |
| `assetliab_id` | string | No | Asset/liability category code |
| `assetliab_desc` | string | No | Category description |
| `issue_code` | string | Yes | Security symbol |
| `isin_code` | string | Yes | ISIN code |
| `issuer` | string | Yes | Issuer name |
| `assetliab_value` | number | No | Market value (THB) |
| `percent_nav` | number | Yes | %NAV |

**Example:**
```http
GET /v2/fund/outstanding/portfolio?proj_id=M0000_2552&start_period=202401&end_period=202403&page_size=10
```

**Sample response:** [`quarterly-portfolio.json`](quarterly-portfolio.json)

---

### 19. Monthly Fund Portfolio by Asset Type

```http
GET /v2/fund/outstanding/portfolio-asset-type?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Filter by project ID |
| `start_period` | period | O | — | Starting period (`YYYYMM`) |
| `end_period` | period | O | — | Ending period (`YYYYMM`) |

**Response Items (`MonthlyPortfolioAssetType`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `assetliab_code` | string | No | Investment category code |
| `assetliab_desc` | string | No | Category description |
| `market_value` | number | No | Market value (THB) |
| `percent_nav` | number | Yes | %NAV |

**Example:**
```http
GET /v2/fund/outstanding/portfolio-asset-type?proj_id=M0000_2552&start_period=202401&end_period=202401&page_size=10
```

**Sample response:** [`monthly-portfolio-asset.json`](monthly-portfolio-asset.json)

---

## Daily Information Endpoints

### 20. Daily Fund NAV

```http
GET /v2/fund/daily-info/nav?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_nav_date={start_nav_date}&end_nav_date={end_nav_date}&fund_class_name={fund_class_name}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Filter by project ID |
| `start_nav_date` | date | O | — | Start date (`YYYY-MM-DD`) |
| `end_nav_date` | date | O | — | End date (`YYYY-MM-DD`) |
| `fund_class_name` | string | O | — | Filter by fund class |

**Response Items (`DailyNAV`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `nav_date` | date | No | NAV date |
| `fund_class_name` | string | No | Fund class |
| `net_asset` | number | Yes | Net asset value (THB) |
| `last_val` | number | Yes | NAV per unit (THB/unit) |
| `sell_price` | number | Yes | Selling price |
| `buy_price` | number | Yes | Redemption price |
| `sell_swap_price` | number | Yes | Switch-in price |
| `buy_swap_price` | number | Yes | Switch-out price |

**Example:**
```http
GET /v2/fund/daily-info/nav?proj_id=M0000_2552&start_nav_date=2024-01-01&end_nav_date=2024-01-31&page_size=10
```

**Sample responses:** [`daily-nav-recent.json`](daily-nav-recent.json), [`daily-nav-range.json`](daily-nav-range.json)

---

### 21. Dividend History

```http
GET /v2/fund/daily-info/dividend-history?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_dividend_date={start_dividend_date}&end_dividend_date={end_dividend_date}&class_abbr_name={class_abbr_name}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Filter by project ID |
| `start_dividend_date` | date | O | — | Start dividend date (`YYYY-MM-DD`) |
| `end_dividend_date` | date | O | — | End dividend date (`YYYY-MM-DD`) |
| `class_abbr_name` | string | O | — | Filter by class abbreviation |

**Response Items (`DividendHistory`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `unique_id` | string | No | AMC identifier |
| `class_abbr_name` | string | No | Class abbreviation |
| `book_close_date` | datetime | Yes | Book close date |
| `dividend_date` | datetime | Yes | Dividend payment date |
| `dividend_value` | float | Yes | Dividend amount per unit |

**Example:**
```http
GET /v2/fund/daily-info/dividend-history?proj_id=M0000_2552&start_dividend_date=2024-01-01&end_dividend_date=2024-12-31&page_size=10
```

**Sample response:** [`dividend-history.json`](dividend-history.json)

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
