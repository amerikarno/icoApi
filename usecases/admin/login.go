package adminLoginUsecases

import (
	"log"
	"time"

	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/token"
	"go.uber.org/zap"
)

type AdminLoginUsecase struct {
	db     IAdminLoginUsecases
	ext    IExternal
	pass   IAdminPassword
	logger *zap.Logger
}

func NewAdminLoginUsecase(db IAdminLoginUsecases,
	ext IExternal,
	pass IAdminPassword,
	logger *zap.Logger) *AdminLoginUsecase {
	return &AdminLoginUsecase{db, ext, pass, logger}
}

func (u *AdminLoginUsecase) Create(req models.AdminCreateRequestModel) (resp models.AdminCreateResponseModel) {
	var admin, create models.AdminLoginRepositoryModel
	var err error
	admin.ID = u.ext.GenUuid()
	admin.HashedUsername = u.ext.HashString(req.Email)
	admin.EncryptedUsername, err = u.ext.Encrypt(req.Email, EncodingKey)
	if err != nil {
		u.logger.Error("error:", zap.Error(err))
		return
	}
	pass := u.pass.GeneratePassword(8)
	admin.HashedPassword = u.ext.HashString(pass)
	admin.EncryptedPassword, err = u.ext.Encrypt(pass, EncodingKey)
	u.logger.Info("username and password", zap.String("username", req.Email), zap.String("password", pass))
	if err != nil {
		u.logger.Error("error:", zap.Error(err))
		return
	}
	admin.Permission = req.Permission
	now := time.Now()
	admin.CreatedAt = now
	admin.ExpiredAt = now.AddDate(1, 0, 0)

	create, err = u.db.Create(admin)
	if err != nil {
		u.logger.Error("error while creating admin account:", zap.Error(err))
		return
	}
	resp = models.AdminCreateResponseModel{
		ID:  create.ID,
		Email: req.Email,
		Password: pass,
	}
	return
}
func (u *AdminLoginUsecase) Verify(email, password string) (resp models.AdminLoginResponse) {
	var login models.AdminAccountsModel
	var err error
	if login, err = u.db.Verify(email); err != nil {
		log.Printf("error: %v", err)
		return
	}

	if password != login.HashedPassword {
		log.Printf("password: %v\nlogin: %v",password, login.HashedPassword)
		return
	}
	user := models.JwtUserModel{
		ID:    login.ID,
		Email: email,
		// Permission:  login.Permission,
		// UserID:      login.CustomerID,
		LoginStatus: Success,
	}

	rc := token.NewRefreshClaims(&user)
	ac := token.NewAccessClaims(&user)

	// var err error
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
