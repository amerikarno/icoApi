package usecases_test

import (
	"testing"

	"github.com/amerikarno/icoApi/usecases"
)

func Test_VerifyIDCardNumber(t *testing.T) {
	inputs := []string{"1234567890121", "3102000378645", "1234567890180"}

	uc := usecases.NewUsecases()

	for i := range inputs {
		input := &inputs[i]
		uc.VerifyIDCardNumber(*input)
	}
}
