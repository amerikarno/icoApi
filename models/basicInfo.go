package models

type InvestmentTypeST struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BusinessTypeST struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type InvestmentObjectivesTypeST struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type EducationST struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CurrentCareerTypeST struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	BusinessFlag  string `json:"business_flag,omitempty"`
	WorkplaceFlag string `json:"workplace_flag,omitempty"`
	PositionFlag  string `json:"position_flag,omitempty"`
}

type MonthlyIncomeTypeST struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	ID0   int    `json:" id ,omitempty"`
	Name0 string `json:" name,omitempty"`
}

type BankST struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}
type BankBranchST struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}
type CountryST struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type BasicInfo struct {
	InvestmentType           []InvestmentTypeST           `json:"investment_type,omitempty"`
	BusinessType             []BusinessTypeST             `json:"business_type,omitempty"`
	CurrentCareerType        []CurrentCareerTypeST        `json:"current_career_type,omitempty"`
	MonthlyIncomeType        []MonthlyIncomeTypeST        `json:"monthly_income_type,omitempty"`
	InvestmentObjectivesType []InvestmentObjectivesTypeST `json:"investment_objectives_type,omitempty"`
	Bank                     []BankST                     `json:"bank,omitempty"`
	BankBranch               []BankBranchST               `json:"bank_branch,omitempty"`
	Country                  []CountryST                  `json:"country,omitempty"`
	Education                []EducationST                `json:"education,omitempty"`
	Province                 any                          `json:"province,omitempty"`
	Title                    any                          `json:"title,omitempty"`
	Success                  bool                         `json:"success,omitempty"`
	Error                    any                          `json:"error,omitempty"`
	Message                  any                          `json:"message,omitempty"`
	NextModule               any                          `json:"next_module,omitempty"`
}
