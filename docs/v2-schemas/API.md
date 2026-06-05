# SEC Open API V2 - Fund Endpoints

## Base URL
```
https://api.sec.or.th
```

## Authentication
All endpoints require the `Ocp-Apim-Subscription-Key` header with your API key.

## Rate Limits
- 5,000 requests per 300 seconds
- Minimum 16ms delay between requests
- HTTP 421 (Misdirected Request) with `Retry-After` header when rate limited

## Pagination
All list endpoints use cursor-based pagination:
- `page_size`: Items per page (1-100, default 100)
- `next_cursor`: Cursor for next page (empty string = no more pages)

Response structure:
```json
{
  "message": "success",
  "page_size": 100,
  "next_cursor": "xxxx-xxx-xxx",
  "items": [...]
}
```

---

## General Information Endpoints

### 1. List Asset Management Companies (AMCs)
```
GET /v2/fund/general-info/amcs
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| unique_id | string | AMC identifier (e.g., C0000000021) |
| comp_name_th | string | Company name in Thai |
| comp_name_en | string | Company name in English |
| last_upd_date | datetime | Last update timestamp |

---

### 2. Fund Profiles
```
GET /v2/fund/general-info/profiles
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| fund_class_name | string | Filter by fund class |
| fund_status | string | Filter by status (IPO, Registered, Expired, Canceled, Liquidated) |
| project_info | string | Search by proj_id, proj_name_th, proj_name_en, proj_abbr_name |
| company_info | string | Search by unique_id, comp_name_th, comp_name_en |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| unique_id | string | AMC identifier |
| proj_id | string | Project number ({Type}{ID}_YYYY) |
| regis_id | string | Fund registration number |
| proj_name_th | string | Fund name in Thai |
| proj_name_en | string | Fund name in English |
| proj_abbr_name | string | Fund abbreviation |
| fund_status | string | Registered, IPO, Expired, Canceled, Liquidated |
| init_date | date | Establishment date |
| regis_date | date | Registration date |
| cancel_date | date | Cancellation date |
| policy_desc | string | Fund policy type |
| management_style | string | AM, PM, IM, LM, BH, SM, OT |
| fund_class_name | string | Fund class abbreviation |
| fund_class_isin_code | string | ISIN code |

---

### 3. Fund Specifications
```
GET /v2/fund/general-info/specifications
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Project number |
| fund_class_name | string | Fund class name |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| spec_code | string | Special characteristic code |
| spec_desc | string | Fund type description |

---

### 4. Mutual Fund Fees
```
GET /v2/fund/general-info/mutual-fund-fees
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Project number |
| fund_class_name | string | Fund class name |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| fee_type_desc | string | Fee type (Management Fee, Trustee Fee, etc.) |
| rate | float | Fee rate |
| rate_unit | string | Rate unit |
| fee_other_desc | string | Additional remarks |

---

### 5. Fund Involve Parties
```
GET /v2/fund/general-info/involve-parties
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Project number |
| entity_type | string | Entity type code (A, U, S, R, V, M, O, P, K, N, F) |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| entity_type | string | Entity type code |
| entity_name_th | string | Name in Thai |
| entity_name_en | string | Name in English |
| address | string | Address |
| last_upd_date | datetime | Last update timestamp |

**Entity Type Values:**
| Code | Description |
|------|-------------|
| A | Auditor |
| U | Underwriter |
| S | Selling Agent |
| R | Registrar |
| V | Mutual Fund Supervisor / Trustee |
| M | Investment Solicitor / Advisor |
| O | Outsource Company (Investment Management) |
| P | Private Equity / Professional Investor |
| K | Market Maker |
| N | Financial Advisor |
| F | Fund Manager |

---

## Daily Information Endpoints

### 20. Daily Fund NAV
```
GET /v2/fund/daily-info/nav
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_nav_date | date | Start date (YYYY-MM-DD) |
| end_nav_date | date | End date (YYYY-MM-DD) |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| nav_date | date | NAV date |
| fund_class_name | string | Fund class |
| net_asset | number | Net asset value (THB) |
| last_val | number | NAV per unit (THB/unit) |
| sell_price | number | Selling price |
| buy_price | number | Redemption price |
| sell_swap_price | number | Switch-in price |
| buy_swap_price | number | Switch-out price |

---

## Factsheet Endpoints

### 14. Fund Fees (Factsheet)
```
GET /v2/fund/factsheet/fees
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| fee_type_desc | string | Front-end Fee, Back-end Fee, etc. |
| rate | float | Specified rate |
| actual_value | float | Actual charged rate |

---

### 6. Fund Factsheet URLs
```
GET /v2/fund/factsheet/urls
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Project number |
| fund_class_name | string | Fund class name |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| prospectus_type | string | Factsheet frequency |
| amc_url_factsheet | string | URL on AMC website |
| pdf_factsheet | string | PDF URL hosted by SEC |
| as_of_date | date | Reference date |
| last_upd_date | datetime | Last update timestamp |

---

### 7. IPO Offering Period
```
GET /v2/fund/factsheet/ipos
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| first_sell_start_date | string | IPO start date |
| first_sell_end_date | string | IPO end date |
| last_upd_date | datetime | Last update timestamp |

---

### 8. Fund Benchmark
```
GET /v2/fund/factsheet/benchmarks
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| group_seq | int | Group sequence number |
| benchmark | string | Benchmark index |
| benchmark_remark | string | Remark |
| last_upd_date | datetime | Last update timestamp |

---

### 9. Subscription/Redemption Minimums
```
GET /v2/fund/factsheet/subscription-redemption-minimums
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| minimum_sub_ipo | float | Minimum initial subscription amount |
| minimum_sub_ipo_cur | string | Currency for minimum initial subscription |
| minimum_sub | float | Minimum subsequent subscription amount |
| minimum_sub_cur | string | Currency for minimum subsequent subscription |
| minimum_sub_unit | string | Minimum subscription units |
| minimum_redempt | float | Minimum redemption amount |
| minimum_redempt_cur | string | Currency for minimum redemption |
| minimum_redempt_unit | string | Minimum redemption units |
| lowbal_val | float | Minimum remaining balance amount |
| lowbal_val_cur | string | Currency for minimum balance |
| lowbal_unit | string | Minimum remaining unit balance |
| last_upd_date | datetime | Last update timestamp |

---

### 10. Subscription/Redemption Periods
```
GET /v2/fund/factsheet/subscription-redemption-periods
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| type | string | Transaction type (subscription or redemption) |
| period | string | Subscription/redemption period |
| redemp_period_oth | string | Additional description if period is 'Other' |
| settlement_period | string | Settlement period for redemption payment |
| last_upd_date | datetime | Last update timestamp |

---

### 12. Fund Statistics
```
GET /v2/fund/factsheet/statistics
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| portfolio_turnover_ratio | string | Portfolio Turnover Ratio |
| recovering_period | string | Recovering Period |
| portfolio_duration_period | string | Portfolio Duration for bond funds |
| maximum_drawdown | string | Maximum drawdown over past 5 years |
| sharpe_ratio | string | Sharpe Ratio (equity funds) |
| beta | string | Beta (equity funds) |
| alpha | string | Alpha (equity funds) |
| fx_hedging | string | FX Hedging |
| tracking_error | string | Tracking Error |
| yield_to_maturity | string | Yield to Maturity |
| last_upd_date | datetime | Last update timestamp |

---

### 13. Fund Dividend Policy
```
GET /v2/fund/factsheet/dividend-policy
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| fund_class_name | string | Fund class |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| dividend_policy | string | Dividend payment policy |
| last_upd_date | datetime | Last update timestamp |

---

### 15. Historical Fund Performance
```
GET /v2/fund/factsheet/performance
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| performance_type_desc | string | Performance type description |
| reference_period | string | Reference period and look-back year |
| performance_value | float | Historical performance value |

---

### 16. Asset Allocation
```
GET /v2/fund/factsheet/asset-allocation
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |
| fund_class_name | string | Filter by fund class |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| asset_seq | int | Asset sequence/order |
| asset_name | string | Asset type name (Thai) |
| asset_ratio | float | Allocation percentage |

---

### 11. Fund Risk Spectrum
```
GET /v2/fund/factsheet/risk-spectrum
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| risk_spectrum | string | Risk level (RS1-RS8, RS81) |
| risk_spectrum_desc | string | Risk level description |

---

### 17. Top 5 Holdings
```
GET /v2/fund/factsheet/top5-holdings
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_date | date | Factsheet effective start date |
| end_date | date | Factsheet effective end date |
| latest | boolean | Return only latest factsheet |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| start_date | date | Effective start date |
| end_date | date | Effective end date (null if current) |
| prospectus_type | string | IPO, Monthly, SignificantFactsheet |
| asset_seq | int | Holding rank (1-5) |
| asset_name | string | Asset/security name |
| asset_ratio | float | %NAV allocation |

---

## Outstanding Endpoints

### 18. Quarterly Fund Portfolio
```
GET /v2/fund/outstanding/portfolio
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_period | string | Starting period (YYYYMM) |
| end_period | string | Ending period (YYYYMM) |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| period | int | Reporting period (YYYYMM) |
| as_of_date | date | Portfolio as-of date |
| assetliab_id | string | Asset/liability category code |
| assetliab_desc | string | Category description |
| issue_code | string | Security symbol |
| isin_code | string | ISIN code |
| issuer | string | Issuer name |
| assetliab_value | number | Market value (THB) |
| percent_nav | number | %NAV |

---

### 19. Monthly Fund Portfolio by Asset Type
```
GET /v2/fund/outstanding/portfolio-asset-type
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_period | string | Starting period (YYYYMM) |
| end_period | string | Ending period (YYYYMM) |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| period | int | Reporting period (YYYYMM) |
| assetliab_code | string | Investment category code |
| assetliab_desc | string | Category description |
| market_value | number | Market value (THB) |
| percent_nav | number | %NAV |

---

### 21. Dividend History
```
GET /v2/fund/daily-info/dividend-history
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page_size | int | Items per page (1-100) |
| next_cursor | string | Pagination cursor |
| proj_id | string | Filter by project ID |
| start_dividend_date | date | Start dividend date (YYYY-MM-DD) |
| end_dividend_date | date | End dividend date (YYYY-MM-DD) |
| class_abbr_name | string | Filter by class abbreviation |

**Response Items:**
| Field | Type | Description |
|-------|------|-------------|
| proj_id | string | Project number |
| unique_id | string | AMC identifier |
| class_abbr_name | string | Class abbreviation |
| book_close_date | datetime | Book close date |
| dividend_date | datetime | Dividend payment date |
| dividend_value | float | Dividend amount per unit |

---

## Fund Status Values

| Status | Description |
|--------|-------------|
| Registered | Officially registered and active |
| IPO | Initial Public Offering period |
| Expired | Offering period ended |
| Canceled | Project canceled |
| Liquidated | Fund liquidated |

## Management Style Values

| Code | Description |
|------|-------------|
| AM | Active Management |
| AN | Thai fund follows master fund (Active) |
| PM | Passive Management/Index Tracking |
| PN | Thai fund follows master fund (Passive) |
| IM | Inverse Management |
| IN | Thai fund follows master fund (Inverse) |
| LM | Leveraged Management |
| LN | Thai fund follows master fund (Leveraged) |
| BH | Buy-and-Hold |
| SM | Index Tracking with occasional alpha |
| OT | Others |

## Retail Type Values

| Code | Description |
|------|-------------|
| A | Non-retail investors |
| B | High net worth investors |
| F | Liquidity support for corporate bonds |
| G | Government policy response |
| H | Non-retail and high net worth |
| N | Institutional investors |
| R | General investors |
| V | Provident funds |
| X | Institutional and ultra-high net worth |
