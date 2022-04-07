// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/session/usecase/manager.go

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	session "golang-stepik-2022q1/reditclone/pkg/session"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CheckExists mocks base method.
func (m *MockRepo) CheckExists(arg0 context.Context, arg1 session.SessionId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckExists indicates an expected call of CheckExists.
func (mr *MockRepoMockRecorder) CheckExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExists", reflect.TypeOf((*MockRepo)(nil).CheckExists), arg0, arg1)
}

// Set mocks base method.
func (m *MockRepo) Set(arg0 context.Context, arg1 session.SessionId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRepoMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRepo)(nil).Set), arg0, arg1)
}
