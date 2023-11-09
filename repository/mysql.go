package repository

import (
	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type MysqlRepository struct{
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (e *MysqlRepository) UpdateExchangeRates(columns []models.PostIDcard) error {
	return e.db.Debug().Create(columns).Error
}