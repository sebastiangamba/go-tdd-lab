// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repositories/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockLunchTimeRepo is a mock of LunchTimeRepo interface
type MockLunchTimeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockLunchTimeRepoMockRecorder
}

// MockLunchTimeRepoMockRecorder is the mock recorder for MockLunchTimeRepo
type MockLunchTimeRepoMockRecorder struct {
	mock *MockLunchTimeRepo
}

// NewMockLunchTimeRepo creates a new mock instance
func NewMockLunchTimeRepo(ctrl *gomock.Controller) *MockLunchTimeRepo {
	mock := &MockLunchTimeRepo{ctrl: ctrl}
	mock.recorder = &MockLunchTimeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLunchTimeRepo) EXPECT() *MockLunchTimeRepoMockRecorder {
	return m.recorder
}

// GetLunchTime mocks base method
func (m *MockLunchTimeRepo) GetLunchTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLunchTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLunchTime indicates an expected call of GetLunchTime
func (mr *MockLunchTimeRepoMockRecorder) GetLunchTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLunchTime", reflect.TypeOf((*MockLunchTimeRepo)(nil).GetLunchTime))
}
