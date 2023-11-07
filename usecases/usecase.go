package usecases

import (
	"fmt"
	"regexp"
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
