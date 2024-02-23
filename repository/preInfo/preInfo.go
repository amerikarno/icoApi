package preinfo

import (
	"fmt"

	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type PreInfoRepository struct {
	models.SettingTitles
	db *gorm.DB
}

func NewPreInfoRepository(db *gorm.DB) *PreInfoRepository {
	return &PreInfoRepository{db: db}
}

func (repo *PreInfoRepository) GetAll() ([]models.SettingTitles, error) {
	data := []models.SettingTitles{}

	if err := repo.db.Find(&data).Error; err != nil {
		return nil, err
	}

	fmt.Printf("data: %+v", data)

	return data, nil
}
