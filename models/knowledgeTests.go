package models

type KnowledgeTestResponse struct {
	ID      string                  `json:"id"`
	Version string                  `json:"version"`
	Items   []KnowledgeTestQuestion `json:"items"`
}

type KnowledgeTestQuestion struct {
	Question   string   `json:"question"`
	ChoiceList []string `json:"choice_list"`
	Ans        int      `json:"ans"`
	AnsDetail  string   `json:"ans_detail"`
}