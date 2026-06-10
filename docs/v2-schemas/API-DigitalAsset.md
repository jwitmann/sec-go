# Digital Asset API (v1)

Base URL: `https://api.sec.or.th/DigitalAsset`

**Returns arrays directly** (no pagination envelope).

---

## Table of Contents

- [Digital Asset Intermediary Profiles](#digital-asset-intermediary-profiles)

---

## Digital Asset Intermediary Profiles

### 82. Digital Asset Intermediary Profiles

```http
POST /profile/intermediary
```

**Request Body:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `IntermediaryName` | string | No | Intermediary name to search. Pass empty string to list all. |

```json
{
  "IntermediaryName": ""
}
```

**Response Items (`DigitalAssetIntermediary`):**

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `unique_id` | string | No | Unique identifier |
| `name_th` | string | No | Name in Thai |
| `name_en` | string | No | Name in English |
| `lic_code` | string | No | License code |
| `lic_action_code` | string | No | License action code |
| `lic_efft_date` | date | No | License effective date |
| `lic_act_date` | date | No | License action date |
| `lic_exp_date` | date | No | License expiration date |
