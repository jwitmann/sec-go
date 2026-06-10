package sec

import (
	"context"
	"encoding/json"
	"fmt"
)

const oneReportBaseURL = "https://api.sec.or.th/onereport"

func fetchOneReport[T any](ctx context.Context, c *Client, path string) ([]T, error) {
	url := oneReportBaseURL + path
	data, err := c.getAbsolute(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get %s: %w", path, err)
	}
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal %s: %w", path, err)
	}
	return items, nil
}

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
	return fetchOneReport[SBOInfo](ctx, c, fmt.Sprintf("/sbo/%s/info/%s", reportYear, language))
}

// GetSBORD returns R&D information for a company.
func (c *Client) GetSBORD(ctx context.Context, reportYear, uniqueID string) ([]SBORD, error) {
	return fetchOneReport[SBORD](ctx, c, fmt.Sprintf("/sbo/%s/rd/%s", reportYear, uniqueID))
}

// GetSBOProductIncome returns product income structure.
func (c *Client) GetSBOProductIncome(ctx context.Context, reportYear, uniqueID string) ([]SBOProductIncome, error) {
	return fetchOneReport[SBOProductIncome](ctx, c, fmt.Sprintf("/sbo/%s/product_income/%s", reportYear, uniqueID))
}

// GetSBOExportIncome returns export income structure.
func (c *Client) GetSBOExportIncome(ctx context.Context, reportYear, uniqueID string) ([]SBOExportIncome, error) {
	return fetchOneReport[SBOExportIncome](ctx, c, fmt.Sprintf("/sbo/%s/export_income/%s", reportYear, uniqueID))
}

// GetSBORisk returns risk management details.
func (c *Client) GetSBORisk(ctx context.Context, reportYear, uniqueID string) ([]SBORisk, error) {
	return fetchOneReport[SBORisk](ctx, c, fmt.Sprintf("/sbo/%s/risk/%s", reportYear, uniqueID))
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
	return fetchOneReport[SustainabilityDetail](ctx, c, fmt.Sprintf("/sustainability/%s/detail/%s", reportYear, uniqueID))
}

// GetSustainabilityEnvironmentIssue returns environmental issues.
func (c *Client) GetSustainabilityEnvironmentIssue(ctx context.Context, reportYear, uniqueID string) ([]SustainabilityEnvironmentIssue, error) {
	return fetchOneReport[SustainabilityEnvironmentIssue](ctx, c, fmt.Sprintf("/sustainability/%s/environment_issue/%s", reportYear, uniqueID))
}

// GetSustainabilityHumanRightsIssue returns human rights issues.
func (c *Client) GetSustainabilityHumanRightsIssue(ctx context.Context, reportYear, uniqueID string) ([]SustainabilityHumanRightsIssue, error) {
	return fetchOneReport[SustainabilityHumanRightsIssue](ctx, c, fmt.Sprintf("/sustainability/%s/humanrights_issue/%s", reportYear, uniqueID))
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
	return fetchOneReport[FinancialStatement](ctx, c, fmt.Sprintf("/fs/%s/financial_statement/%s", reportYear, uniqueID))
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
	return fetchOneReport[CGPGovernance](ctx, c, fmt.Sprintf("/cgp/%s/governance/%s", reportYear, uniqueID))
}

// GetCGPDirector returns director policies and practices.
func (c *Client) GetCGPDirector(ctx context.Context, reportYear, uniqueID string) ([]CGPDirector, error) {
	return fetchOneReport[CGPDirector](ctx, c, fmt.Sprintf("/cgp/%s/director/%s", reportYear, uniqueID))
}

// GetCGPCodeOfConduct returns business code of conduct.
func (c *Client) GetCGPCodeOfConduct(ctx context.Context, reportYear, uniqueID string) ([]CGPCodeOfConduct, error) {
	return fetchOneReport[CGPCodeOfConduct](ctx, c, fmt.Sprintf("/cgp/%s/code_of_conduct/%s", reportYear, uniqueID))
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
	return fetchOneReport[CGSBoard](ctx, c, fmt.Sprintf("/cgs/%s/board/%s", reportYear, uniqueID))
}

// GetCGSAuditorCompany returns auditor company information.
func (c *Client) GetCGSAuditorCompany(ctx context.Context, reportYear, uniqueID string) ([]CGSAuditorCompany, error) {
	return fetchOneReport[CGSAuditorCompany](ctx, c, fmt.Sprintf("/cgs/%s/auditor_company/%s", reportYear, uniqueID))
}

// GetCGSDirectorPerformance returns board meeting attendance.
func (c *Client) GetCGSDirectorPerformance(ctx context.Context, reportYear, uniqueID string) ([]CGSDirectorPerformance, error) {
	return fetchOneReport[CGSDirectorPerformance](ctx, c, fmt.Sprintf("/cgs/%s/director_performance/%s", reportYear, uniqueID))
}

// GetCGSEmployee returns employee information.
func (c *Client) GetCGSEmployee(ctx context.Context, reportYear, uniqueID string) ([]CGSEmployee, error) {
	return fetchOneReport[CGSEmployee](ctx, c, fmt.Sprintf("/cgs/%s/employee/%s", reportYear, uniqueID))
}

// GetCGSBods returns board of directors information.
func (c *Client) GetCGSBods(ctx context.Context, reportYear, uniqueID string) ([]CGSBods, error) {
	return fetchOneReport[CGSBods](ctx, c, fmt.Sprintf("/cgs/%s/bods/%s", reportYear, uniqueID))
}

// GetCGSExecutives returns executive information.
func (c *Client) GetCGSExecutives(ctx context.Context, reportYear, uniqueID string) ([]CGSExecutives, error) {
	return fetchOneReport[CGSExecutives](ctx, c, fmt.Sprintf("/cgs/%s/executives/%s", reportYear, uniqueID))
}

// GetCGSCommitteesOthers returns other committees information.
func (c *Client) GetCGSCommitteesOthers(ctx context.Context, reportYear, uniqueID string) ([]CGSCommitteesOthers, error) {
	return fetchOneReport[CGSCommitteesOthers](ctx, c, fmt.Sprintf("/cgs/%s/committees/%s/others", reportYear, uniqueID))
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
	return fetchOneReport[SCPLaborDispute](ctx, c, fmt.Sprintf("/scp/%s/labor_dispute/%s", reportYear, uniqueID))
}

// GetSCPCSRActivity returns CSR activities.
func (c *Client) GetSCPCSRActivity(ctx context.Context, reportYear, uniqueID string) ([]SCPCSRActivity, error) {
	return fetchOneReport[SCPCSRActivity](ctx, c, fmt.Sprintf("/scp/%s/csr_activity/%s", reportYear, uniqueID))
}

// GetSCPEmployeeInfo returns employee and compensation information.
func (c *Client) GetSCPEmployeeInfo(ctx context.Context, reportYear, uniqueID string) ([]SCPEmployeeInfo, error) {
	return fetchOneReport[SCPEmployeeInfo](ctx, c, fmt.Sprintf("/scp/%s/employee_info/%s", reportYear, uniqueID))
}

// GetSCPEmployeeDevelopment returns training and safety information.
func (c *Client) GetSCPEmployeeDevelopment(ctx context.Context, reportYear, uniqueID string) ([]SCPEmployeeDevelopment, error) {
	return fetchOneReport[SCPEmployeeDevelopment](ctx, c, fmt.Sprintf("/scp/%s/employee_development/%s", reportYear, uniqueID))
}
