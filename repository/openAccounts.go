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

func (e *OpenAccountsRepository) CreateOpenAccount(columns models.IDCardOpenAccounts) error {
	return e.db.Debug().Create(columns).Error
}
