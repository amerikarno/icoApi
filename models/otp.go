package models

type OTPNextModule struct {
	Module     string `json:"module,omitempty"`
	GroupType  int    `json:"group_type,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	Page       string `json:"page,omitempty"`
	Sub        string `json:"sub,omitempty"`
	Step       string `json:"step,omitempty"`
	URL        string `json:"url,omitempty"`
	Success    bool   `json:"success,omitempty"`
	Error      string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	NextModule string `json:"next_module,omitempty"`
}

type OTPdata struct {
	UID         string  `json:"uid,omitempty"`
	Email       string  `json:"email,omitempty"`
	Mobile      string  `json:"mobile,omitempty"`
	RefID       string  `json:"ref_id,omitempty"`
	VerifySms   string  `json:"verify_sms,omitempty"`
	VerifyEmail string  `json:"verify_email,omitempty"`
	Info        OTPinfo `json:"info,omitempty"`
	Success     bool    `json:"success,omitempty"`
	Error       string  `json:"error,omitempty"`
	Message     string  `json:"message,omitempty"`
	NextModule  string  `json:"next_module,omitempty"`
}

type OTPinfo struct {
	ProjectKey string `json:"project_key,omitempty"`
	Phone      string `json:"phone,omitempty"`
	RefCode    string `json:"ref_code,omitempty"`
	Status     string `json:"Status,omitempty"`
	TxID       string `json:"TxID,omitempty"`
	Message    string `json:"Message,omitempty"`
}

type OTPresponse struct {
	Data             OTPdata       `json:"data,omitempty"`
	UID              string        `json:"uid,omitempty"`
	Module           string        `json:"module,omitempty"`
	CreatedTimestamp string        `json:"created_timestamp,omitempty"`
	UpdatedTimestamp string        `json:"updated_timestamp,omitempty"`
	Email            string        `json:"email,omitempty"`
	Name             string        `json:"name,omitempty"`
	Surname          string        `json:"surname,omitempty"`
	SessionOTP       string        `json:"sessionOTP,omitempty"`
	Success          bool          `json:"success,omitempty"`
	Error            string        `json:"error,omitempty"`
	Message          string        `json:"message,omitempty"`
	NextModule       OTPNextModule `json:"next_module,omitempty"`
}
