package usecases

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/amerikarno/icoApi/models"
)

type OpenAccountUsecases struct {
	oaRepository IOpenAccountsRepository
	external     IExternal
}

func NewOpenAccountUsecases(oaRepository IOpenAccountsRepository, external IExternal) *OpenAccountUsecases {
	return &OpenAccountUsecases{oaRepository: oaRepository, external: external}
}

func (u *OpenAccountUsecases) VerifyEmailFormat(email string) bool {
	emailPattern := regexp.MustCompile("[a-zA-Z0-9]@[a-zA-Z0-9].[a-zA-Z]")
	return emailPattern.MatchString(email)
}

func (u *OpenAccountUsecases) VerifyMobileNoFormat(mobileno string) bool {
	localMobilePattern := regexp.MustCompile("^0[0-9]")
	interMobilePattern := regexp.MustCompile("^66[0-9]")
	if localMobilePattern.MatchString(mobileno) && len(mobileno) == 10 {
		fmt.Printf("local mobile no: %s", mobileno)
	}
	if interMobilePattern.MatchString(mobileno) && len(mobileno) == 11 {
		fmt.Printf("inter mobile no: %s", mobileno)
	}
	return localMobilePattern.MatchString(mobileno) && len(mobileno) == 10 || interMobilePattern.MatchString(mobileno) && len(mobileno) == 11
}

func (u *OpenAccountUsecases) VerifyIDCardNumber(idcard string) bool {
	if len(idcard) == 13 {
		idcardbytes := []byte(idcard)
		sum := 0
		for i := 1; i < 13; i++ {
			digit, _ := strconv.Atoi(string(idcardbytes[i-1]))
			sum += digit * (14 - i)
		}
		last := (11 - (sum % 11)) % 10
		log.Printf("sum: %v, last: %v", sum, last)
		lastdigit, _ := strconv.Atoi(string(idcardbytes[12]))
		return last == lastdigit
	}

	return false
}

func (u *OpenAccountUsecases) CreateCustomerInformationUsecase(idcard models.CustomerInformations) (accountID string, err error) {
	accountID = u.external.GenUuid()
	idcard.AccountID = accountID
	err = u.oaRepository.CreateCustomerInformation(idcard)
	return
}

func (u *OpenAccountUsecases) UpdateCustomerPersonalInformationUsecase(personalInfo models.PersonalInformationPostRequests) (accountID string, err error) {
	accountID = personalInfo.CID
	now := time.Now().Local()
	var pInfo models.PersonalInformations
	pInfo.CustomerInformation.AccountID = accountID
	pInfo.CustomerInformation.SourceOfIncome = personalInfo.Occupation.SourceOfIncome
	pInfo.CustomerInformation.CurrentOccupation = personalInfo.Occupation.CurrentOccupation
	pInfo.CustomerInformation.OfficeName = personalInfo.Occupation.OfficeName
	pInfo.CustomerInformation.TypeOfBusiness = personalInfo.Occupation.TypeOfBusiness
	pInfo.CustomerInformation.PositionName = personalInfo.Occupation.PositionName
	pInfo.CustomerInformation.SalaryRange = personalInfo.Occupation.SalaryRange
	pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressRequest{
		AccountID:       accountID,
		HomeNumber:      personalInfo.RegisteredAddress.HomeNumber,
		VillageNumber:   personalInfo.RegisteredAddress.VillageNumber,
		VillageName:     personalInfo.RegisteredAddress.VillageName,
		SubStreetName:   personalInfo.RegisteredAddress.StreetName,
		StreetName:      personalInfo.RegisteredAddress.StreetName,
		SubDistrictName: personalInfo.RegisteredAddress.SubDistrictName,
		DistrictName:    personalInfo.RegisteredAddress.DistrictName,
		ProvinceName:    personalInfo.RegisteredAddress.ProvinceName,
		ZipCode:         personalInfo.RegisteredAddress.ZipCode,
		CountryName:     personalInfo.RegisteredAddress.CountryName,
		TypeOfAddress:   "r",
		Create:          now,
	})

	if personalInfo.CurrentAddress.TypeOfAddress == "r" {
		pInfo.CustomerAddresseLists[0].TypeOfAddress = "r|c"
	} else {
		pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressRequest{
			AccountID:       accountID,
			HomeNumber:      personalInfo.CurrentAddress.HomeNumber,
			VillageNumber:   personalInfo.CurrentAddress.VillageNumber,
			VillageName:     personalInfo.CurrentAddress.VillageName,
			SubStreetName:   personalInfo.CurrentAddress.StreetName,
			StreetName:      personalInfo.CurrentAddress.StreetName,
			SubDistrictName: personalInfo.CurrentAddress.SubDistrictName,
			DistrictName:    personalInfo.CurrentAddress.DistrictName,
			ProvinceName:    personalInfo.CurrentAddress.ProvinceName,
			ZipCode:         personalInfo.CurrentAddress.ZipCode,
			CountryName:     personalInfo.CurrentAddress.CountryName,
			TypeOfAddress:   "c",
			Create:          now,
		})
	}

	if personalInfo.CurrentAddress.TypeOfAddress == "r" && (personalInfo.OfficeAddress.TypeOfAddress == "r" || personalInfo.OfficeAddress.TypeOfAddress == "c") {
		pInfo.CustomerAddresseLists[0].TypeOfAddress = "r|c|o"
	} else if personalInfo.CurrentAddress.TypeOfAddress == "c" && personalInfo.OfficeAddress.TypeOfAddress == "r" {
		pInfo.CustomerAddresseLists[0].TypeOfAddress = "r|o"
	} else if personalInfo.CurrentAddress.TypeOfAddress == "c" && personalInfo.OfficeAddress.TypeOfAddress == "c" {
		pInfo.CustomerAddresseLists[1].TypeOfAddress = "c|o"
	} else {
		pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressRequest{
			AccountID:       accountID,
			HomeNumber:      personalInfo.OfficeAddress.HomeNumber,
			VillageNumber:   personalInfo.OfficeAddress.VillageNumber,
			VillageName:     personalInfo.OfficeAddress.VillageName,
			SubStreetName:   personalInfo.OfficeAddress.StreetName,
			StreetName:      personalInfo.OfficeAddress.StreetName,
			SubDistrictName: personalInfo.OfficeAddress.SubDistrictName,
			DistrictName:    personalInfo.OfficeAddress.DistrictName,
			ProvinceName:    personalInfo.OfficeAddress.ProvinceName,
			ZipCode:         personalInfo.OfficeAddress.ZipCode,
			CountryName:     personalInfo.OfficeAddress.CountryName,
			TypeOfAddress:   "o",
			Create:          now,
		})
	}
	// err = u.oaRepository.UpdatePersonalInformation(pInfo, accountID)
	fmt.Printf("accountID: %v\npersonal info: %v\n", accountID, pInfo)
	return
}
