package adminLoginUsecases

import "github.com/amerikarno/icoApi/models"

type IAdminLoginUsecases interface {
	Create(admin models.AdminLoginRepositoryModel) models.AdminLoginRepositoryModel
	Verify(email string) models.AdminLoginRepositoryModel
}

type IExternal interface {
	GenUuid() (uid string)
}

type IAdminPassword interface {
	IsValidPassword(password string) bool
	GeneratePassword(length int) string
}