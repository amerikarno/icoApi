// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/amerikarno/icoApi/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIOpenAccountsRepository is a mock of IOpenAccountsRepository interface.
type MockIOpenAccountsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIOpenAccountsRepositoryMockRecorder
}

// MockIOpenAccountsRepositoryMockRecorder is the mock recorder for MockIOpenAccountsRepository.
type MockIOpenAccountsRepositoryMockRecorder struct {
	mock *MockIOpenAccountsRepository
}

// NewMockIOpenAccountsRepository creates a new mock instance.
func NewMockIOpenAccountsRepository(ctrl *gomock.Controller) *MockIOpenAccountsRepository {
	mock := &MockIOpenAccountsRepository{ctrl: ctrl}
	mock.recorder = &MockIOpenAccountsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOpenAccountsRepository) EXPECT() *MockIOpenAccountsRepositoryMockRecorder {
	return m.recorder
}

// CheckReisteredCitizenID mocks base method.
func (m *MockIOpenAccountsRepository) CheckReisteredCitizenID(citizenID string) models.CustomerInformations {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReisteredCitizenID", citizenID)
	ret0, _ := ret[0].(models.CustomerInformations)
	return ret0
}

// CheckReisteredCitizenID indicates an expected call of CheckReisteredCitizenID.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CheckReisteredCitizenID(citizenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReisteredCitizenID", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CheckReisteredCitizenID), citizenID)
}

// CheckReisteredEmail mocks base method.
func (m *MockIOpenAccountsRepository) CheckReisteredEmail(email string) models.CustomerInformations {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReisteredEmail", email)
	ret0, _ := ret[0].(models.CustomerInformations)
	return ret0
}

// CheckReisteredEmail indicates an expected call of CheckReisteredEmail.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CheckReisteredEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReisteredEmail", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CheckReisteredEmail), email)
}

// CheckReisteredMobileNo mocks base method.
func (m *MockIOpenAccountsRepository) CheckReisteredMobileNo(mobileno string) models.CustomerInformations {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReisteredMobileNo", mobileno)
	ret0, _ := ret[0].(models.CustomerInformations)
	return ret0
}

// CheckReisteredMobileNo indicates an expected call of CheckReisteredMobileNo.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CheckReisteredMobileNo(mobileno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReisteredMobileNo", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CheckReisteredMobileNo), mobileno)
}

// CreateCustomerConfirms mocks base method.
func (m *MockIOpenAccountsRepository) CreateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomerConfirms", customerConfirms)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomerConfirms indicates an expected call of CreateCustomerConfirms.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CreateCustomerConfirms(customerConfirms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomerConfirms", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CreateCustomerConfirms), customerConfirms)
}

// CreateCustomerExams mocks base method.
func (m *MockIOpenAccountsRepository) CreateCustomerExams(customerExams models.CustomerExamsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomerExams", customerExams)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomerExams indicates an expected call of CreateCustomerExams.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CreateCustomerExams(customerExams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomerExams", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CreateCustomerExams), customerExams)
}

// CreateCustomerInformation mocks base method.
func (m *MockIOpenAccountsRepository) CreateCustomerInformation(columns models.CustomerInformations) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomerInformation", columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomerInformation indicates an expected call of CreateCustomerInformation.
func (mr *MockIOpenAccountsRepositoryMockRecorder) CreateCustomerInformation(columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomerInformation", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).CreateCustomerInformation), columns)
}

// GetHTMLTemplate mocks base method.
func (m *MockIOpenAccountsRepository) GetHTMLTemplate(thaiName, uid, token string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTMLTemplate", thaiName, uid, token)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHTMLTemplate indicates an expected call of GetHTMLTemplate.
func (mr *MockIOpenAccountsRepositoryMockRecorder) GetHTMLTemplate(thaiName, uid, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTMLTemplate", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).GetHTMLTemplate), thaiName, uid, token)
}

// QueryCustomerConfirmsExpireDT mocks base method.
func (m *MockIOpenAccountsRepository) QueryCustomerConfirmsExpireDT(tokenID string) models.CustomerConfirmsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCustomerConfirmsExpireDT", tokenID)
	ret0, _ := ret[0].(models.CustomerConfirmsRequest)
	return ret0
}

// QueryCustomerConfirmsExpireDT indicates an expected call of QueryCustomerConfirmsExpireDT.
func (mr *MockIOpenAccountsRepositoryMockRecorder) QueryCustomerConfirmsExpireDT(tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCustomerConfirmsExpireDT", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).QueryCustomerConfirmsExpireDT), tokenID)
}

// UpdateCustomerConfirms mocks base method.
func (m *MockIOpenAccountsRepository) UpdateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomerConfirms", customerConfirms)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomerConfirms indicates an expected call of UpdateCustomerConfirms.
func (mr *MockIOpenAccountsRepositoryMockRecorder) UpdateCustomerConfirms(customerConfirms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomerConfirms", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).UpdateCustomerConfirms), customerConfirms)
}

// UpdatePersonalInformation mocks base method.
func (m *MockIOpenAccountsRepository) UpdatePersonalInformation(personalInfos models.PersonalInformations, cid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePersonalInformation", personalInfos, cid)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePersonalInformation indicates an expected call of UpdatePersonalInformation.
func (mr *MockIOpenAccountsRepositoryMockRecorder) UpdatePersonalInformation(personalInfos, cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePersonalInformation", reflect.TypeOf((*MockIOpenAccountsRepository)(nil).UpdatePersonalInformation), personalInfos, cid)
}

// MockIExternal is a mock of IExternal interface.
type MockIExternal struct {
	ctrl     *gomock.Controller
	recorder *MockIExternalMockRecorder
}

// MockIExternalMockRecorder is the mock recorder for MockIExternal.
type MockIExternalMockRecorder struct {
	mock *MockIExternal
}

// NewMockIExternal creates a new mock instance.
func NewMockIExternal(ctrl *gomock.Controller) *MockIExternal {
	mock := &MockIExternal{ctrl: ctrl}
	mock.recorder = &MockIExternalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExternal) EXPECT() *MockIExternalMockRecorder {
	return m.recorder
}

// GenUuid mocks base method.
func (m *MockIExternal) GenUuid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenUuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenUuid indicates an expected call of GenUuid.
func (mr *MockIExternalMockRecorder) GenUuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenUuid", reflect.TypeOf((*MockIExternal)(nil).GenUuid))
}
