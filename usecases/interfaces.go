package usecases

import "github.com/amerikarno/icoApi/models"

//go:generate mockgen -source=interfaces.go -destination=mock/mock.go -package=mock
type IOpenAccountsRepository interface {
	CreateOpenAccount(columns models.IDCardOpenAccounts) error
}

type IExternal interface {
	GenUuid() (uid string)
}