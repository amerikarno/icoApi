package adminLoginRepository

import (
	"crypto/rand"
	"math/big"
	"regexp"
)

/*
Password policy

1. password length at least 8 characters
2. must have
  - a-z at least 1 character
  - A-Z at least 1 character
  - 0-9 at least 1 character
  - special characters  (@#$%*&) at least 1 character
  - cannot duplicate character more than 3 charactor such as aaaa1111

3. new password must not be duplicated with last 4 passwords
4. password will be expired after 90 days
5. alert before 14days when password expire
6. account will be locked after incorrect password more than 3 times and need to contact IT to unlock
7. hidden password as default when user login
*/
type AdminPassword struct{}

func NewAdminPassword() *AdminPassword {
	return &AdminPassword{}
}

func (r *AdminPassword) GeneratePassword(length int) string {
	var password string
	charTypes := [][]rune{
		[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), // Uppercase
		[]rune("abcdefghijklmnopqrstuvwxyz"), // Lowercase
		[]rune("0123456789"),                 // Digits
		[]rune("@#$%*&()"),                   // Special characters
	}
	typeCount := make([]int, len(charTypes))

	for len(password) < length {
		charTypeIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charTypes))))
		charType := charTypes[charTypeIndex.Int64()]
		newCharIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charType))))
		newChar := charType[newCharIndex.Int64()]

		if canAddChar(password, newChar, 3) {
			password += string(newChar)
			typeCount[charTypeIndex.Int64()]++
		}
	}

	if !r.IsValidPassword(password) {
		return r.GeneratePassword(length) // Regenerate if criteria not met
	}

	return password
}

func canAddChar(password string, char rune, maxRepeats int) bool {
	count := 0
	for _, c := range password {
		if c == char {
			count++
			if count >= maxRepeats {
				return false
			}
		}
	}
	return true
}

/*
  check:
  - password length at least 8 characters
  - a-z at least 1 character
  - A-Z at least 1 character
  - 0-9 at least 1 character
  - special characters  (@#$%*&) at least 1 character
  - cannot duplicate character more than 3 charactor such as aaaa1111
*/
func (r *AdminPassword) IsValidPassword(password string) bool {
	isUpperCase := regexp.MustCompile(`[A-Z]`)
	isLowerCase := regexp.MustCompile(`[a-z]`)
	isSpecialCase := regexp.MustCompile(`[@#$%*&()]`)
	isNumber := regexp.MustCompile(`\d`)
	if len(password) < 8 {
		return false
	}
	if r.isDuplicateCharactor(password, 3) {
		return false
	}
	if !isUpperCase.MatchString(password) {
		return false
	}
	if !isLowerCase.MatchString(password) {
		return false
	}
	if !isNumber.MatchString(password) {
		return false
	}
	if !isSpecialCase.MatchString(password) {
		return false
	}
	return true
}

func (r *AdminPassword) isDuplicateCharactor(password string, times int) bool {
	last := rune(' ')
	i := 0

	for _, c := range password {
		if c == last {
			i++
			if i >= times {
				return true
			}
		} else {
			last = c
			i = 0
		}
	}

	return false
}
