package preinfo

import (
	"log"

	"github.com/amerikarno/icoApi/models"
	preinfo "github.com/amerikarno/icoApi/repository/preInfo"
)

type PreInfoUsecase struct {
	repo *preinfo.PreInfoRepository
}

func NewPreInfoUsecase(repo *preinfo.PreInfoRepository) *PreInfoUsecase {
	return &PreInfoUsecase{repo: repo}
}

func (u *PreInfoUsecase) GetTitles() ([]models.SettingTitles) {
	titles, err := u.repo.GetAll()
	if err != nil {
		log.Printf("error: %v", err)
	}
	return titles
}
