package models

type RiskCountryModel struct {
	RiskCountryID string `json:"riskCountryID" gorm:"risk_countriy_id"`
	CountryName   string `json:"countryName" gorm:"country_name"`
	GroupName     string `json:"groupName" gorm:"group_name"`
	RiskLevel     int    `json:"risk_level" gorm:"risk_level"`
}

func (RiskCountryModel) TableName() string { return "risk_countries" }
