# Fund API (v2)

Base URL: `https://api.sec.or.th`

All Fund endpoints use the standard [paginated envelope](../API.md#pagination) with `message`, `page_size`, `next_cursor`, and `items`.

---

## Table of Contents

- [General Information](#general-information)
- [Factsheet](#factsheet)
- [Outstanding](#outstanding)
- [Daily Information](#daily-information)

---

## General Information

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

---

### 4. Mutual Fund Fees

```http
GET /v2/fund/general-info/mutual-fund-fees?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&fund_class_name={fund_class_name}
```

**Query Parameters:** Same as endpoint 3.

**Response Items (`MutualFundFee`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `fee_type_desc` | string | No | Fee type (`Management Fee`, `Trustee Fee`, etc.) |
| `rate` | float | Yes | Fee rate |
| `rate_unit` | string | Yes | Rate unit |
| `fee_other_desc` | string | Yes | Additional remarks |

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

---

## Factsheet

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

---

## Outstanding

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

---

### 19. Monthly Fund Portfolio by Asset Type

```http
GET /v2/fund/outstanding/portfolio-asset-type?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 18.

**Response Items (`MonthlyPortfolioAssetType`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `assetliab_code` | string | No | Investment category code |
| `assetliab_desc` | string | No | Category description |
| `market_value` | number | No | Market value (THB) |
| `percent_nav` | number | Yes | %NAV |

---

## Daily Information

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
