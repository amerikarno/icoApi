package usecases

import (
	"github.com/amerikarno/icoApi/models"
)

//go:generate mockgen -source=interfaces.go -destination=mock/mock.go -package=mock
type IOpenAccountsRepository interface {
	CreateCustomerInformation(columns models.CustomerInformations) error
	UpdatePersonalInformation(personalInfos models.PersonalInformations, cid string) error
	CheckReisteredEmail(email string) models.CustomerInformations
	CheckReisteredMobileNo(mobileno string) models.CustomerInformations
}

type IExternal interface {
	GenUuid() (uid string)
}
