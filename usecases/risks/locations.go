package risks

type LocationRisk struct {
	Country    string
	Province    string
}

func NewLocationRisksUsecases() *LocationRisk {
	return &LocationRisk{}
}

func (r *LocationRisk) ProvinceUsecase() int {
	provinceLists := []string{
		"ยะลา",
		"ปัตตานี",
		"นาราธิวาส",
	}

	for i := range provinceLists {
		if r.Province == provinceLists[i] {
			return 4
		}
	}
	return 0
}

func (r *LocationRisk) CountryUsecase() int {
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
		if r.Country == jurisdictCountryLists[i] {
			return 4
		}
	}

	for i := range FATFCountryLists {
		if r.Country == FATFCountryLists[i] {
			return 2
		}
	}

	return 0
}