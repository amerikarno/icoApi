package risks

type CustomerRisk struct {
	Occupation string
	PoliticalStatus bool
	AMLlist    string
	IsInThailand bool
	Country    string
}

func NewCustomerRisksUsecases() *CustomerRisk {
	return &CustomerRisk{}
}

func (r *CustomerRisk) OccupationUsecase() int {
	switch r.Occupation {
	case "นักการเมืองในประเทศ", "พระภิกษุ/นักบวช":
		return 4
	case "กิจการครอบครัว", "พนักงานบริษัท", "เจ้าของธุรกิจ/ธุรกิจส่วนตัว", "อาชีพอิสระ", "เกษตรกร":
		return 2
	case "ข้าราชการ", "อาจารย์", "ครู", "แพทย์", "พยาบาล", "พนักงานรัฐวิสาหกิจ":
		return 1
	}
	return 0
}

func (r *CustomerRisk) PoliticalStatusUsecase() int {
	if r.PoliticalStatus {
		return 4
	}
	return 0
}

func (r *CustomerRisk) AMLUsecase() int {
	lists := []string{
		"HR-08-RISK",
		"HR-02",
		"Watch Lis",
		"STR Lis",
		"UN126",
		"UN171",
		"UN223",
		"OFA",
		"H",
		"UN",
		"OFACC",
		"E",
	}
	for i := range lists {
		if r.AMLlist == lists[i] {
			return 4
		}
	}
	return 0
}

func (r *CustomerRisk) IsInThailandUsecase() int {
	if r.IsInThailand {
		return 0
	}
	return 4
}

