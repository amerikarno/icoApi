package handlers

import "github.com/amerikarno/icoApi/models"

type IAdminUsecases interface {
	Verify(email, password string) (resp models.AdminLoginResponse)
	Create(req models.AdminCreateRequestModel) (resp models.AdminCreateResponseModel)
}
