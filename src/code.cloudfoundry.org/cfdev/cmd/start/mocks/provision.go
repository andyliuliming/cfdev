// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/start (interfaces: Provisioner)

// Package mocks is a generated GoMock package.
package mocks

import (
	provision "code.cloudfoundry.org/cfdev/provision"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvisioner is a mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// DeployBosh mocks base method
func (m *MockProvisioner) DeployBosh() error {
	ret := m.ctrl.Call(m, "DeployBosh")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployBosh indicates an expected call of DeployBosh
func (mr *MockProvisionerMockRecorder) DeployBosh() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployBosh", reflect.TypeOf((*MockProvisioner)(nil).DeployBosh))
}

// DeployCloudFoundry mocks base method
func (m *MockProvisioner) DeployCloudFoundry(arg0 []string) error {
	ret := m.ctrl.Call(m, "DeployCloudFoundry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployCloudFoundry indicates an expected call of DeployCloudFoundry
func (mr *MockProvisionerMockRecorder) DeployCloudFoundry(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployCloudFoundry", reflect.TypeOf((*MockProvisioner)(nil).DeployCloudFoundry), arg0)
}

// DeployServices mocks base method
func (m *MockProvisioner) DeployServices(arg0 provision.UI, arg1 []provision.Service) error {
	ret := m.ctrl.Call(m, "DeployServices", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployServices indicates an expected call of DeployServices
func (mr *MockProvisionerMockRecorder) DeployServices(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployServices", reflect.TypeOf((*MockProvisioner)(nil).DeployServices), arg0, arg1)
}

// GetServices mocks base method
func (m *MockProvisioner) GetServices() ([]provision.Service, string, error) {
	ret := m.ctrl.Call(m, "GetServices")
	ret0, _ := ret[0].([]provision.Service)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServices indicates an expected call of GetServices
func (mr *MockProvisionerMockRecorder) GetServices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockProvisioner)(nil).GetServices))
}

// Ping mocks base method
func (m *MockProvisioner) Ping() error {
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockProvisionerMockRecorder) Ping() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockProvisioner)(nil).Ping))
}

// ReportProgress mocks base method
func (m *MockProvisioner) ReportProgress(arg0 provision.UI, arg1 string) {
	m.ctrl.Call(m, "ReportProgress", arg0, arg1)
}

// ReportProgress indicates an expected call of ReportProgress
func (mr *MockProvisionerMockRecorder) ReportProgress(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportProgress", reflect.TypeOf((*MockProvisioner)(nil).ReportProgress), arg0, arg1)
}
