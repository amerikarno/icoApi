package models

type RiskOccupationModel struct {
	OccupationID   string `json:"occupationID" gorm:"column:occupation_id"`
	OccupationCode int    `json:"occupationCode" gorm:"column:occupation_code"`
	OccupationName string `json:"occupationName" gorm:"column:occupation_name"`
	BusinessType   string `json:"businessType" gorm:"column:business_type"`
	RiskGroup      int    `json:"riskGroup" gorm:"column:risk_group"`
}

func (RiskOccupationModel) TableName() string { return "risk_occupations" }
