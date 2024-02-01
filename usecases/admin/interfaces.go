package adminLoginUsecases

import "github.com/amerikarno/icoApi/models"

type IAdminLoginUsecases interface {
	Create(admin models.AdminLoginRepositoryModel) (models.AdminLoginRepositoryModel, error)
	Verify(email string) (models.AdminAccountsModel, error)
}

type IExternal interface {
	GenUuid() (uid string)
	HashString(input string) string
	Encrypt(plainText string, keyString string) (string, error)
	Decrypt(cipherText string, keyString string) (string, error)
}

type IAdminPassword interface {
	IsValidPassword(password string) bool
	GeneratePassword(length int) string
}
