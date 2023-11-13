package models

type VerifyEmailMobileResponse struct {
	IsRegistedEmail         bool   `json:"isRegisteredEmail"`
	IsInvalidEmailFormat    bool   `json:"isInvalidEmailFormat"`
	RegistedEmail           string `json:"registeredEmail"`
	IsRegistedMobileno      bool   `json:"isRegisteredMobileNo"`
	IsInvalidMobileNoFormat bool   `json:"isInvalidMobileNoFormat"`
	RegistedMobileNo        string `json:"registeredMobileNo"`
	RegistedPage            string `json:"registeredPage"`
}
type VerifyEmailResponse struct {
	IsRegistedEmail      bool   `json:"isRegisteredEmail"`
	IsInvalidEmailFormat bool   `json:"isInvalidEmailFormat"`
	RegistedEmail        string `json:"registeredEmail"`
	RegistedPage         string `json:"registeredPage"`
}
type VerifyMobileNoResponse struct {
	IsRegistedMobileno      bool   `json:"isRegisteredMobileNo"`
	IsInvalidMobileNoFormat bool   `json:"isInvalidMobileNoFormat"`
	RegistedMobileNo        string `json:"registeredMobileNo"`
	RegistedPage            string `json:"registeredPage"`
}
type ProvinceList struct {
	Id          int           `json:"id"`
	NameTh      string        `json:"name_th"`
	NameEn      string        `json:"name_en"`
	GeographyId int           `json:"geography_id"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	DeletedAt   string        `json:"deleted_at"`
	Amphure     []AmphureList `json:"amphure"`
}

type AmphureList struct {
	Id         int          `json:"id"`
	NameTh     string       `json:"name_th"`
	NameEn     string       `json:"name_en"`
	ProvinceId int          `json:"province_id"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
	DeletedAt  string       `json:"deleted_at"`
	Tambon     []TambonList `json:"tambon"`
}

type TambonList struct {
	Id        int    `json:"id"`
	ZipCode   int    `json:"zip_code"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	AmphureId int    `json:"amphure_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type IDCardOpenAccounts struct {
	AccountID      string `json:"id" gorm:"column:id"`
	BirthDate      string `json:"birthDate" gorm:"column:birth_date"`
	MarriageStatus string `json:"marriageStatus" gorm:"column:marriage_status"`
	IDCard         string `json:"idCard" gorm:"column:id_card"`
	LaserCode      string `json:"laserCode" gorm:"column:laser_code"`
}

func (IDCardOpenAccounts) TableName() string { return "personal_informations" }
