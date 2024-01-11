package risks

import "math"

type ProductRisk struct{}

func NewProductRisksUsecases() *ProductRisk {
	return &ProductRisk{}
}

func (u *ProductRisk) CashExchange() int {
	return 1
}
func (u *ProductRisk) ChangeOwner() int {
	return 4
}
func (u *ProductRisk) InternationalUse() int {
	return 4
}
func (u *ProductRisk) HighRisk() int {
	return 4
}

func (u *ProductRisk) Sum() int {
	return u.CashExchange() + u.CashExchange() + u.InternationalUse() + u.HighRisk()
}

func (u *ProductRisk) Percentage() float64 {
	return math.Floor(((float64(u.Sum()) / 16) * 100) / 100)
}

func (u *ProductRisk) Level() string {
	sum := u.Sum()
	if sum <= 7 {
		return "ความเสี่ยงต่ำ"
	} else if sum <= 11 {
		return "ความเสี่ยงปานกลาง"
	} else {
		return "ความเสี่ยงสูง"
	}
}
