package models

type SettingTitles struct {
	ID          int    `json:"id,omitempty" `
	TitleNameTh string `json:"titleNameTh,omitempty" gorm:"column:titleNameTh"`
	TitleNameEn string `json:"titleNameEn,omitempty" gorm:"column:titleNameEn"`
}

func (SettingTitles) TableName() string {
	return "setting_titles"
}