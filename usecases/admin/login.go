package adminLoginUsecases

import (
	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/token"
)

type AdminLoginUsecase struct {
	db   IAdminLoginUsecases
	ext  IExternal
	pass IAdminPassword
}

func NewAdminLoginUsecase(db   IAdminLoginUsecases,
	ext  IExternal,
	pass IAdminPassword) *AdminLoginUsecase {
	return &AdminLoginUsecase{db, ext, pass}
}

func (u *AdminLoginUsecase) Create(email, permission string) (resp models.JwtUserModel) {
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
	resp = models.JwtUserModel{
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
		return
	}

	if password != login.Password {
		resp.Error = login.Error
		return

	}
	user := models.JwtUserModel{
		ID:          login.ID,
		Email:       email,
		Permission:  login.Permission,
		UserID:      login.UserID,
		LoginStatus: Success,
	}

	rc := token.NewRefreshClaims(&user)
	ac := token.NewAccessClaims(&user)

	var err error
	resp.RequestToken, err = rc.JwtString()
	if err != nil {
		resp.Error = err
		return
	}

	resp.AccessToken, err = ac.JwtString()
	if err != nil {
		resp.Error = err
		return
	}

	return
}
