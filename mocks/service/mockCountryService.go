// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luschnat-ziegler/cc_backend_go/service (interfaces: CountryService)

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	dto "github.com/luschnat-ziegler/cc_backend_go/dto"
	errs "github.com/luschnat-ziegler/cc_backend_go/errs"
	reflect "reflect"
)

// MockCountryService is a mock of CountryService interface
type MockCountryService struct {
	ctrl     *gomock.Controller
	recorder *MockCountryServiceMockRecorder
}

// MockCountryServiceMockRecorder is the mock recorder for MockCountryService
type MockCountryServiceMockRecorder struct {
	mock *MockCountryService
}

// NewMockCountryService creates a new mock instance
func NewMockCountryService(ctrl *gomock.Controller) *MockCountryService {
	mock := &MockCountryService{ctrl: ctrl}
	mock.recorder = &MockCountryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCountryService) EXPECT() *MockCountryServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockCountryService) GetAll() ([]dto.GetCountryResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]dto.GetCountryResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockCountryServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCountryService)(nil).GetAll))
}
