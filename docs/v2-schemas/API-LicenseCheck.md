# License Check API (v1)

Base URL: `https://api.sec.or.th`

All License Check endpoints use the standard [paginated envelope](../API.md#pagination) with `message`, `page_size`, `next_cursor`, and `items`.

Uses standard pagination parameters (`page_size`, `next_cursor`) on all endpoints.

---

## Table of Contents

- [Common Parameters](#common-parameters)
- [Common Response Structure](#common-response-structure)
- [Endpoints](#endpoints)

---

## Common Parameters

Most endpoints support these filters:

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page_size` | int | O | 100 | Items per page (1-100) |
| `next_cursor` | string | O | `""` | Pagination cursor |
| `license_no` | string | O | — | Filter by license number |
| `entity_name` | string | O | — | Filter by entity name (Thai/English) |
| `license_status` | string | O | — | Filter by license status |

---

## Common Response Structure

All endpoints return `LicenseCheckResult` items:

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `license_id` | string | No | License identifier |
| `license_no` | string | No | License number |
| `license_type` | string | No | License type |
| `entity_name_th` | string | Yes | Entity name in Thai |
| `entity_name_en` | string | Yes | Entity name in English |
| `license_status` | string | No | License status |
| `issue_date` | date | No | Issue date |
| `expire_date` | date | No | Expiration date |
| `license_detail` | string | Yes | License detail |
| `remark` | string | Yes | Remark |
| `last_upd_date` | datetime | No | Last update timestamp |

---

## Endpoints

### 43. Check Bond Sale Representatives

```http
GET /v1/license-check/bond-sale-rep?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&rep_name={rep_name}
```

**Additional Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `rep_name` | string | O | — | Filter by representative name |

---

### 44. Check Securities Companies

```http
GET /v1/license-check/securities-company?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&company_name={company_name}
```

**Additional Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `company_name` | string | O | — | Filter by company name |

---

### 45. Check Derivatives Companies

```http
GET /v1/license-check/derivatives-company?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 46. Check Securities Brokers

```http
GET /v1/license-check/securities-broker?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 47. Check Derivatives Brokers

```http
GET /v1/license-check/derivatives-broker?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 48. Check Investment Advisors

```http
GET /v1/license-check/investment-advisor?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 49. Check Securities Fund Managers

```http
GET /v1/license-check/securities-fund-manager?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 50. Check Fund Supervisors

```http
GET /v1/license-check/fund-supervisor?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 51. Check Auditors

```http
GET /v1/license-check/auditor?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 52. Check Credit Rating Companies

```http
GET /v1/license-check/credit-rating-company?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 53. Check Private Funds

```http
GET /v1/license-check/private-fund?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 54. Check Derivatives Fund Managers

```http
GET /v1/license-check/derivatives-fund-manager?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 55. Check Securities Borrowing & Lending

```http
GET /v1/license-check/securities-borrowing-lending?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 56. Check Financial Advisors

```http
GET /v1/license-check/financial-advisor?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 57. Check Acquirers

```http
GET /v1/license-check/acquirer?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```

---

### 58. Check Venture Capitals

```http
GET /v1/license-check/venture-capital?page_size={page_size}&next_cursor={next_cursor}&license_no={license_no}&entity_name={entity_name}&license_status={license_status}
```
