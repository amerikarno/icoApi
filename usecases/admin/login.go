package adminLoginUsecases

import (
	"github.com/amerikarno/icoApi/models"
)

type AdminLoginUsecase struct {
	db   IAdminLoginUsecases
	ext  IExternal
	pass IAdminPassword
}

func NewAdminLoginUsecase() *AdminLoginUsecase {
	return &AdminLoginUsecase{}
}

func (u *AdminLoginUsecase) Create(email, permission string) (resp models.AdminLoginResponse) {
	var admin, create models.AdminLoginRepositoryModel
	admin.ID = u.ext.GenUuid()
	admin.Email = email
	admin.Password = u.pass.GeneratePassword(13)
	admin.Permission = permission

	if create = u.db.Create(admin); create.Error != nil {
		resp.Error = create.Error
		resp.LoginStatus = Failed
		return
	}
	resp = models.AdminLoginResponse{
		ID:          admin.ID,
		Email:       email,
		Permission:  permission,
		UserID:      create.UserID,
		LoginStatus: Success,
	}
	return
}
func (u *AdminLoginUsecase) Verify(email, password string) (resp models.AdminLoginResponse) {
	var login models.AdminLoginRepositoryModel
	if login = u.db.Verify(email); login.Error != nil {
		resp.Error = login.Error
		resp.LoginStatus = Failed
		return
	}
	if password != login.Password {
		resp.Error = login.Error
		resp.LoginStatus = Failed
		return

	}
	resp = models.AdminLoginResponse{
		ID:          login.ID,
		Email:       email,
		Permission:  login.Permission,
		UserID:      login.UserID,
		LoginStatus: Success,
	}
	return
}
