package repository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/amerikarno/icoApi/models"
)

type PATRepository struct {
	filename string
	provincelists []models.ProvinceList
}

func NewPATRepository(filename string) *PATRepository { return &PATRepository{filename: filename} }

func (r *PATRepository) LoadPAT() *PATRepository {
	file, err := os.ReadFile(r.filename)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(file, &r.provincelists); err != nil {
		log.Fatal(err)
	}

	return r
}

func (r *PATRepository) GetProvinceAmphureTambonLists() (provinces []string, amphure, tambon map[string][]string, zipcode map[string]int) {
	amphure = make(map[string][]string)
	tambon = make(map[string][]string)
	zipcode = make(map[string]int)

	for i := range r.provincelists {
		provincelist := &r.provincelists[i]
		provinces = append(provinces, provincelist.NameTh)
		for j := range provincelist.Amphure {
			amphurelist := &provincelist.Amphure[j]
			amphure[provincelist.NameTh] = append(amphure[provincelist.NameTh],amphurelist.NameTh)
			for k := range amphurelist.Tambon {
				tambonlist := &amphurelist.Tambon[k]
				tambon[amphurelist.NameTh] = append(tambon[amphurelist.NameTh], tambonlist.NameTh)
				zipcode[provincelist.NameTh+amphurelist.NameTh+tambonlist.NameTh] = tambonlist.ZipCode
			}
		}
	}
	return
}
