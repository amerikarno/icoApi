package risks

type ProductRiskUsecase struct {}

func NewProductRisksUsecases() *ProductRiskUsecase {
	return &ProductRiskUsecase{}
}

func (r *ProductRiskUsecase) Get() int {
	return 4
}

