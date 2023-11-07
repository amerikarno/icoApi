package main

import (
	"fmt"

	"github.com/amerikarno/icoApi/repository"
)

func main() {
	filename := "json/api_province_with_amphure_tambon.json"
	repo := repository.NewPATRepository(filename)
	repo.LoadPAT()
	provinces, amphures, tambons, zipcode := repo.GetProvinceAmphureTambonLists()
	fmt.Printf("province: %v", provinces)
	fmt.Printf("amphure: %v", amphures["กรุงเทพมหานคร"])
	fmt.Printf("tambon: %v", tambons["เขตบางกอกน้อย"])
	fmt.Printf("zipcode: %v", zipcode["กรุงเทพมหานครเขตบางกอกน้อยบางขุนศรี"])
}