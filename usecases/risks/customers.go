package risks

type CustomerRisk struct {
	occupation      string // from occupation in openaccount
	politicalStatus bool   // from CDD
	amllist         string // from CDD
	country         string // registered address in Thailand
	Occupation      int    // from occupation in openaccount
	PoliticalStatus int    // from CDD
	Amllist         int    // from CDD
	Country         int
	Summary         int
}

// type CustomerRiskReturn struct {
// 	Occupation      int // from occupation in openaccount
// 	PoliticalStatus int // from CDD
// 	Amllist         int // from CDD
// 	Country         int
// 	Summary         int
// }

func NewCustomerRisksUsecases(
	occupation string,
	politicalStatus bool,
	amllist string,
	country string,
) *CustomerRisk {
	return &CustomerRisk{
		occupation:      occupation,
		politicalStatus: politicalStatus,
		amllist:         amllist,
		country:         country,
	}
}

func (r *CustomerRisk) getOccupationUsecase() *CustomerRisk {
	switch r.occupation {
	case "นักการเมืองในประเทศ", "พระภิกษุ/นักบวช":
		r.Occupation = 4
		return r
	case "กิจการครอบครัว", "พนักงานบริษัท", "เจ้าของธุรกิจ/ธุรกิจส่วนตัว", "อาชีพอิสระ", "เกษตรกร":
		r.Occupation = 2
		return r
	case "ข้าราชการ", "อาจารย์", "ครู", "แพทย์", "พยาบาล", "พนักงานรัฐวิสาหกิจ":
		r.Occupation = 1
		return r
	}
	return r
}

func (r *CustomerRisk) getPoliticalStatusUsecase() *CustomerRisk {
	// this data from CDD API
	if r.politicalStatus {
		r.PoliticalStatus = 4
		return r
	}
	return r
}

func (r *CustomerRisk) getAMLUsecase() *CustomerRisk {
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
		if r.amllist == lists[i] {
			r.Amllist = 4
			return r
		}
	}
	return r
}

func (r *CustomerRisk) getIsInThailandUsecase() *CustomerRisk {
	if r.country == "Thailand" || r.country == "thailand" || r.country == "ไทย" {
		return r
	}
	r.Country = 4
	return r
}

func (r *CustomerRisk) GetSum() *CustomerRisk {
	// oc := r.getOccupationUsecase()
	// ps := r.getPoliticalStatusUsecase()
	// aml := r.getAMLUsecase()
	// cnt := r.getIsInThailandUsecase()
	// r.Occupation = oc.Occupation
	// r.PoliticalStatus = ps.PoliticalStatus
	// r.Amllist = aml.Amllist
	// r.Country = cnt.Country
	// r.Summary = oc.Occupation + ps.PoliticalStatus + aml.Amllist + cnt.Country
	r.Summary = r.getOccupationUsecase().Occupation + r.getAMLUsecase().Amllist + r.getPoliticalStatusUsecase().PoliticalStatus + r.getIsInThailandUsecase().Country
	return r
}
