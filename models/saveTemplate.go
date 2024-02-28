package models

type SaveTemplate struct {
	Module     string  `json:"module,omitempty"`
	Session    int     `json:"session,omitempty"`
	PreInfo
}
