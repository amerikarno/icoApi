package usecases

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/amerikarno/icoApi/models"
)

type Usecases struct{}

func NewUsecases() *Usecases {
	return &Usecases{}
}

func (u *Usecases) VerifyEmailFormat(email string) bool {
	emailPattern := regexp.MustCompile("[a-zA-Z0-9]@[a-zA-Z0-9].[a-zA-Z]")
	return emailPattern.MatchString(email)
}

func (u *Usecases) VerifyMobileNoFormat(mobileno string) bool {
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

func (u *Usecases) VerifyIDCardNumber(idcard string) bool {
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

func (u *Usecases) PostIDcardService(idcard models.PostIDcard) (err error) {


	return
}
