package risks

import "log"

type RiskAssessmentUsecase struct {
	ProductRisk  IProductRisk
	CustomerRisk ICustomerRisk
	LocationRisk ILocationRisk
	ChannelRisk  IChannelRisk
}

func NewRiskAssessmentUsecase(
	ProductRisk IProductRisk,
	CustomerRisk ICustomerRisk,
	LocationRisk ILocationRisk,
	ChannelRisk IChannelRisk,
) *RiskAssessmentUsecase {
	return &RiskAssessmentUsecase{ProductRisk, CustomerRisk, LocationRisk, ChannelRisk}
}

func (r *RiskAssessmentUsecase) GetSum() {
	product := r.ProductRisk.Get()
	customer := r.CustomerRisk.GetSum()
	location := r.LocationRisk.GetSum()
	channel := r.ChannelRisk.GetChannelRisk()
	sum := product + customer.Summary + location.SummaryPoint + channel.ChannelPoint
	log.Println(sum)
}
