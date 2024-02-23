package models

type CheckExist struct {
	Success    bool   `json:"success,omitempty"`
	Error      string `json:"error,omitempty"`
	Message    any    `json:"message,omitempty"`
	NextModule any    `json:"next_module,omitempty"`
}