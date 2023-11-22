package repository

import (
	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type OpenAccountsRepository struct {
	db *gorm.DB
}

func NewOpenAccountsRepository(db *gorm.DB) *OpenAccountsRepository {
	return &OpenAccountsRepository{db: db}
}

func (e *OpenAccountsRepository) CreateCustomerInformation(columns models.CustomerInformations) error {
	return e.db.Debug().Create(columns).Error
}

func (e *OpenAccountsRepository) CreateCustomerAddresses(columns models.CustomerAddresses) error {
	return e.db.Debug().Create(columns).Error
}

func (e *OpenAccountsRepository) CreateCustomerBookbanks(columns models.CustomerBookbanks) error {
	return e.db.Debug().Create(columns).Error
}

func (e *OpenAccountsRepository) UpdatePersonalInformation(personalInfos models.PersonalInformations, cid string) error {
	tx := e.db.Begin()

	if err := tx.Updates(personalInfos.CustomerInformation).Where(cid).Error; err != nil {
		return err
	}

	if err := tx.Create(personalInfos.CustomerAddresseLists).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(personalInfos.CustomerBookbankLists).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
