package models

import "time"

type CustomerExamsRequest struct {
	AccountID           string    `json:"id" gorm:"column:customer_id"`
	SuiteTestResult     string    `json:"suiteTestResult" gorm:"column:suite_test_result"`
	IsFatca             bool      `json:"isFatca" gorm:"column:is_fatca"`
	FatcaInfo           string    `json:"fatcaInfo" gorm:"column:fatca_info"`
	IsKnowledgeDone     bool      `json:"isKnowledgeDone" gorm:"column:is_knowledge_done"`
	KnowledgeTestResult string    `json:"knowledgeTestResult" gorm:"column:knowledge_test_result"`
	CreateAt            time.Time `json:"createAt" gorm:"column:create_at"`
}

func (CustomerExamsRequest) TableName() string { return "customer_exams" }
