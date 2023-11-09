package models

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
	Id int `json:"id"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
	GeographyId int `json:"geography_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Amphure []AmphureList `json:"amphure"`
}

type AmphureList struct {
	Id int `json:"id"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
	ProvinceId int `json:"province_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Tambon []TambonList `json:"tambon"`
}

type TambonList struct {
	Id int `json:"id"`
	ZipCode int `json:"zip_code"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
	AmphureId int `json:"amphure_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type PostIDcard struct {
	BirthDate string `json:"birthDate"`
	MarriageStatus string `json:"marriageStatus"`
	IDCard string `json:"idCard"`
	LaserCode string `json:"laserCode"`
}