package models

import (
	"time"
)

type AdminLoginRepositoryModel struct {
	ID                string    `json:"id" gorm:"column:id"`
	HashedUsername    string    `json:"hashedUsername" gorm:"column:hashed_username"`
	EncryptedUsername string    `json:"encryptedUsername" gorm:"column:encrypted_usernames"`
	HashedPassword    string    `json:"hashedPassword" gorm:"column:hashed_password"`
	EncryptedPassword string    `json:"encryptedPassword" gorm:"column:encrypted_passwords"`
	CustomerID        string    `json:"customerId" gorm:"column:customer_id"`
	Permission        string    `json:"permission" gorm:"column:permission"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:update_at"`
	ExpiredAt         time.Time `json:"expiredAt" gorm:"column:expire_at"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:create_at"`
	ConfirmedAt       time.Time `json:"confirmedAt" gorm:"column:confirm_at"`
}

func (AdminLoginRepositoryModel) TableName() string {
	return "admin_accounts"
}

type JwtUserModel struct {
	ID          string `json:"id" gorm:"column:id"`
	Email       string `json:"email" gorm:"column:email"`
	Permission  string `json:"permission" gorm:"column:permission"`
	UserID      string `json:"userId" gorm:"column:user_id"`
	LoginStatus string `json:"loginStatus" gorm:"column:login_status"`
	Error       error
}

type AdminLoginResponse struct {
	AccessToken  string `json:"accessToken" gorm:"column:access_token"`
	RequestToken string `json:"requestToken" gorm:"column:request_token"`
	Error        error
}

type AdminLoginRequest struct {
	HashedUsername string `json:"hashedUsername"`
	HashedPassword string `json:"hashedPassword"`
}

type AdminAccountsModel struct {
	ID             string `json:"id" gorm:"column:id"`
	HashedUsername string `json:"hashedUsername" gorm:"column:hashed_username"`
	HashedPassword string `json:"hashedPassword" gorm:"column:hashed_password"`
}

func (AdminAccountsModel) TableName() string { return "admin_accounts" }

type AdminCreateRequestModel struct {
	Email      string `json:"email" gorm:"column:email"`
	Permission string `json:"permission" gorm:"column:permission"`
}

func (AdminCreateRequestModel) TableName() string { return "admin_accounts" }

type AdminCreateResponseModel struct {
	ID       string `json:"id" gorm:"column:id"`
	Email    string `json:"email" gorm:"column:hashedUsername"`
	Password string `json:"password" gorm:"column:hashedPassword"`
}
