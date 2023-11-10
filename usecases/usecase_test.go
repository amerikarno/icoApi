package usecases_test

import (
	"testing"

	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/usecases"
	"github.com/amerikarno/icoApi/usecases/mock"
	"golang.org/x/tools/go/expect"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite

	controller *gomock.Controller
	oa         *mock.MockIOpenAccountsRepository
	ex         *mock.MockIExternal
	uc         *usecases.OpenAccountUsecases
}

func (s *ServiceTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())

	s.oa = mock.NewMockIOpenAccountsRepository(s.controller)
	s.uc = usecases.NewOpenAccountUsecases(s.oa)
}

func TestIntegratedTest(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) Test_VerifyIDCardNumber() {
	inputs := []string{"1234567890121", "3102000378645", "1234567890180"}
	expected := true

	for i := range inputs {
		input := &inputs[i]
		actual := s.uc.VerifyIDCardNumber(*input)
		s.Equal(actual, expected)
	}
}

func (s *ServiceTestSuite) Test_CreateIDCardOpenAccountUsecase() {
	idcard := models.IDCardOpenAccounts{
		BirthDate:      "1 เม.ย. 2521",
		MarriageStatus: "โสด",
		IDCard:         "3102000378645",
		LaserCode:      "ME-1234567890",
	}
	actual, err := s.uc.CreateIDCardOpenAccountUsecase(idcard)

	expected := "account-id"
	s.Equal(expected, actual)
	s.NoError(err)
}
