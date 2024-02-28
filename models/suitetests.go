package models

type SuiteTestQuestions struct {
	Question   string   `json:"question"`
	ChoiceList []string `json:"choice_list"`
	Type       int      `json:"type"`
}

type SuiteTestResponse struct {
	ID      string               `json:"id"`
	Version string               `json:"version"`
	Items   []SuiteTestQuestions `json:"items"`
}

type SuiteSelect struct {
	ID   string `json:"id"`
	Risk string `json:"risk"`
}

type SuiteSelectLists struct {
	InvestmentRisk []SuiteSelect `json:"investment_risk"`
}