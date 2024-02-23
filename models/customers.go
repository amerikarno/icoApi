package models

type Customer struct {
	CustomerInformations
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
	BankName          string    `json:"bankName" gorm:"column:bank_name"`
	BankBranchName    string    `json:"bankBranchName" gorm:"column:bank_branch_name"`
	BankAccountNumber string    `json:"bankAccountNumber" gorm:"column:bank_account_number"`
	IsDefalut         bool      `json:"isDefalut" gorm:"column:is_defalut"`
	ShortTermInvestment bool   `json:"shortTermInvestment" gorm:"column:is_short_term"`
	LongTermInvestment  bool   `json:"longTermInvestment" gorm:"column:is_long_term"`
	TaxesInvestment     bool   `json:"taxesInvestment" gorm:"column:is_taxes"`
	RetireInvestment    bool   `json:"retireInvestment" gorm:"column:is_retirement"`
	SuiteTestResult     string    `json:"suiteTestResult" gorm:"column:suite_test_result"`
	IsFatca             bool      `json:"isFatca" gorm:"column:is_fatca"`
	FatcaInfo           string    `json:"fatcaInfo" gorm:"column:fatca_info"`
	IsKnowledgeDone     bool      `json:"isKnowledgeDone" gorm:"column:is_knowledge_done"`
	KnowledgeTestResult string    `json:"knowledgeTestResult" gorm:"column:knowledge_test_result"`
}
