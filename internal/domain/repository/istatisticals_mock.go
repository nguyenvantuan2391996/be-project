// Code generated by MockGen. DO NOT EDIT.
// Source: istatisticals.go

// Package repository is a generated GoMock package.
package repository

import (
	model "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStatisticalRepositoryInterface is a mock of IStatisticalRepositoryInterface interface.
type MockIStatisticalRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIStatisticalRepositoryInterfaceMockRecorder
}

// MockIStatisticalRepositoryInterfaceMockRecorder is the mock recorder for MockIStatisticalRepositoryInterface.
type MockIStatisticalRepositoryInterfaceMockRecorder struct {
	mock *MockIStatisticalRepositoryInterface
}

// NewMockIStatisticalRepositoryInterface creates a new mock instance.
func NewMockIStatisticalRepositoryInterface(ctrl *gomock.Controller) *MockIStatisticalRepositoryInterface {
	mock := &MockIStatisticalRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockIStatisticalRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStatisticalRepositoryInterface) EXPECT() *MockIStatisticalRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIStatisticalRepositoryInterface) Create(record *model.Statistical) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIStatisticalRepositoryInterfaceMockRecorder) Create(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIStatisticalRepositoryInterface)(nil).Create), record)
}

// GetByQueries mocks base method.
func (m *MockIStatisticalRepositoryInterface) GetByQueries(queries map[string]interface{}) (*model.Statistical, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQueries", queries)
	ret0, _ := ret[0].(*model.Statistical)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQueries indicates an expected call of GetByQueries.
func (mr *MockIStatisticalRepositoryInterfaceMockRecorder) GetByQueries(queries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQueries", reflect.TypeOf((*MockIStatisticalRepositoryInterface)(nil).GetByQueries), queries)
}

// UpdateWithMap mocks base method.
func (m *MockIStatisticalRepositoryInterface) UpdateWithMap(record *model.Statistical, queries map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithMap", record, queries)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithMap indicates an expected call of UpdateWithMap.
func (mr *MockIStatisticalRepositoryInterfaceMockRecorder) UpdateWithMap(record, queries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithMap", reflect.TypeOf((*MockIStatisticalRepositoryInterface)(nil).UpdateWithMap), record, queries)
}