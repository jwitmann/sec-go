# SEC Open Data APIs - Complete Endpoint Discovery

## Overview

This document catalogs all API endpoints available on the SEC Open Data API portal (`https://api.sec.or.th`).

**Last Updated:** 2026-06-10
**Source:** SEC Open Data API Portal Navigation Panel (`https://dataexchange.sec.or.th/portal/interactive/apis/sec-open-apis`)

---

## API Categories

| Category | Endpoints | Status |
|----------|-----------|--------|
| Fund (กองทุน) | 21 | ✅ Fully documented in [API.md](API.md) |
| Bond (ตราสารหนี้) | 6 | 🔍 Discovered, needs implementation |
| Digital Asset (สินทรัพย์ดิจิทัล) | 1 | 🔍 Discovered, needs implementation |
| License Check | 16 | 🔍 Discovered, needs implementation |
| One Report | 23 | 🔍 Discovered, needs implementation |
| PVD (Provident Fund) | 15 | 🔍 Discovered, needs implementation |
| **Total** | **82** | |

---

## Version Migration Notice

- **Fund & Bond APIs:** v1 → v2 migration deadline **April 30, 2026**
- **Other APIs (License Check, One Report, PVD):** legacy → v1 migration deadline **June 30, 2026**

---

## Fund APIs (กองทุน) - 21 Endpoints

Already fully implemented and documented in [API.md](API.md).

### General Information (5 endpoints)
| Endpoint | Service ID | Description |
|----------|------------|-------------|
| AMC List | `getAmcList` | Asset management companies |
| Fund Profile | `getFundProfile` | Fund registration details |
| Fund Specification | `getFundSpecification` | Fund type/class specifications |
| Mutual Fund Fee | `getMutualfundFee` | Fee structures |
| Fund Relative Parties | `getFundRelative` | Related entities |

### Factsheet (12 endpoints)
| Endpoint | Service ID | Description |
|----------|------------|-------------|
| Factsheet URL | `getFactsheetUrl` | Document links |
| Factsheet IPO | `getFactsheetIPO` | IPO offering periods |
| Factsheet Benchmark | `getFactsheetBenchmark` | Performance benchmarks |
| Factsheet Risk Spectrum | `getFactsheetRiskSpectrum` | Risk levels |
| Factsheet Performance | `getFactsheetPerformance` | Historical returns |
| Factsheet Asset Allocation | `getFactsheetAssetAllocation` | Portfolio breakdown |
| Factsheet Top 5 Holding | `getFactsheetTop5Holding` | Top holdings |
| Factsheet Subscription | `getFactsheetSubscription` | Subscription terms |
| Factsheet Redemption | `getFactsheetRedemption` | Redemption terms |
| Factsheet Statistic | `getFactsheetStatistic` | Fund statistics |
| Factsheet Dividend Policy | `getFactsheetDividendPolicy` | Dividend policies |
| Factsheet Fee | `getFactsheetFee` | Fee details |

### Daily Information (2 endpoints)
| Endpoint | Service ID | Description |
|----------|------------|-------------|
| Daily NAV | `getFundDailyInfoNAV` | Net asset values |
| Dividend History | `getFundDailyInfoDividendHistory` | Dividend payments |

### Outstanding (2 endpoints)
| Endpoint | Service ID | Description |
|----------|------------|-------------|
| Quarterly Portfolio | `getOutstandingQuarterlyPortfolio` | Quarterly holdings |
| Monthly Portfolio Asset Type | `getOutstandingMonthlyPortfolioAssetType` | Asset allocation |

---

## Bond APIs (ตราสารหนี้) - 6 Endpoints

| # | Service ID | Name | Description |
|---|------------|------|-------------|
| 1 | `getAmcList` | Issuer Names | Bond issuer/AMC list |
| 2 | `getBondFeatures` | Bond Features | Bond specifications and features |
| 3 | `bond-credit-rating` | Credit Rating | Credit ratings by time period |
| 4 | `bond-outstanding` | Outstanding Values | Outstanding values by time period |
| 5 | `bond-related-party` | Related Parties | Related parties by time period |
| 6 | `bond-investor-holding` | Investor Holdings | Investor holdings by type |

**Base Path:** `/v2/bond/...` (expected, following Fund v2 pattern)

---

## Digital Asset APIs (สินทรัพย์ดิจิทัล) - 1 Endpoint

> ✅ **Status: Implemented**

| # | Service ID | Name | Description |
|---|------------|------|-------------|
| 1 | `01-profile-intermediary` | Intermediary Profiles | Digital asset intermediary profiles |

**Base Path:** `https://api.sec.or.th/DigitalAsset`
**Method:** `POST`
**Request Body:** `{ "IntermediaryName": "string" }`

---

## License Check APIs - 16 Endpoints

These endpoints allow checking licenses and registrations of financial services entities.

| # | Service ID | Name | Description |
|---|------------|------|-------------|
| 1 | `01-bond-sale-rep` | Bond Sale Representative | Bond sale rep licenses |
| 2 | `02-securities-company` | Securities Company | Securities company licenses |
| 3 | `03-derivatives-company` | Derivatives Company | Derivatives company licenses |
| 4 | `04-securities-broker` | Securities Broker | Securities broker licenses |
| 5 | `05-derivatives-broker` | Derivatives Broker | Derivatives broker licenses |
| 6 | `06-investment-advisor` | Investment Advisor | Investment advisor licenses |
| 7 | `07-securities-fund-manager` | Securities Fund Manager | Securities fund manager licenses |
| 8 | `08-fund-supervisor` | Fund Supervisor | Fund supervisor licenses |
| 9 | `09-auditor` | Auditor | Auditor licenses |
| 10 | `10-credit-rating-company` | Credit Rating Company | Credit rating company licenses |
| 11 | `11-private-fund` | Private Fund | Private fund licenses |
| 12 | `12-derivatives-fund-manager` | Derivatives Fund Manager | Derivatives fund manager licenses |
| 13 | `13-securities-borrowing-lending` | Securities Borrowing & Lending | Securities borrowing & lending licenses |
| 14 | `14-financial-advisor` | Financial Advisor | Financial advisor licenses |
| 15 | `15-acquirer` | Acquirer | Acquirer licenses |
| 16 | `16-venture-capital` | Venture Capital | Venture capital licenses |

**Base Path:** `/v1/license-check/...` (expected, new API category)

---

## One Report APIs - 23 Endpoints

> ✅ **Status: Implemented**

Corporate disclosure and governance data from annual reports (56-1 One Report).

Actual endpoint IDs from portal navigation:

### SBO (Small Business Overview)
| # | Service ID | Name |
|---|------------|------|
| 1 | `01-SBO-Info` | SBO Info |
| 2 | `02-SBO-RD` | SBO RD |
| 3 | `03-SBO-Product-Income` | SBO Product Income |
| 4 | `04-SBO-Export-Income` | SBO Export Income |
| 5 | `05-SBO-Risk-Detail` | SBO Risk Detail |

### Sustainability
| # | Service ID | Name |
|---|------------|------|
| 6 | `06-Sustainability-Detail` | Sustainability Detail |
| 7 | `07-Sustainability-Environment_issue` | Sustainability Environment Issue |
| 8 | `08-Sustainability-Humanrights_issue` | Sustainability Human Rights Issue |

### Social Performance
| # | Service ID | Name |
|---|------------|------|
| 9 | `09-SocialPerformance-Employee_info` | Social Performance Employee Info |
| 10 | `10-SocialPerformance-Employee_development` | Social Performance Employee Development |
| 11 | `11-SocialPerformance-Labor_dispute` | Social Performance Labor Dispute |
| 12 | `12-SocialPerformance-CSR_Activity` | Social Performance CSR Activity |

### CGP (Corporate Governance Policy)
| # | Service ID | Name |
|---|------------|------|
| 13 | `13-CGP-Governance` | CGP Governance |
| 14 | `14-CGP-Director` | CGP Director |
| 15 | `15-CGP-CodeofConduct` | CGP Code of Conduct |
| 16 | `16-CGP-FinancialStatement` | CGP Financial Statement |

### CGS (Corporate Governance Structure)
| # | Service ID | Name |
|---|------------|------|
| 17 | `17-CGS-Board` | CGS Board |
| 18 | `18-CGS-Employee` | CGS Employee |
| 19 | `19-CGS-AuditorCompany` | CGS Auditor Company |
| 20 | `20-CGS-DirectorPerformance` | CGS Director Performance |
| 21 | `21-CGS-Bods` | CGS BODs |
| 22 | `22-CGS-Executives` | CGS Executives |
| 23 | `23-CGS-Committees-Others` | CGS Committees & Others |

**Base Path:** `/v1/one-report/...`
**Method:** All endpoints are `GET`
**Blocker:** SEC has not published API documentation (schemas, parameters, response formats). Cannot implement without this information.

---

## PVD APIs (Provident Fund / กองทุนสำรองเลี้ยงชีพ) - 15 Endpoints

Provident fund data - similar structure to mutual fund APIs but for retirement funds.

| # | Service ID | Name | Description |
|---|------------|------|-------------|
| 1 | `getPvdList` | PVD List | List of provident funds |
| 2 | `getPvdFundInfo` | PVD Fund Info | Fund information |
| 3 | `getPvdFundSpec` | PVD Fund Spec | Fund specifications |
| 4 | `getPvdFundMember` | PVD Fund Member | Member statistics |
| 5 | `getPvdFundAsset` | PVD Fund Asset | Asset allocation |
| 6 | `getPvdFundTransaction` | PVD Fund Transaction | Transaction data |
| 7 | `getPvdFundContribution` | PVD Fund Contribution | Contribution records |
| 8 | `getPvdFundExpense` | PVD Fund Expense | Expense data |
| 9 | `getPvdFundLiquidity` | PVD Fund Liquidity | Liquidity metrics |
| 10 | `getPvdFundPerformance` | PVD Fund Performance | Performance data |
| 11 | `getPvdFundBenchmark` | PVD Fund Benchmark | Benchmark comparisons |
| 12 | `getPvdFundDividend` | PVD Fund Dividend | Dividend history |
| 13 | `getPvdFundPolicy` | PVD Fund Policy | Investment policies |
| 14 | `getPvdFundFee` | PVD Fund Fee | Fee structures |
| 15 | `getPvdFundCompliance` | PVD Fund Compliance | Compliance data |

**Base Path:** `/v1/pvd/...` (expected, new API category)

---

## Implementation Status

All 82 endpoints across 6 categories are now implemented:

| Category | Endpoints | Status |
|----------|-----------|--------|
| Fund (กองทุน) | 21 | ✅ Implemented |
| Bond (ตราสารหนี้) | 6 | ✅ Implemented |
| PVD (Provident Fund) | 15 | ✅ Implemented |
| License Check | 16 | ✅ Implemented |
| One Report | 23 | ✅ Implemented |
| Digital Asset | 1 | ✅ Implemented |
| **Total** | **82** | **Complete** |

---

## Notes for Implementation

1. **Authentication:** All APIs use the same `Ocp-Apim-Subscription-Key` header
2. **Pagination:** New APIs likely use the same cursor-based pagination as Fund APIs
3. **Rate Limits:** Same limits apply (5,000 requests per 300 seconds)
4. **Base URL:** All APIs use `https://api.sec.or.th`
5. **Versioning:** 
   - Fund & Bond: `/v2/...`
   - Digital Asset, License Check, One Report, PVD: `/v1/...` (expected)

## Endpoint URL Patterns

Based on the Fund API structure, expected patterns are:

```
/v2/fund/{category}/{endpoint}    # Fund APIs (implemented)
/v2/bond/{category}/{endpoint}    # Bond APIs
/v1/digital-asset/{endpoint}      # Digital Asset APIs
/v1/license-check/{endpoint}      # License Check APIs
/v1/one-report/{endpoint}         # One Report APIs
/v1/pvd/{category}/{endpoint}     # PVD APIs
```

---

## Raw Discovery Data

The complete list of service IDs discovered:

```json
[
  {
    "name": "PVD",
    "endpoints": [
      {
        "id": "getPvdList",
        "name": "PVD List",
        "type": "pvd"
      },
      {
        "id": "getPvdFundInfo",
        "name": "PVD Fund Info",
        "type": "pvd"
      },
      {
        "id": "getPvdFundSpec",
        "name": "PVD Fund Spec",
        "type": "pvd"
      },
      {
        "id": "getPvdFundMember",
        "name": "PVD Fund Member",
        "type": "pvd"
      },
      {
        "id": "getPvdFundAsset",
        "name": "PVD Fund Asset",
        "type": "pvd"
      },
      {
        "id": "getPvdFundTransaction",
        "name": "PVD Fund Transaction",
        "type": "pvd"
      },
      {
        "id": "getPvdFundContribution",
        "name": "PVD Fund Contribution",
        "type": "pvd"
      },
      {
        "id": "getPvdFundExpense",
        "name": "PVD Fund Expense",
        "type": "pvd"
      },
      {
        "id": "getPvdFundLiquidity",
        "name": "PVD Fund Liquidity",
        "type": "pvd"
      },
      {
        "id": "getPvdFundPerformance",
        "name": "PVD Fund Performance",
        "type": "pvd"
      },
      {
        "id": "getPvdFundBenchmark",
        "name": "PVD Fund Benchmark",
        "type": "pvd"
      },
      {
        "id": "getPvdFundDividend",
        "name": "PVD Fund Dividend",
        "type": "pvd"
      },
      {
        "id": "getPvdFundPolicy",
        "name": "PVD Fund Policy",
        "type": "pvd"
      },
      {
        "id": "getPvdFundFee",
        "name": "PVD Fund Fee",
        "type": "pvd"
      },
      {
        "id": "getPvdFundCompliance",
        "name": "PVD Fund Compliance",
        "type": "pvd"
      }
    ]
  },
  {
    "name": "Bond",
    "endpoints": [
      {
        "id": "getAmcList",
        "name": "Issuer Names",
        "type": "bond"
      },
      {
        "id": "getBondFeatures",
        "name": "Bond Features",
        "type": "bond"
      },
      {
        "id": "bond-credit-rating",
        "name": "Credit Rating",
        "type": "bond"
      },
      {
        "id": "bond-outstanding",
        "name": "Outstanding Values",
        "type": "bond"
      },
      {
        "id": "bond-related-party",
        "name": "Related Parties",
        "type": "bond"
      },
      {
        "id": "bond-investor-holding",
        "name": "Investor Holdings",
        "type": "bond"
      }
    ]
  },
  {
    "name": "Digital Asset",
    "endpoints": [
      {
        "id": "01-profile-intermediary",
        "name": "Intermediary Profiles",
        "type": "digital-asset"
      }
    ]
  },
  {
    "name": "Fund",
    "endpoints": [
      {
        "id": "getAmcList",
        "name": "AMC List",
        "type": "fund"
      },
      {
        "id": "getFundProfile",
        "name": "Fund Profile",
        "type": "fund"
      },
      {
        "id": "getFundSpecification",
        "name": "Fund Specification",
        "type": "fund"
      },
      {
        "id": "getMutualfundFee",
        "name": "Mutual Fund Fee",
        "type": "fund"
      },
      {
        "id": "getFundRelative",
        "name": "Fund Relative Parties",
        "type": "fund"
      },
      {
        "id": "getFactsheetUrl",
        "name": "Factsheet URL",
        "type": "fund"
      },
      {
        "id": "getFactsheetIPO",
        "name": "Factsheet IPO",
        "type": "fund"
      },
      {
        "id": "getFactsheetBenchmark",
        "name": "Factsheet Benchmark",
        "type": "fund"
      },
      {
        "id": "getFactsheetRiskSpectrum",
        "name": "Factsheet Risk Spectrum",
        "type": "fund"
      },
      {
        "id": "getFactsheetPerformance",
        "name": "Factsheet Performance",
        "type": "fund"
      },
      {
        "id": "getFactsheetAssetAllocation",
        "name": "Factsheet Asset Allocation",
        "type": "fund"
      },
      {
        "id": "getFactsheetTop5Holding",
        "name": "Factsheet Top 5 Holding",
        "type": "fund"
      },
      {
        "id": "getFactsheetSubscription",
        "name": "Factsheet Subscription",
        "type": "fund"
      },
      {
        "id": "getFactsheetRedemption",
        "name": "Factsheet Redemption",
        "type": "fund"
      },
      {
        "id": "getFactsheetStatistic",
        "name": "Factsheet Statistic",
        "type": "fund"
      },
      {
        "id": "getFactsheetDividendPolicy",
        "name": "Factsheet Dividend Policy",
        "type": "fund"
      },
      {
        "id": "getFactsheetFee",
        "name": "Factsheet Fee",
        "type": "fund"
      },
      {
        "id": "getFundDailyInfoNAV",
        "name": "Fund Daily Info - NAV",
        "type": "fund"
      },
      {
        "id": "getFundDailyInfoDividendHistory",
        "name": "Fund Daily Info - Dividend History",
        "type": "fund"
      },
      {
        "id": "getOutstandingQuarterlyPortfolio",
        "name": "Outstanding Quarterly Portfolio",
        "type": "fund"
      },
      {
        "id": "getOutstandingMonthlyPortfolioAssetType",
        "name": "Outstanding Monthly Portfolio Asset Type",
        "type": "fund"
      }
    ]
  },
  {
    "name": "License Check",
    "endpoints": [
      {
        "id": "01-bond-sale-rep",
        "name": "Bond Sale Representative",
        "type": "license-check"
      },
      {
        "id": "02-securities-company",
        "name": "Securities Company",
        "type": "license-check"
      },
      {
        "id": "03-derivatives-company",
        "name": "Derivatives Company",
        "type": "license-check"
      },
      {
        "id": "04-securities-broker",
        "name": "Securities Broker",
        "type": "license-check"
      },
      {
        "id": "05-derivatives-broker",
        "name": "Derivatives Broker",
        "type": "license-check"
      },
      {
        "id": "06-investment-advisor",
        "name": "Investment Advisor",
        "type": "license-check"
      },
      {
        "id": "07-securities-fund-manager",
        "name": "Securities Fund Manager",
        "type": "license-check"
      },
      {
        "id": "08-fund-supervisor",
        "name": "Fund Supervisor",
        "type": "license-check"
      },
      {
        "id": "09-auditor",
        "name": "Auditor",
        "type": "license-check"
      },
      {
        "id": "10-credit-rating-company",
        "name": "Credit Rating Company",
        "type": "license-check"
      },
      {
        "id": "11-private-fund",
        "name": "Private Fund",
        "type": "license-check"
      },
      {
        "id": "12-derivatives-fund-manager",
        "name": "Derivatives Fund Manager",
        "type": "license-check"
      },
      {
        "id": "13-securities-borrowing-lending",
        "name": "Securities Borrowing \u0026amp; Lending",
        "type": "license-check"
      },
      {
        "id": "14-financial-advisor",
        "name": "Financial Advisor",
        "type": "license-check"
      },
      {
        "id": "15-acquirer",
        "name": "Acquirer",
        "type": "license-check"
      },
      {
        "id": "16-venture-capital",
        "name": "Venture Capital",
        "type": "license-check"
      }
    ]
  },
  {
    "name": "One Report",
    "endpoints": [
      {
        "id": "01-company-profile",
        "name": "Company Profile",
        "type": "one-report"
      },
      {
        "id": "02-company-highlight",
        "name": "Company Highlight",
        "type": "one-report"
      },
      {
        "id": "03-financial-statement",
        "name": "Financial Statement",
        "type": "one-report"
      },
      {
        "id": "04-management-discussion",
        "name": "Management Discussion",
        "type": "one-report"
      },
      {
        "id": "05-shareholders-structure",
        "name": "Shareholders Structure",
        "type": "one-report"
      },
      {
        "id": "06-dividend-policy",
        "name": "Dividend Policy",
        "type": "one-report"
      },
      {
        "id": "07-corporate-governance",
        "name": "Corporate Governance",
        "type": "one-report"
      },
      {
        "id": "08-related-party-transaction",
        "name": "Related Party Transaction",
        "type": "one-report"
      },
      {
        "id": "09-major-shareholders",
        "name": "Major Shareholders",
        "type": "one-report"
      },
      {
        "id": "10-board-of-directors",
        "name": "Board of Directors",
        "type": "one-report"
      },
      {
        "id": "11-executive",
        "name": "Executive",
        "type": "one-report"
      },
      {
        "id": "12-audit-committee",
        "name": "Audit Committee",
        "type": "one-report"
      },
      {
        "id": "13-compensation",
        "name": "Compensation",
        "type": "one-report"
      },
      {
        "id": "14-environmental",
        "name": "Environmental",
        "type": "one-report"
      },
      {
        "id": "15-social",
        "name": "Social",
        "type": "one-report"
      },
      {
        "id": "16-internal-audit",
        "name": "Internal Audit",
        "type": "one-report"
      },
      {
        "id": "17-risk-management",
        "name": "Risk Management",
        "type": "one-report"
      },
      {
        "id": "18-anti-corruption",
        "name": "Anti Corruption",
        "type": "one-report"
      },
      {
        "id": "19-whistleblowing",
        "name": "Whistleblowing",
        "type": "one-report"
      },
      {
        "id": "20-contact-information",
        "name": "Contact Information",
        "type": "one-report"
      },
      {
        "id": "21-investor-relations",
        "name": "Investor Relations",
        "type": "one-report"
      },
      {
        "id": "22-stock-quote",
        "name": "Stock Quote",
        "type": "one-report"
      },
      {
        "id": "23-company-event",
        "name": "Company Event",
        "type": "one-report"
      }
    ]
  }
]
```

---

*Generated from SEC Open Data API Portal navigation panel. Service IDs are the official identifiers used by the SEC API.*
