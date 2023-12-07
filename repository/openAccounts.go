package repository

import (
	"log"

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

func (e *OpenAccountsRepository) CheckReisteredEmail(email string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("email=?", email).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking registered email: %v", err)
	}
	return cust
}

func (e *OpenAccountsRepository) CheckReisteredMobileNo(mobileno string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("mobile_no=?", mobileno).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking registered mobile number: %v", err)
	}
	return cust
}

func (e *OpenAccountsRepository) UpdatePersonalInformation(personalInfos models.PersonalInformations, cid string) error {
	tx := e.db.Begin()

	if err := tx.Updates(personalInfos.CustomerInformation).Where(cid).Error; err != nil {
		log.Printf("error1: %v", err)
		return err
	}

	if err := tx.Create(personalInfos.CustomerAddresseLists).Error; err != nil {
		log.Printf("error2: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Create(personalInfos.CustomerBookbankLists).Error; err != nil {
		log.Printf("error3: %v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
