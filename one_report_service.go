package sec

import (
	"context"
	"encoding/json"
	"fmt"
)

const oneReportBaseURL = "https://api.sec.or.th/onereport"

// --- SBO (Small Business Overview) ---

// SBOInfo represents general company information from the One Report.
type SBOInfo struct {
	LastUpdDate          DateTime `json:"last_upd_date"`
	ReportYear           string   `json:"report_year"`
	UniqueID             string   `json:"unique_id"`
	Language             string   `json:"language"`
	CorpName             string   `json:"corp_name"`
	Symbol               string   `json:"symbol"`
	Address              string   `json:"address"`
	Province             string   `json:"province"`
	ZipCode              string   `json:"zip_code"`
	BusinessType         string   `json:"business_type"`
	RegisteredNumber     string   `json:"registered_number"`
	Telephone            string   `json:"telephone"`
	Fax                  string   `json:"fax"`
	Website              string   `json:"website"`
	Email                string   `json:"email"`
	CommonPaidupShare    float64  `json:"common_paidup_share"`
	PreferredPaidupShare float64  `json:"preferred_paidup_share"`
}

// SBORD represents R&D information.
type SBORD struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SBOProductIncome represents product income structure.
type SBOProductIncome struct {
	LastUpdDate                     DateTime `json:"last_upd_date"`
	ReportYear                      string   `json:"report_year"`
	UniqueID                        string   `json:"unique_id"`
	BusinessIncomeCode              string   `json:"business_income_code"`
	Sequence                        int      `json:"sequence"`
	BusinessIncomeDesc              string   `json:"business_income_desc"`
	AsofYear                        float64  `json:"asof_year"`
	AsofYesteryear                  float64  `json:"asof_yesteryear"`
	AsofYearBeforeYesteryear        float64  `json:"asof_year_before_yesteryear"`
	AsofYearPercent                 float64  `json:"asof_year_percent"`
	AsofYesteryearPercent           float64  `json:"asof_yesteryear_percent"`
	AsofYearBeforeYesteryearPercent float64  `json:"asof_year_before_yesteryear_percent"`
}

// SBOExportIncome represents export income structure.
type SBOExportIncome struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SBORisk represents risk management details.
type SBORisk struct {
	LastUpdDate  DateTime `json:"last_upd_date"`
	ReportYear   string   `json:"report_year"`
	UniqueID     string   `json:"unique_id"`
	RiskCategory string   `json:"risk_category"`
	RiskCode     string   `json:"risk_code"`
	Sequence     int      `json:"sequence"`
	Choice       string   `json:"choice"`
	HolderRisk   string   `json:"holder_risk"`
	ForeignRisk  string   `json:"foreign_risk"`
}

// GetSBOInfo returns general company information.
// Language: "T" for Thai, "E" for English.
func (c *Client) GetSBOInfo(ctx context.Context, reportYear, language string) ([]SBOInfo, error) {
	url := fmt.Sprintf("%s/sbo/%s/info/%s", oneReportBaseURL, reportYear, language)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sbo info: %w", err)
	}
	var items []SBOInfo
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sbo info: %w", err)
	}
	return items, nil
}

// GetSBORD returns R&D information for a company.
func (c *Client) GetSBORD(ctx context.Context, reportYear, uniqueID string) ([]SBORD, error) {
	url := fmt.Sprintf("%s/sbo/%s/rd/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sbo rd: %w", err)
	}
	var items []SBORD
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sbo rd: %w", err)
	}
	return items, nil
}

// GetSBOProductIncome returns product income structure.
func (c *Client) GetSBOProductIncome(ctx context.Context, reportYear, uniqueID string) ([]SBOProductIncome, error) {
	url := fmt.Sprintf("%s/sbo/%s/product_income/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sbo product income: %w", err)
	}
	var items []SBOProductIncome
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sbo product income: %w", err)
	}
	return items, nil
}

// GetSBOExportIncome returns export income structure.
func (c *Client) GetSBOExportIncome(ctx context.Context, reportYear, uniqueID string) ([]SBOExportIncome, error) {
	url := fmt.Sprintf("%s/sbo/%s/export_income/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sbo export income: %w", err)
	}
	var items []SBOExportIncome
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sbo export income: %w", err)
	}
	return items, nil
}

// GetSBORisk returns risk management details.
func (c *Client) GetSBORisk(ctx context.Context, reportYear, uniqueID string) ([]SBORisk, error) {
	url := fmt.Sprintf("%s/sbo/%s/risk/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sbo risk: %w", err)
	}
	var items []SBORisk
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sbo risk: %w", err)
	}
	return items, nil
}

// --- Sustainability ---

// SustainabilityDetail represents sustainability policy and targets.
type SustainabilityDetail struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SustainabilityEnvironmentIssue represents environmental issues.
type SustainabilityEnvironmentIssue struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SustainabilityHumanRightsIssue represents human rights issues.
type SustainabilityHumanRightsIssue struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// GetSustainabilityDetail returns sustainability policy and targets.
func (c *Client) GetSustainabilityDetail(ctx context.Context, reportYear, uniqueID string) ([]SustainabilityDetail, error) {
	url := fmt.Sprintf("%s/sustainability/%s/detail/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sustainability detail: %w", err)
	}
	var items []SustainabilityDetail
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sustainability detail: %w", err)
	}
	return items, nil
}

// GetSustainabilityEnvironmentIssue returns environmental issues.
func (c *Client) GetSustainabilityEnvironmentIssue(ctx context.Context, reportYear, uniqueID string) ([]SustainabilityEnvironmentIssue, error) {
	url := fmt.Sprintf("%s/sustainability/%s/environment_issue/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sustainability environment issue: %w", err)
	}
	var items []SustainabilityEnvironmentIssue
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sustainability environment issue: %w", err)
	}
	return items, nil
}

// GetSustainabilityHumanRightsIssue returns human rights issues.
func (c *Client) GetSustainabilityHumanRightsIssue(ctx context.Context, reportYear, uniqueID string) ([]SustainabilityHumanRightsIssue, error) {
	url := fmt.Sprintf("%s/sustainability/%s/humanrights_issue/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get sustainability humanrights issue: %w", err)
	}
	var items []SustainabilityHumanRightsIssue
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal sustainability humanrights issue: %w", err)
	}
	return items, nil
}

// --- Financial Statement ---

// FinancialStatement represents financial statements and ratios.
type FinancialStatement struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// GetFinancialStatement returns financial statements and ratios.
func (c *Client) GetFinancialStatement(ctx context.Context, reportYear, uniqueID string) ([]FinancialStatement, error) {
	url := fmt.Sprintf("%s/fs/%s/financial_statement/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get financial statement: %w", err)
	}
	var items []FinancialStatement
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal financial statement: %w", err)
	}
	return items, nil
}

// --- CGP (Corporate Governance Policy) ---

// CGPGovernance represents corporate governance policy.
type CGPGovernance struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGPDirector represents director policies and practices.
type CGPDirector struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGPCodeOfConduct represents business code of conduct.
type CGPCodeOfConduct struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// GetCGPGovernance returns corporate governance policy.
func (c *Client) GetCGPGovernance(ctx context.Context, reportYear, uniqueID string) ([]CGPGovernance, error) {
	url := fmt.Sprintf("%s/cgp/%s/governance/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgp governance: %w", err)
	}
	var items []CGPGovernance
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgp governance: %w", err)
	}
	return items, nil
}

// GetCGPDirector returns director policies and practices.
func (c *Client) GetCGPDirector(ctx context.Context, reportYear, uniqueID string) ([]CGPDirector, error) {
	url := fmt.Sprintf("%s/cgp/%s/director/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgp director: %w", err)
	}
	var items []CGPDirector
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgp director: %w", err)
	}
	return items, nil
}

// GetCGPCodeOfConduct returns business code of conduct.
func (c *Client) GetCGPCodeOfConduct(ctx context.Context, reportYear, uniqueID string) ([]CGPCodeOfConduct, error) {
	url := fmt.Sprintf("%s/cgp/%s/code_of_conduct/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgp code of conduct: %w", err)
	}
	var items []CGPCodeOfConduct
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgp code of conduct: %w", err)
	}
	return items, nil
}

// --- CGS (Corporate Governance Structure) ---

// CGSBoard represents board composition.
type CGSBoard struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSAuditorCompany represents auditor company information.
type CGSAuditorCompany struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSDirectorPerformance represents board meeting attendance.
type CGSDirectorPerformance struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSEmployee represents employee information.
type CGSEmployee struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSBods represents board of directors information.
type CGSBods struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSExecutives represents executive information.
type CGSExecutives struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// CGSCommitteesOthers represents other committees information.
type CGSCommitteesOthers struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// GetCGSBoard returns board composition.
func (c *Client) GetCGSBoard(ctx context.Context, reportYear, uniqueID string) ([]CGSBoard, error) {
	url := fmt.Sprintf("%s/cgs/%s/board/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs board: %w", err)
	}
	var items []CGSBoard
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs board: %w", err)
	}
	return items, nil
}

// GetCGSAuditorCompany returns auditor company information.
func (c *Client) GetCGSAuditorCompany(ctx context.Context, reportYear, uniqueID string) ([]CGSAuditorCompany, error) {
	url := fmt.Sprintf("%s/cgs/%s/auditor_company/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs auditor company: %w", err)
	}
	var items []CGSAuditorCompany
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs auditor company: %w", err)
	}
	return items, nil
}

// GetCGSDirectorPerformance returns board meeting attendance.
func (c *Client) GetCGSDirectorPerformance(ctx context.Context, reportYear, uniqueID string) ([]CGSDirectorPerformance, error) {
	url := fmt.Sprintf("%s/cgs/%s/director_performance/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs director performance: %w", err)
	}
	var items []CGSDirectorPerformance
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs director performance: %w", err)
	}
	return items, nil
}

// GetCGSEmployee returns employee information.
func (c *Client) GetCGSEmployee(ctx context.Context, reportYear, uniqueID string) ([]CGSEmployee, error) {
	url := fmt.Sprintf("%s/cgs/%s/employee/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs employee: %w", err)
	}
	var items []CGSEmployee
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs employee: %w", err)
	}
	return items, nil
}

// GetCGSBods returns board of directors information.
func (c *Client) GetCGSBods(ctx context.Context, reportYear, uniqueID string) ([]CGSBods, error) {
	url := fmt.Sprintf("%s/cgs/%s/bods/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs bods: %w", err)
	}
	var items []CGSBods
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs bods: %w", err)
	}
	return items, nil
}

// GetCGSExecutives returns executive information.
func (c *Client) GetCGSExecutives(ctx context.Context, reportYear, uniqueID string) ([]CGSExecutives, error) {
	url := fmt.Sprintf("%s/cgs/%s/executives/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs executives: %w", err)
	}
	var items []CGSExecutives
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs executives: %w", err)
	}
	return items, nil
}

// GetCGSCommitteesOthers returns other committees information.
func (c *Client) GetCGSCommitteesOthers(ctx context.Context, reportYear, uniqueID string) ([]CGSCommitteesOthers, error) {
	url := fmt.Sprintf("%s/cgs/%s/committees/%s/others", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get cgs committees others: %w", err)
	}
	var items []CGSCommitteesOthers
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal cgs committees others: %w", err)
	}
	return items, nil
}

// --- SCP (Social Performance) ---

// SCPLaborDispute represents labor dispute information.
type SCPLaborDispute struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SCPCSRActivity represents CSR activities.
type SCPCSRActivity struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SCPEmployeeInfo represents employee and compensation information.
type SCPEmployeeInfo struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// SCPEmployeeDevelopment represents training and safety information.
type SCPEmployeeDevelopment struct {
	LastUpdDate DateTime `json:"last_upd_date"`
	ReportYear  string   `json:"report_year"`
	UniqueID    string   `json:"unique_id"`
}

// GetSCPLaborDispute returns labor dispute information.
func (c *Client) GetSCPLaborDispute(ctx context.Context, reportYear, uniqueID string) ([]SCPLaborDispute, error) {
	url := fmt.Sprintf("%s/scp/%s/labor_dispute/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get scp labor dispute: %w", err)
	}
	var items []SCPLaborDispute
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal scp labor dispute: %w", err)
	}
	return items, nil
}

// GetSCPCSRActivity returns CSR activities.
func (c *Client) GetSCPCSRActivity(ctx context.Context, reportYear, uniqueID string) ([]SCPCSRActivity, error) {
	url := fmt.Sprintf("%s/scp/%s/csr_activity/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get scp csr activity: %w", err)
	}
	var items []SCPCSRActivity
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal scp csr activity: %w", err)
	}
	return items, nil
}

// GetSCPEmployeeInfo returns employee and compensation information.
func (c *Client) GetSCPEmployeeInfo(ctx context.Context, reportYear, uniqueID string) ([]SCPEmployeeInfo, error) {
	url := fmt.Sprintf("%s/scp/%s/employee_info/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get scp employee info: %w", err)
	}
	var items []SCPEmployeeInfo
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal scp employee info: %w", err)
	}
	return items, nil
}

// GetSCPEmployeeDevelopment returns training and safety information.
func (c *Client) GetSCPEmployeeDevelopment(ctx context.Context, reportYear, uniqueID string) ([]SCPEmployeeDevelopment, error) {
	url := fmt.Sprintf("%s/scp/%s/employee_development/%s", oneReportBaseURL, reportYear, uniqueID)
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get scp employee development: %w", err)
	}
	var items []SCPEmployeeDevelopment
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal scp employee development: %w", err)
	}
	return items, nil
}
