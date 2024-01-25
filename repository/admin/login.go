package adminLoginRepository

import (
	"time"

	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type AdminLoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *AdminLoginRepository {
	return &AdminLoginRepository{db}
}

func (r *AdminLoginRepository) Verify(email string) models.AdminLoginRepositoryModel {
	admin := models.AdminLoginRepositoryModel{}
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		admin.Error = err
	}
	return admin
}

func (r *AdminLoginRepository) Create(admin models.AdminLoginRepositoryModel) models.AdminLoginRepositoryModel {
	cust := models.CustomerInformations{}
	if err := r.db.Select("id").Table("customer_informations").Where("email = ?", admin.Email).First(&cust).Error; err != nil {
		admin.Error = err
		return admin
	}

	admin.UserID = cust.AccountID
	now := time.Now()
	end := now.AddDate(0, 0, 90)
	admin.CreatedAt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	admin.ExpiredAt = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, time.Local)

	columns := []string{
		"id",
		"email",
		"password",
		"permission",
		"userId",
		"expiredAt",
		"createdAt",
	}
	if err := r.db.Select(columns).Create(admin).Error; err != nil {
		admin.Error = err
		return admin
	}

	return admin
}
