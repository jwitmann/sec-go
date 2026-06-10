# One Report API (v1)

Base URL: `https://api.sec.or.th/onereport`

**All endpoints return arrays directly** (no pagination envelope). All are `GET` requests with path parameters.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `report_year` | string | Yes | Report year (e.g., `2024`) |
| `unique_id` | string | Yes | Company unique identifier |
| `language` | string | Yes | Language code: `T` for Thai, `E` for English |

---

## Table of Contents

- [Common Response Fields](#common-response-fields)
- [SBO (Small Business Overview)](#sbo-small-business-overview)
- [Sustainability](#sustainability)
- [Financial Statement](#financial-statement)
- [CGP (Corporate Governance Policy)](#cgp-corporate-governance-policy)
- [CGS (Corporate Governance Structure)](#cgs-corporate-governance-structure)
- [SCP (Social Performance)](#scp-social-performance)

---

## Common Response Fields

Most One Report response items contain at minimum:

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

Additional fields are endpoint-specific. Detailed schemas for many endpoints are still being finalized by SEC.

---

## SBO (Small Business Overview)

### 59. SBO Company Information

```http
GET /sbo/{report_year}/info/{language}
```

**Response Items (`SBOInfo`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |
| `language` | string | No | Language code |
| `corp_name` | string | No | Corporate name |
| `symbol` | string | No | Stock symbol |
| `address` | string | No | Address |
| `province` | string | No | Province |
| `zip_code` | string | No | ZIP code |
| `business_type` | string | No | Business type |
| `registered_number` | string | No | Registered number |
| `telephone` | string | No | Telephone |
| `fax` | string | No | Fax number |
| `website` | string | No | Website URL |
| `email` | string | No | Email address |
| `common_paidup_share` | float | No | Common paid-up shares |
| `preferred_paidup_share` | float | No | Preferred paid-up shares |

---

### 60. SBO R&D Information

```http
GET /sbo/{report_year}/rd/{unique_id}
```

**Response Items (`SBORD`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 61. SBO Product Income

```http
GET /sbo/{report_year}/product_income/{unique_id}
```

**Response Items (`SBOProductIncome`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |
| `business_income_code` | string | No | Business income code |
| `sequence` | int | No | Sequence number |
| `business_income_desc` | string | No | Business income description |
| `asof_year` | float | No | Current year value |
| `asof_yesteryear` | float | No | Previous year value |
| `asof_year_before_yesteryear` | float | No | Two years ago value |
| `asof_year_percent` | float | No | Current year percentage |
| `asof_yesteryear_percent` | float | No | Previous year percentage |
| `asof_year_before_yesteryear_percent` | float | No | Two years ago percentage |

---

### 62. SBO Export Income

```http
GET /sbo/{report_year}/export_income/{unique_id}
```

**Response Items (`SBOExportIncome`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 63. SBO Risk Details

```http
GET /sbo/{report_year}/risk/{unique_id}
```

**Response Items (`SBORisk`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |
| `risk_category` | string | No | Risk category |
| `risk_code` | string | No | Risk code |
| `sequence` | int | No | Sequence number |
| `choice` | string | No | Choice |
| `holder_risk` | string | No | Holder risk |
| `foreign_risk` | string | No | Foreign risk |

---

## Sustainability

### 64. Sustainability Detail

```http
GET /sustainability/{report_year}/detail/{unique_id}
```

**Response Items (`SustainabilityDetail`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 65. Sustainability Environment Issue

```http
GET /sustainability/{report_year}/environment_issue/{unique_id}
```

**Response Items (`SustainabilityEnvironmentIssue`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 66. Sustainability Human Rights Issue

```http
GET /sustainability/{report_year}/humanrights_issue/{unique_id}
```

**Response Items (`SustainabilityHumanRightsIssue`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

## Financial Statement

### 67. Financial Statement

```http
GET /fs/{report_year}/financial_statement/{unique_id}
```

**Response Items (`FinancialStatement`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

## CGP (Corporate Governance Policy)

### 68. CGP Governance

```http
GET /cgp/{report_year}/governance/{unique_id}
```

**Response Items (`CGPGovernance`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 69. CGP Director

```http
GET /cgp/{report_year}/director/{unique_id}
```

**Response Items (`CGPDirector`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 70. CGP Code of Conduct

```http
GET /cgp/{report_year}/code_of_conduct/{unique_id}
```

**Response Items (`CGPCodeOfConduct`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

## CGS (Corporate Governance Structure)

### 71. CGS Board

```http
GET /cgs/{report_year}/board/{unique_id}
```

**Response Items (`CGSBoard`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 72. CGS Employee

```http
GET /cgs/{report_year}/employee/{unique_id}
```

**Response Items (`CGSEmployee`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 73. CGS Auditor Company

```http
GET /cgs/{report_year}/auditor_company/{unique_id}
```

**Response Items (`CGSAuditorCompany`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 74. CGS Director Performance

```http
GET /cgs/{report_year}/director_performance/{unique_id}
```

**Response Items (`CGSDirectorPerformance`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 75. CGS Board of Directors (BODs)

```http
GET /cgs/{report_year}/bods/{unique_id}
```

**Response Items (`CGSBods`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 76. CGS Executives

```http
GET /cgs/{report_year}/executives/{unique_id}
```

**Response Items (`CGSExecutives`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 77. CGS Committees Others

```http
GET /cgs/{report_year}/committees/{unique_id}/others
```

**Response Items (`CGSCommitteesOthers`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

## SCP (Social Performance)

### 78. SCP Labor Dispute

```http
GET /scp/{report_year}/labor_dispute/{unique_id}
```

**Response Items (`SCPLaborDispute`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 79. SCP CSR Activity

```http
GET /scp/{report_year}/csr_activity/{unique_id}
```

**Response Items (`SCPCSRActivity`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 80. SCP Employee Info

```http
GET /scp/{report_year}/employee_info/{unique_id}
```

**Response Items (`SCPEmployeeInfo`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |

---

### 81. SCP Employee Development

```http
GET /scp/{report_year}/employee_development/{unique_id}
```

**Response Items (`SCPEmployeeDevelopment`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `last_upd_date` | datetime | No | Last update timestamp |
| `report_year` | string | No | Report year |
| `unique_id` | string | No | Company unique identifier |
