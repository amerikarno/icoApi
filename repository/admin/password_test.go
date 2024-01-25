package adminLoginRepository_test

import (
	"testing"

	adminLoginRepository "github.com/amerikarno/icoApi/repository/admin"
)

type passwordTest struct {
	password string
	expected bool
}

func TestIsValidate(t *testing.T) {
	passwordTestLists := []passwordTest{
		{"aA0$", false},
		{"", false},
		{"aaaa1111", false},
		{"BBAA1122", false},
		{"BBAAaab&", false},
		{"aaAA1122", false},
		{"aA0$aA0$", true},
	}
	p := adminLoginRepository.NewAdminPassword()
	for i := range passwordTestLists {
		test := &passwordTestLists[i]
		actual := p.IsValidPassword(test.password)
		if actual != test.expected {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}

	}
}
