package models

type PreInfo struct {
	PrefixT           string `json:"PrefixT,omitempty"`
	FNameT            string `json:"FNameT,omitempty"`
	LNameT            string `json:"LNameT,omitempty"`
	PrefixE           string `json:"PrefixE,omitempty"`
	FNameE            string `json:"FNameE,omitempty"`
	LNameE            string `json:"LNameE,omitempty"`
	Email             string `json:"Email,omitempty"`
	Mobile            string `json:"Mobile,omitempty"`
	MobileCountryCode any    `json:"MobileCountryCode,omitempty"`
	Token             any    `json:"token,omitempty"`
	StatusModule      string `json:"status_module,omitempty"`
}