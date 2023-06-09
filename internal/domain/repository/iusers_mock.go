// Code generated by MockGen. DO NOT EDIT.
// Source: iusers.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	model "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepositoryInterface is a mock of IUserRepositoryInterface interface.
type MockIUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryInterfaceMockRecorder
}

// MockIUserRepositoryInterfaceMockRecorder is the mock recorder for MockIUserRepositoryInterface.
type MockIUserRepositoryInterfaceMockRecorder struct {
	mock *MockIUserRepositoryInterface
}

// NewMockIUserRepositoryInterface creates a new mock instance.
func NewMockIUserRepositoryInterface(ctrl *gomock.Controller) *MockIUserRepositoryInterface {
	mock := &MockIUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepositoryInterface) EXPECT() *MockIUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserRepositoryInterface) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepositoryInterfaceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepositoryInterface)(nil).CreateUser), ctx, user)
}
