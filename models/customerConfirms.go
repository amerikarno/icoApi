package models

import "time"

type CustomerConfirmsRequest struct {
	AccountID    string    `json:"id" gorm:"column:customer_id"`
	TokenID      string    `json:"tokenID" gorm:"column:token"`
	IsConfirm    bool      `json:"isConfirm" gorm:"column:is_confirm"`
	ConfirmTypes string    `json:"confirmTypes" gorm:"column:confirm_types"`
	CreateAt     time.Time `json:"createAt" gorm:"column:create_at"`
	ExpireAt     time.Time `json:"expireAt" gorm:"column:expire_at"`
	ConfirmAt    time.Time `json:"confirmAt" gorm:"column:confirm_at"`
}

func (CustomerConfirmsRequest) TableName() string { return "customer_confirms" }
