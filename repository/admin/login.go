package adminRepository

import (
	"fmt"
	"time"

	"github.com/amerikarno/icoApi/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdminLoginRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewLoginRepository(db *gorm.DB, logger *zap.Logger) *AdminLoginRepository {
	return &AdminLoginRepository{db, logger}
}

func (r *AdminLoginRepository) Verify(email string) (models.AdminAccountsModel, error) {
	admin := models.AdminAccountsModel{}
	// var adm string
	columns := []string{
		// "id",
		// "hashed_username",
		"hashed_password",
		// "customer_id",
		// "permission",
		// "expire_at",
	}
	if err := r.db.Debug().Select(columns).
		// Table("admin_accounts").
		Where("hashed_username = ?", email).First(&admin).Error; err != nil {
		r.logger.Error("error while selecting admin account by", zap.String("hashed_username", email), zap.Error(err))
		return admin, err
		// admin.Error = err
	}
	return admin, nil
}

func (r *AdminLoginRepository) Create(admin models.AdminLoginRepositoryModel) (models.AdminLoginRepositoryModel, error) {
	// cust := models.CustomerInformations{}
	// if err := r.db.Select("id").Table("customer_informations").Where("email = ?", admin.HashedUsername).First(&cust).Error; err != nil {
	// 	admin.Error = err
	// 	return admin
	// }

	// admin.CustomerID = cust.AccountID
	now := time.Now()
	end := now.AddDate(0, 0, 90)
	admin.CreatedAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	admin.ExpiredAt = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, time.Local)

	columns := []string{
		"id",
		"hashed_username",
		"encrypted_username",
		"hashed_password",
		"encrypted_password",
		"permission",
		"expire_at",
		"create_att",
	}
	fmt.Printf("admin: %+v\n", admin)
	if err := r.db.Debug().Select(columns).Create(&admin).Error; err != nil {
		r.logger.Error("error whild created admin account", zap.Error(err))
		return admin, err
	}

	return admin, nil
}
