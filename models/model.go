package models

import "time"

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

type CustomerInformations struct {
	AccountID         string    `json:"id" gorm:"column:id"`
	ThTitle           string    `json:"thTitle" gorm:"column:th_title"`
	ThName            string    `json:"thName" gorm:"column:th_name"`
	ThSurname         string    `json:"thSurname" gorm:"column:th_surname"`
	EngTitle          string    `json:"engTitle" gorm:"column:en_title"`
	EngName           string    `json:"engName" gorm:"column:en_name"`
	EngSurname        string    `json:"engSurname" gorm:"column:en_surname"`
	Email             string    `json:"email" gorm:"column:email"`
	Mobile            string    `json:"mobile" gorm:"column:mobile_no"`
	Agreement         bool      `json:"agreement" gorm:"column:personal_agreement"`
	BirthDate         string    `json:"birthDate" gorm:"column:birth_date"`
	MarriageStatus    string    `json:"marriageStatus" gorm:"column:marriage_status"`
	IDCard            string    `json:"idCard" gorm:"column:id_card"`
	LaserCode         string    `json:"laserCode" gorm:"column:laser_code"`
	SourceOfIncome    string    `json:"sourceOfIncome" gorm:"column:source_of_income"`
	CurrentOccupation string    `json:"currentOccupation" gorm:"column:current_occupation"`
	OfficeName        string    `json:"officeName" gorm:"column:office_name"`
	TypeOfBusiness    string    `json:"typeOfBusiness" gorm:"column:type_of_business"`
	PositionName      string    `json:"positionName" gorm:"column:position_name"`
	SalaryRange       string    `json:"salaryRange" gorm:"column:salary_range"`
	Pages             bool      `json:"pages" gorm:"column:personal_pages"`
	Update            time.Time `json:"update" gorm:"column:update_at"`
	Create            time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerInformations) TableName() string { return "customer_informations" }

type CustomerAddressRequest struct {
	AccountID           string    `json:"id" gorm:"column:customer_id"`
	HomeNumber          string    `json:"homeNumber" gorm:"column:home_number"`
	VillageNumber       string    `json:"villageNumber" gorm:"column:village_number"`
	VillageName         string    `json:"villageName" gorm:"column:village_name"`
	SubStreetName       string    `json:"subStreetName" gorm:"column:sub_street_name"`
	StreetName          string    `json:"streetName" gorm:"column:street_name"`
	SubDistrictName     string    `json:"subDistrictName" gorm:"column:sub_district_name"`
	DistrictName        string    `json:"districtName" gorm:"column:district_name"`
	ProvinceName        string    `json:"provinceName" gorm:"column:province_name"`
	ZipCode             string    `json:"zipCode" gorm:"column:zipcode"`
	CountryName         string    `json:"countryName" gorm:"column:country_name"`
	TypeOfAddress       string    `json:"typeOfAddress"` //gorm:"column:type_of_address"`
	IsRegisteredAddress bool      `json:"isRegisteredAddress" gorm:"column:is_registered_address"`
	IsCurrentAddress    bool      `json:"isCurrentAddress" gorm:"column:is_current_address"`
	IsOfficeAddress     bool      `json:"isOfficeAddress" gorm:"column:is_office_address"`
	Create              time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerAddressRequest) TableName() string { return "customer_addresses" }

type CustomerAddressResponse struct {
	AccountID       string `json:"id" gorm:"column:customer_id"`
	HomeNumber      string `json:"homeNumber" gorm:"column:home_number"`
	VillageNumber   string `json:"villageNumber" gorm:"column:village_number"`
	VillageName     string `json:"villageName" gorm:"column:village_name"`
	SubStreetName   string `json:"subStreetName" gorm:"column:sub_street_name"`
	StreetName      string `json:"streetName" gorm:"column:street_name"`
	SubDistrictName string `json:"subDistrictName" gorm:"column:sub_district_name"`
	DistrictName    string `json:"districtName" gorm:"column:district_name"`
	ProvinceName    string `json:"provinceName" gorm:"column:province_name"`
	ZipCode         string `json:"zipCode" gorm:"column:zipcode"`
	CountryName     string `json:"countryName" gorm:"column:country_name"`
	// TypeOfAddress   string    `json:"typeOfAddress"` //gorm:"column:type_of_address"`
	IsRegisteredAddress bool      `json:"isRegisteredAddress" gorm:"column:is_registered_address"`
	IsCurrentAddress    bool      `json:"isCurrentAddress" gorm:"column:is_current_address"`
	IsOfficeAddress     bool      `json:"isOfficeAddress" gorm:"column:is_office_address"`
	Create              time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerAddressResponse) TableName() string { return "customer_addresses" }

type CustomerAddresses struct {
	AccountID           string    `json:"id" gorm:"column:customer_id"`
	HomeNumber          string    `json:"homeNumber" gorm:"column:home_number"`
	VillageNumber       string    `json:"villageNumber" gorm:"column:village_number"`
	VillageName         string    `json:"villageName" gorm:"column:village_name"`
	SubStreetName       string    `json:"subStreetName" gorm:"column:sub_street_name"`
	StreetName          string    `json:"streetName" gorm:"column:street_name"`
	SubDistrictName     string    `json:"subDistrictName" gorm:"column:sub_district_name"`
	DistrictName        string    `json:"districtName" gorm:"column:district_name"`
	ProvinceName        string    `json:"provinceName" gorm:"column:province_name"`
	ZipCode             string    `json:"zipCode" gorm:"column:zipcode"`
	CountryName         string    `json:"countryName" gorm:"column:country_name"`
	IsRegisteredAddress bool      `json:"isRegisteredAddress" gorm:"column:is_registered_address"`
	IsCurrentAddress    bool      `json:"isCurrentAddress" gorm:"column:is_current_address"`
	IsOfficeAddress     bool      `json:"isOfficeAddress" gorm:"column:is_office_address"`
	SourceOfIncome      string    `json:"sourceOfIncome" gorm:"column:source_of_income"`
	CurrentOccupation   string    `json:"currentOccupation" gorm:"column:current_occupation"`
	OfficeName          string    `json:"officeName" gorm:"column:office_name"`
	TypeOfBusiness      string    `json:"typeOfBusiness" gorm:"column:type_of_business"`
	PositionName        string    `json:"positionName" gorm:"column:position_name"`
	SalaryRange         string    `json:"salaryRange" gorm:"column:salary_range"`
	Create              time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerAddresses) TableName() string { return "customer_addresses" }

type CustomerBookbankRequest struct {
	AccountID         string    `json:"id" gorm:"column:customer_id"`
	BankName          string    `json:"bankName" gorm:"column:bank_name"`
	BankBranchName    string    `json:"bankBranchName" gorm:"column:bank_branch_name"`
	BankAccountNumber string    `json:"bankAccountNumber" gorm:"column:bank_account_number"`
	IsDefalut         bool      `json:"isDefalut" gorm:"column:is_defalut"`
	AccountType       string    `json:"accountType" gorm:"column:account_type"`
	Create            time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerBookbankRequest) TableName() string { return "customer_bookbanks" }

type CustomerBookbankResponse struct {
	AccountID         string `json:"id" gorm:"column:customer_id"`
	BankName          string `json:"bankName" gorm:"column:bank_name"`
	BankBranchName    string `json:"bankBranchName" gorm:"column:bank_branch_name"`
	BankAccountNumber string `json:"bankAccountNumber" gorm:"column:bank_account_number"`
	IsDefalut         bool   `json:"isDefalut" gorm:"column:is_default"`
	// AccountType       string    `json:"accountType" gorm:"column:account_type"`
	IsDeposit  bool      `json:"isDeposit" gorm:"column:is_deposit"`
	IsWithdraw bool      `json:"isWithdraw" gorm:"column:is_withdraw"`
	Create     time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerBookbankResponse) TableName() string { return "customer_bookbanks" }

type CustomerBookbanks struct {
	AccountID         string    `json:"id" gorm:"column:customer_id"`
	BankName          string    `json:"bankName" gorm:"column:bank_name"`
	BankBranchName    string    `json:"bankBranchName" gorm:"column:bank_branch_name"`
	BankAccountNumber string    `json:"bankAccountNumber" gorm:"column:bank_account_number"`
	IsDefalut         bool      `json:"isDefalut" gorm:"column:is_defalut"`
	IsDeposit         bool      `json:"isDeposit" gorm:"column:is_deposit"`
	IsWithdraw        bool      `json:"isWithdraw" gorm:"column:is_withdraw"`
	Create            time.Time `json:"create" gorm:"column:create_at"`
}

func (CustomerBookbanks) TableName() string { return "customer_bookbanks" }

type PersonalInformations struct {
	CustomerInformation   CustomerInformations      `json:"customerInformation"`
	CustomerAddresseLists []CustomerAddressResponse `json:"customerAddresseLists"`
	CustomerBookbankLists []CustomerBookbankResponse `json:"customerBookbankLits"`
}

type PersonalInformationPostRequests struct {
	CID               string                    `json:"cid"`
	RegisteredAddress CustomerAddressRequest    `json:"registeredAddress"`
	CurrentAddress    CustomerAddressRequest    `json:"currentAddress"`
	OfficeAddress     CustomerAddressRequest    `json:"officeAddress"`
	Occupation        CustomerOccupationRequest `json:"occupation"`
	FirstBankAccount  CustomerBookbankRequest   `json:"firstBankAccount"`
	SecondBankAccount CustomerBookbankRequest   `json:"secondBankAccount"`
}

type CustomerOccupationRequest struct {
	AccountID         string `json:"id" gorm:"column:customer_id"`
	SourceOfIncome    string `json:"sourceOfIncome" gorm:"column:source_of_income"`
	CurrentOccupation string `json:"currentOccupation" gorm:"column:current_occupation"`
	OfficeName        string `json:"officeName" gorm:"column:office_name"`
	TypeOfBusiness    string `json:"typeOfBusiness" gorm:"column:type_of_business"`
	PositionName      string `json:"positionName" gorm:"column:position_name"`
	SalaryRange       string `json:"salaryRange" gorm:"column:salary_range"`
}
