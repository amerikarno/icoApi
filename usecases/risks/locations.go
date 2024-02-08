package risks

type LocationRisk struct {
	province      string // from registered / current / office addresses
	country       string // from registered / current /
	ProvincePoint int
	CountryPoint  int
	SummaryPoint  int
}

func NewLocationRisksUsecases(
	province string, // from registered / current / office addresses
	country string, // from registered / current /
) *LocationRisk {
	return &LocationRisk{province: province, country: country}
}

func (r *LocationRisk) ProvinceUsecase() *LocationRisk {
	provinceLists := []string{
		"ยะลา",
		"ปัตตานี",
		"นาราธิวาส",
	}

	for i := range provinceLists {
		if r.province == provinceLists[i] {
			r.ProvincePoint = 4
			return r
		}
	}
	return r
}

func (r *LocationRisk) CountryUsecase() *LocationRisk {
	jurisdictCountryLists := []string{
		"Albania",
		"Barbados",
		"Burkina Faso",
		"Cambodia",
		"Cayman Islands",
		"Gibraltar",
		"Haiti",
		"Jamaica",
		"Jordan",
		"Mali",
		"Morocco",
		"Myanmar",
		"Nicaragua",
		"Pakistan",
		"Panama",
		"Philippines",
		"Senegal",
		"South Sudan",
		"Syria",
		"Türkiye",
		"Uganda",
		"United Arab Emirates",
		"Yemen",
	}

	FATFCountryLists := []string{
		"Argentina",
		"Australia",
		"Austria",
		"Belgium",
		"Brazil",
		"Canada",
		"China",
		"Denmark",
		"European Commission",
		"Finland",
		"France",
		"Germany",
		"Greece",
		"Gulf Co-operation Council",
		"Hong Kong, China",
		"Iceland",
		"India",
		"Ireland",
		"Israel",
		"Italy",
		"Japan",
		"Republic of Korea",
		"Luxembourg",
		"Malaysia",
		"Mexico",
		"Netherlands, Kingdom of",
		"New Zealand",
		"Norway",
		"Portugal",
		"Russian Federation",
		"Saudi Arabia",
		"Singapore",
		"South Africa",
		"Spain",
		"Sweden",
		"Switzerland",
		"Türkiye",
		"United Kingdom",
		"United States",
	}

	for i := range jurisdictCountryLists {
		if r.country == jurisdictCountryLists[i] {
			r.CountryPoint = 4
			return r
		}
	}

	for i := range FATFCountryLists {
		if r.country == FATFCountryLists[i] {
			r.CountryPoint = 2
			return r
		}
	}

	return r
}

func (r *LocationRisk) GetSum() *LocationRisk {
	r.SummaryPoint = r.CountryPoint + r.ProvincePoint
	return r
}
