package adminRepository

import (
	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type Customers struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *Customers {
	return &Customers{db: db}
}

func (r *Customers) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	var custInfos  []models.CustomerInformations
	if err := r.db.Find(&custInfos).Error; err != nil {
		return nil, err
	}


	return customers, nil
}
