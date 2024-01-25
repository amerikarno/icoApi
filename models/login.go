package models

import "time"

type AdminLoginRepositoryModel struct {
	ID         string    `json:"id" gorm:"id"`
	Email      string    `json:"email" gorm:"email"`
	Password   string    `json:"password" gorm:"password"`
	Permission string    `json:"permission" gorm:"permission"`
	UserID     string       `json:"userId" gorm:"user_id"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"updated_at"`
	ExpiredAt  time.Time `json:"expiredAt" gorm:"expired_at"`
	CreatedAt  time.Time `json:"createdAt" gorm:"created_at"`

	Error error
}

type AdminLoginResponse struct {
	ID          string `json:"id" gorm:"id"`
	Email       string `json:"email" gorm:"email"`
	Permission  string `json:"permission" gorm:"permission"`
	UserID      string    `json:"userId" gorm:"user_id"`
	LoginStatus string `json:"loginStatus" gorm:"login_status"`
	Error       error
}
