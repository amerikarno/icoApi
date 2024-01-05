package usecases

import (
	"fmt"
	"log"
	"net/mail"
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
	// emailPattern1 := regexp.MustCompile("[a-zA-Z0-9]@[a-zA-Z0-9][.][a-zA-Z]")
	// return emailPattern1.Match([]byte(email))
	if _, err := mail.ParseAddress(email); err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}
	return true
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

func (u *OpenAccountUsecases) CheckedEmailUsecase(email string) bool {
	customer := u.oaRepository.CheckReisteredEmail(email)
	return len(customer.AccountID) != 0
}
func (u *OpenAccountUsecases) CheckedMobileUsecase(mobileno string) bool {
	customer := u.oaRepository.CheckReisteredMobileNo(mobileno)
	return len(customer.AccountID) != 0
}
func (u *OpenAccountUsecases) CheckedCitizenIDUsecase(citizenID string) bool {
	customer := u.oaRepository.CheckReisteredCitizenID(citizenID)
	return len(customer.AccountID) != 0
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
	pInfo.CustomerInformation.Update = now
	pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressResponse{
		AccountID:           accountID,
		HomeNumber:          personalInfo.RegisteredAddress.HomeNumber,
		VillageNumber:       personalInfo.RegisteredAddress.VillageNumber,
		VillageName:         personalInfo.RegisteredAddress.VillageName,
		SubStreetName:       personalInfo.RegisteredAddress.StreetName,
		StreetName:          personalInfo.RegisteredAddress.StreetName,
		SubDistrictName:     personalInfo.RegisteredAddress.SubDistrictName,
		DistrictName:        personalInfo.RegisteredAddress.DistrictName,
		ProvinceName:        personalInfo.RegisteredAddress.ProvinceName,
		ZipCode:             personalInfo.RegisteredAddress.ZipCode,
		CountryName:         personalInfo.RegisteredAddress.CountryName,
		IsRegisteredAddress: true,
		Create:              now,
	})

	if personalInfo.CurrentAddress.TypeOfAddress == "SelectedCurrentAddressEnum.registered" {
		pInfo.CustomerAddresseLists[0].IsCurrentAddress = true
		// fmt.Printf("current1: %+v\n", personalInfo.CurrentAddress.TypeOfAddress)
	} else if personalInfo.CurrentAddress.TypeOfAddress == "SelectedCurrentAddressEnum.current" {
		fmt.Printf("current2: %+v\n", personalInfo.CurrentAddress.TypeOfAddress)
		pInfo.CustomerAddresseLists[0].IsCurrentAddress = false
		pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressResponse{
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
			// TypeOfAddress:   "c",
			IsCurrentAddress: true,
			Create:           now,
		})
	}

	fmt.Printf("address list: %+v\n", pInfo.CustomerAddresseLists)

	if personalInfo.CurrentAddress.TypeOfAddress == "SelectedCurrentAddressEnum.registered" &&
		(personalInfo.OfficeAddress.TypeOfAddress == "SelectedOfficeAddressEnum.registered" || personalInfo.OfficeAddress.TypeOfAddress == "SelectedOfficeAddressEnum.current") {
		// pInfo.CustomerAddresseLists[0].TypeOfAddress = "r|c|o"
		pInfo.CustomerAddresseLists[0].IsOfficeAddress = true
		fmt.Printf("office1: %+v\n", personalInfo.OfficeAddress.TypeOfAddress)
	} else if personalInfo.CurrentAddress.TypeOfAddress == "SelectedCurrentAddressEnum.current" && personalInfo.OfficeAddress.TypeOfAddress == "SelectedOfficeAddressEnum.registered" {
		// pInfo.CustomerAddresseLists[0].TypeOfAddress = "r|o"
		pInfo.CustomerAddresseLists[0].IsOfficeAddress = true
		pInfo.CustomerAddresseLists[1].IsOfficeAddress = false
		fmt.Printf("office2: %+v\n", personalInfo.OfficeAddress.TypeOfAddress)
	} else if personalInfo.CurrentAddress.TypeOfAddress == "SelectedCurrentAddressEnum.current" && personalInfo.OfficeAddress.TypeOfAddress == "SelectedOfficeAddressEnum.current" {
		// pInfo.CustomerAddresseLists[1].TypeOfAddress = "c|o"
		pInfo.CustomerAddresseLists[0].IsOfficeAddress = false
		pInfo.CustomerAddresseLists[1].IsOfficeAddress = true
		fmt.Printf("office3: %+v\n", personalInfo.OfficeAddress.TypeOfAddress)
	} else if personalInfo.OfficeAddress.TypeOfAddress == "SelectedOfficeAddressEnum.office" {
		fmt.Printf("office4: %+v\n", personalInfo.OfficeAddress.TypeOfAddress)
		pInfo.CustomerAddresseLists[0].IsOfficeAddress = false
		pInfo.CustomerAddresseLists[1].IsOfficeAddress = false
		pInfo.CustomerAddresseLists = append(pInfo.CustomerAddresseLists, models.CustomerAddressResponse{
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
			// TypeOfAddress:   "o",
			IsOfficeAddress: true,
			Create:          now,
		})
	}

	pInfo.CustomerBookbankLists = append(pInfo.CustomerBookbankLists, models.CustomerBookbankResponse{
		AccountID:         accountID,
		BankName:          personalInfo.FirstBankAccount.BankName,
		BankBranchName:    personalInfo.FirstBankAccount.BankBranchName,
		BankAccountNumber: personalInfo.FirstBankAccount.BankAccountNumber,
		IsDefalut:         true,
		IsDeposit:         true,
		IsWithdraw:        true,
		Create:            now,
	})

	if personalInfo.SecondBankAccount.BankName != "" {
		pInfo.CustomerBookbankLists = append(pInfo.CustomerBookbankLists, models.CustomerBookbankResponse{
			AccountID:         accountID,
			BankName:          personalInfo.SecondBankAccount.BankName,
			BankBranchName:    personalInfo.SecondBankAccount.BankBranchName,
			BankAccountNumber: personalInfo.SecondBankAccount.BankAccountNumber,
			IsDeposit:         true,
			Create:            now,
		})
	}

	fmt.Printf("address list: %+v\n", pInfo.CustomerAddresseLists)
	if err = u.oaRepository.UpdatePersonalInformation(pInfo, accountID); err != nil {
		accountID = ""
		return
	}
	fmt.Printf("accountID: %v\npersonal info: %v\n", accountID, pInfo)
	return
}

func (u *OpenAccountUsecases) CreateCustomerExamsUsecase(customerExams models.CustomerExamsRequest) (accountID string, err error) {
	accountID = customerExams.AccountID
	customerExams.CreateAt = time.Now().Local()
	if err = u.oaRepository.CreateCustomerExams(customerExams); err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (u *OpenAccountUsecases) CreateCustomerConfirmsUsecase(customerConfirms models.CustomerConfirmsRequest) (tokenID string, err error) {
	now := time.Now().Local()
	customerConfirms.CreateAt = now
	customerConfirms.ExpireAt = now.Add(time.Hour * 24)
	customerConfirms.TokenID = u.external.GenUuid()
	customerConfirms.IsConfirm = false
	tokenID = customerConfirms.TokenID

	if err = u.oaRepository.CreateCustomerConfirms(customerConfirms); err != nil {
		log.Println(err.Error())
		return
	}

	return
}

func (u *OpenAccountUsecases) UpdateConfirmsUsecase(token string) (tokenID string, err error) {
	now := time.Now().Local()
	query := u.oaRepository.QueryCustomerConfirmsExpireDT(token)

	if query.ExpireAt.Before(now) {
		err = fmt.Errorf("expired: %v, now: %v", query.ExpireAt, now)
		log.Println(err.Error())
		return
	}
	customerConfirms := models.CustomerConfirmsRequest{}
	customerConfirms.ConfirmAt = now
	customerConfirms.IsConfirm = true

	if err = u.oaRepository.UpdateCustomerConfirms(customerConfirms); err != nil {
		log.Println(err.Error())
		return
	}

	return token, nil
}
