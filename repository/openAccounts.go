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

func (e *OpenAccountsRepository) CreateCustomerInformation(datas models.CustomerInformations) error {
	columns := []string{"id", "th_title", "th_name", "th_surname", "en_title", "en_name", "en_surname", "email", "mobile_no", "personal_agreement", "birth_date", "marriage_status", "id_card", "laser_code", "create_at"}
	return e.db.Debug().Select(columns).Create(datas).Error
	// return e.db.Debug().Select("id", "th_title", "th_name", "th_surname", "en_title", "en_name", "en_surname", "email", "mobile_no", "personal_agreement", "birth_date", "marriage_status", "id_card", "laser_code", "create_at").Create(datas).Error
}

func (e *OpenAccountsRepository) CreateCustomerAddresses(datas models.CustomerAddresses) error {
	return e.db.Debug().Create(datas).Error
}

func (e *OpenAccountsRepository) CreateCustomerBookbanks(datas models.CustomerBookbanks) error {
	return e.db.Debug().Create(datas).Error
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
func (e *OpenAccountsRepository) CheckReisteredCitizenID(cid string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("id_card=?", cid).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking citizen id: %v", err)
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

func (e *OpenAccountsRepository) CreateCustomerExams(customerExams models.CustomerExamsRequest) error {
	tx := e.db.Begin()

	if err := tx.Create(customerExams).Error; err != nil {
		log.Printf("customer exams update error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) CreateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	columns := []string{
		"id",
		"token",
		"is_confirm",
		"confirm_types",
		"create_at",
		"expire_at",
	}
	tx := e.db.Begin()

	if err := tx.Select(columns).Create(customerConfirms).Error; err != nil {
		log.Printf("create customer confirm error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) UpdateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	columns := []string{
		"is_confirm",
		"confirm_at",
	}
	tx := e.db.Begin()

	if err := tx.Select(columns).Updates(customerConfirms).Where(customerConfirms.TokenID).Error; err != nil {
		log.Printf("create customer confirm error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) QueryCustomerConfirmsExpireDT(tokenID string) models.CustomerConfirmsRequest {
	custConfirmDetail := models.CustomerConfirmsRequest{}

	if err := e.db.Debug().Where(tokenID).First(&custConfirmDetail).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking customer token %s, error: %v", tokenID, err)
	}

	return custConfirmDetail
}
