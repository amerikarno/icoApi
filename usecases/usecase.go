package usecases

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

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

func (u *OpenAccountUsecases) UpdateCustomerPersonalInformationUsecase(personalInfo models.PersonalInformations) (accountID string, err error) {
	accountID = personalInfo.CustomerInformation.AccountID
	err = u.oaRepository.UpdatePersonalInformation(personalInfo)
	return
}
