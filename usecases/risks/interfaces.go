package risks

import "github.com/amerikarno/icoApi/models"

type Risk interface {
	GetAllRiskCountry() (riskCountries []models.RiskCountryModel)
	GetAllRiskOccupation() (riskOccupations []models.RiskOccupationModel)
}
