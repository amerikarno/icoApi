package risks

import "github.com/amerikarno/icoApi/models"

type Risk interface {
	GetAllRiskCountry() (riskCountries []models.RiskCountryModel)
	GetAllRiskOccupation() (riskOccupations []models.RiskOccupationModel)
}

type ICustomerRisk interface {
	GetOccupationUsecase() *CustomerRisk
	GetPoliticalStatusUsecase() *CustomerRisk
	GetAMLUsecase() *CustomerRisk
	GetIsInThailandUsecase() *CustomerRisk
	GetSum() *CustomerRisk
}

type IProductRisk interface {
	Get() int
}

type ILocationRisk interface {
	GetSum() *LocationRisk
}

type IChannelRisk interface {
	GetChannelRisk() *ChannelRisk
}
