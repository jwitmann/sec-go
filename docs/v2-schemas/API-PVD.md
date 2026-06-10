# PVD API (v1)

Base URL: `https://api.sec.or.th`

All PVD endpoints use the standard [paginated envelope](../API.md#pagination) with `message`, `page_size`, `next_cursor`, and `items`.

Uses standard pagination parameters (`page_size`, `next_cursor`) on all endpoints.

---

## Table of Contents

- [General Information](#general-information)
- [Fund Members](#fund-members)
- [Fund Assets](#fund-assets)
- [Fund Transactions](#fund-transactions)
- [Fund Contributions](#fund-contributions)
- [Fund Expenses](#fund-expenses)
- [Fund Liquidity](#fund-liquidity)
- [Fund Performance](#fund-performance)
- [Fund Benchmarks](#fund-benchmarks)
- [Fund Dividends](#fund-dividends)
- [Fund Policies](#fund-policies)
- [Fund Fees](#fund-fees)
- [Fund Compliance](#fund-compliance)

---

## General Information

### 28. List Provident Funds

```http
GET /v1/pvd/general-info/list?page_size={page_size}&next_cursor={next_cursor}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |

**Response Items (`PVDListItem`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `unique_id` | string | No | Fund identifier |
| `comp_name_th` | string | No | Company name in Thai |
| `comp_name_en` | string | No | Company name in English |
| `last_upd_date` | datetime | No | Last update timestamp |

---

### 29. PVD Fund Information

```http
GET /v1/pvd/general-info/fund-info?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |

**Response Items (`PVDFundInfo`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `proj_name_th` | string | No | Fund name in Thai |
| `proj_name_en` | string | No | Fund name in English |
| `proj_abbr_name` | string | No | Fund abbreviation |
| `fund_status` | string | No | Fund status |
| `init_date` | date | No | Establishment date |
| `regis_date` | date | No | Registration date |
| `cancel_date` | date | No | Cancellation date |
| `policy_desc` | string | No | Policy description |
| `management_style` | string | No | Management style code |
| `fund_class_name` | string | No | Fund class name |
| `last_upd_date` | datetime | No | Last update timestamp |

---

### 30. PVD Fund Specifications

```http
GET /v1/pvd/general-info/fund-spec?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:** Same as endpoint 29.

**Response Items (`PVDFundSpec`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `spec_code` | string | No | Specification code |
| `spec_desc` | string | No | Specification description |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Members

### 31. PVD Fund Members

```http
GET /v1/pvd/fund-member?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |
| `start_period` | period | O | — | Starting period (`YYYYMM`) |
| `end_period` | period | O | — | Ending period (`YYYYMM`) |

**Response Items (`PVDFundMember`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `total_members` | int | No | Total members |
| `active_members` | int | No | Active members |
| `retired_members` | int | No | Retired members |
| `last_upd_date` | date | No | Last update date |

---

## Fund Assets

### 32. PVD Fund Assets

```http
GET /v1/pvd/fund-asset?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundAsset`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `asset_seq` | int | No | Asset sequence |
| `asset_name` | string | No | Asset name |
| `asset_ratio` | float | No | Asset ratio (%) |
| `market_value` | float | No | Market value (THB) |
| `last_upd_date` | date | No | Last update date |

---

## Fund Transactions

### 33. PVD Fund Transactions

```http
GET /v1/pvd/fund-transaction?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundTransaction`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `trans_type` | string | No | Transaction type |
| `trans_value` | float | No | Transaction value (THB) |
| `trans_count` | int | No | Transaction count |
| `last_upd_date` | date | No | Last update date |

---

## Fund Contributions

### 34. PVD Fund Contributions

```http
GET /v1/pvd/fund-contribution?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundContribution`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `employee_contrib` | float | No | Employee contribution (THB) |
| `employer_contrib` | float | No | Employer contribution (THB) |
| `total_contrib` | float | No | Total contribution (THB) |
| `last_upd_date` | date | No | Last update date |

---

## Fund Expenses

### 35. PVD Fund Expenses

```http
GET /v1/pvd/fund-expense?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundExpense`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `expense_type` | string | No | Expense type |
| `expense_value` | float | No | Expense value (THB) |
| `expense_ratio` | float | No | Expense ratio (%) |
| `last_upd_date` | date | No | Last update date |

---

## Fund Liquidity

### 36. PVD Fund Liquidity

```http
GET /v1/pvd/fund-liquidity?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundLiquidity`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `cash_ratio` | float | No | Cash ratio |
| `current_ratio` | float | No | Current ratio |
| `quick_ratio` | float | No | Quick ratio |
| `last_upd_date` | date | No | Last update date |

---

## Fund Performance

### 37. PVD Fund Performance

```http
GET /v1/pvd/fund-performance?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundPerformance`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `performance_type` | string | No | Performance type |
| `reference_period` | string | No | Reference period |
| `performance_value` | float | No | Performance value |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Benchmarks

### 38. PVD Fund Benchmarks

```http
GET /v1/pvd/fund-benchmark?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundBenchmark`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `benchmark` | string | No | Benchmark name |
| `benchmark_value` | float | No | Benchmark value |
| `fund_value` | float | No | Fund value |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Dividends

### 39. PVD Fund Dividends

```http
GET /v1/pvd/fund-dividend?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |

**Response Items (`PVDFundDividend`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `dividend_date` | date | No | Dividend date |
| `dividend_value` | float | No | Dividend value per unit |
| `book_close_date` | date | No | Book close date |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Policies

### 40. PVD Fund Policies

```http
GET /v1/pvd/fund-policy?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:** Same as endpoint 39.

**Response Items (`PVDFundPolicy`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `policy_type` | string | No | Policy type |
| `policy_desc` | string | No | Policy description |
| `min_equity_ratio` | float | No | Minimum equity ratio (%) |
| `max_equity_ratio` | float | No | Maximum equity ratio (%) |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Fees

### 41. PVD Fund Fees

```http
GET /v1/pvd/fund-fee?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:** Same as endpoint 39.

**Response Items (`PVDFundFee`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `fund_class_name` | string | No | Fund class |
| `fee_type_desc` | string | No | Fee type description |
| `rate` | float | No | Fee rate |
| `rate_unit` | string | No | Rate unit |
| `actual_value` | float | No | Actual fee value |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Fund Compliance

### 42. PVD Fund Compliance

```http
GET /v1/pvd/fund-compliance?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 31.

**Response Items (`PVDFundCompliance`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `compliance_type` | string | No | Compliance type |
| `compliance_status` | string | No | Compliance status |
| `remark` | string | No | Remark |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Convenience Helpers

The Go client provides higher-level helpers for common PVD operations:

### Fetch All Provident Funds

Auto-paginates through all pages:

```go
pvds, err := client.FetchAllPVDs(ctx)
```

### Find Provident Fund

Search by Thai name, English name, or unique ID:

```go
pvd, err := client.FindPVD(ctx, "กรุงศรี")
// or
pvd, err := client.FindPVD(ctx, "PVD123456")
```

### Fetch All PVD Fund Info

Auto-paginates through all pages for a given project:

```go
fundInfos, err := client.FetchAllPVDFundInfo(ctx, sec.PVDProjOptions{ProjID: "PVD123456"})
```
