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
	ThTitle        string `json:"thTitle" gorm:"column:th_title"`
	ThName         string `json:"thName" gorm:"column:th_name"`
	ThSurname      string `json:"thSurname" gorm:"column:th_surname"`
	EngTitle       string `json:"engTitle" gorm:"column:en_title"`
	EngName        string `json:"engName" gorm:"column:en_name"`
	EngSurname     string `json:"engSurname" gorm:"column:en_surname"`
	Email          string `json:"email" gorm:"column:email"`
	Mobile         string `json:"mobile" gorm:"column:mobile_no"`
	Agreement      string `json:"agreement" gorm:"column:personal_agreement"`
	BirthDate      string `json:"birthDate" gorm:"column:birth_date"`
	MarriageStatus string `json:"marriageStatus" gorm:"column:marriage_status"`
	IDCard         string `json:"idCard" gorm:"column:id_card"`
	LaserCode      string `json:"laserCode" gorm:"column:laser_code"`
	Pages          string `json:"pages" gorm:"column:personal_pages"`
	Create         string `json:"create" gorm:"column:create_at"`
}

func (IDCardOpenAccounts) TableName() string { return "personal_informations" }
