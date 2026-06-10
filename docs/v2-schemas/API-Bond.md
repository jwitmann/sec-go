# Bond API (v2)

Base URL: `https://api.sec.or.th`

All Bond endpoints use the standard [paginated envelope](../API.md#pagination) with `message`, `page_size`, `next_cursor`, and `items`.

Uses standard pagination parameters (`page_size`, `next_cursor`) on all endpoints.

---

## Table of Contents

- [General Information](#general-information)
- [Credit Ratings](#credit-ratings)
- [Outstanding](#outstanding)
- [Related Parties](#related-parties)
- [Investor Holdings](#investor-holdings)

---

## General Information

### 22. List Bond Issuers (AMCs)

```http
GET /v2/bond/general-info/amcs?page_size={page_size}&next_cursor={next_cursor}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |

**Response Items (`BondIssuer`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `unique_id` | string | No | AMC identifier (e.g., `C0000000021`) |
| `comp_name_th` | string | No | Company name in Thai |
| `comp_name_en` | string | No | Company name in English |
| `last_upd_date` | datetime | No | Last update timestamp |

---

### 23. Bond Features

```http
GET /v2/bond/general-info/features?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |

**Response Items (`BondFeature`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `proj_name_th` | string | No | Project name in Thai |
| `proj_name_en` | string | Yes | Project name in English |
| `proj_abbr_name` | string | No | Project abbreviation |
| `issue_date` | date | No | Issue date (`YYYY-MM-DD`) |
| `maturity_date` | date | No | Maturity date (`YYYY-MM-DD`) |
| `coupon_rate` | float | No | Coupon rate (%) |
| `face_value` | float | No | Face value per unit (THB) |
| `issue_value` | float | No | Total issue value (THB) |
| `bond_type` | string | No | Bond type code |
| `bond_sub_type` | string | No | Bond sub-type code |
| `secured_flag` | string | No | Secured flag (`Y`/`N`) |
| `guarantee_flag` | string | No | Guarantee flag (`Y`/`N`) |
| `collateral_desc` | string | Yes | Collateral description |
| `purpose_th` | string | Yes | Purpose in Thai |
| `purpose_en` | string | Yes | Purpose in English |
| `remark_th` | string | Yes | Remarks in Thai |
| `remark_en` | string | Yes | Remarks in English |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Credit Ratings

### 24. Bond Credit Ratings

```http
GET /v2/bond/credit-rating?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `proj_id` | string | O | — | Project number |
| `start_period` | period | O | — | Starting period (`YYYYMM`) |
| `end_period` | period | O | — | Ending period (`YYYYMM`) |

**Response Items (`BondCreditRating`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `rating_agency` | string | No | Rating agency name |
| `rating` | string | No | Credit rating |
| `outlook` | string | No | Rating outlook |
| `as_of_date` | date | No | Rating as-of date |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Outstanding

### 25. Bond Outstanding

```http
GET /v2/bond/outstanding?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 24.

**Response Items (`BondOutstanding`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `outstanding_qty` | float | No | Outstanding quantity |
| `outstanding_val` | float | No | Outstanding value (THB) |
| `held_by_investor` | float | No | Held by investors |
| `held_by_issuer` | float | No | Held by issuer |
| `last_upd_date` | date | No | Last update date |

---

## Related Parties

### 26. Bond Related Parties

```http
GET /v2/bond/related-party?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 24.

**Response Items (`BondRelatedParty`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `party_type` | string | No | Party type |
| `party_name_th` | string | No | Party name in Thai |
| `party_name_en` | string | Yes | Party name in English |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Investor Holdings

### 27. Bond Investor Holdings

```http
GET /v2/bond/investor-holding?page_size={page_size}&next_cursor={next_cursor}&proj_id={proj_id}&start_period={start_period}&end_period={end_period}
```

**Query Parameters:** Same as endpoint 24.

**Response Items (`BondInvestorHolding`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `proj_id` | string | No | Project number |
| `period` | int | No | Reporting period (`YYYYMM`) |
| `investor_type` | string | No | Investor type |
| `holding_qty` | float | No | Holding quantity |
| `holding_val` | float | No | Holding value (THB) |
| `holding_pct` | float | No | Holding percentage |
| `last_upd_date` | date | No | Last update date |

---

## Convenience Helpers

The Go client provides higher-level helpers for common Bond operations:

### Fetch All Bond Issuers

Auto-paginates through all pages:

```go
issuers, err := client.FetchAllBondIssuers(ctx)
```

### Find Bond Issuer

Search by Thai name, English name, or unique ID:

```go
issuer, err := client.FindBondIssuer(ctx, "กรุงไทย")
// or
issuer, err := client.FindBondIssuer(ctx, "C0000000021")
```

### Fetch All Bond Features

Auto-paginates through all pages for a given project:

```go
features, err := client.FetchAllBondFeatures(ctx, sec.BondFeatureOptions{ProjID: "B123456"})
```
