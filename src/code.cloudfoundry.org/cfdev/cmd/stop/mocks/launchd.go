// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/stop (interfaces: Launchd)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLaunchd is a mock of Launchd interface
type MockLaunchd struct {
	ctrl     *gomock.Controller
	recorder *MockLaunchdMockRecorder
}

// MockLaunchdMockRecorder is the mock recorder for MockLaunchd
type MockLaunchdMockRecorder struct {
	mock *MockLaunchd
}

// NewMockLaunchd creates a new mock instance
func NewMockLaunchd(ctrl *gomock.Controller) *MockLaunchd {
	mock := &MockLaunchd{ctrl: ctrl}
	mock.recorder = &MockLaunchdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLaunchd) EXPECT() *MockLaunchdMockRecorder {
	return m.recorder
}

// RemoveDaemon mocks base method
func (m *MockLaunchd) RemoveDaemon(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveDaemon", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDaemon indicates an expected call of RemoveDaemon
func (mr *MockLaunchdMockRecorder) RemoveDaemon(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDaemon", reflect.TypeOf((*MockLaunchd)(nil).RemoveDaemon), arg0)
}
