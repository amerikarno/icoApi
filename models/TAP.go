package models

type TAP struct {
	ProvinceID      int    `json:"province_id,omitempty"`
	ProvinceName    string `json:"province_name,omitempty"`
	DistrictID      int    `json:"district_id,omitempty"`
	DistrictName    string `json:"district_name,omitempty"`
	SubDistrictID   int    `json:"sub_district_id,omitempty"`
	SubDistrictName string `json:"sub_district_name,omitempty"`
	AddressText     string `json:"addressText,omitempty"`
	Zipcode         string `json:"zipcode,omitempty"`
}

type TAPresponse struct {
	ResultData []TAP `json:"resultData,omitempty"`
	Success    bool `json:"success,omitempty"`
	Error      any  `json:"error,omitempty"`
	Message    any  `json:"message,omitempty"`
	NextModule any  `json:"next_module,omitempty"`
}
