// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/kauthproxy/pkg/adaptors/env (interfaces: Interface)

// Package mock_env is a generated GoMock package.
package mock_env

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AllocateLocalPort mocks base method
func (m *MockInterface) AllocateLocalPort() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateLocalPort")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocateLocalPort indicates an expected call of AllocateLocalPort
func (mr *MockInterfaceMockRecorder) AllocateLocalPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateLocalPort", reflect.TypeOf((*MockInterface)(nil).AllocateLocalPort))
}

// OpenBrowser mocks base method
func (m *MockInterface) OpenBrowser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenBrowser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenBrowser indicates an expected call of OpenBrowser
func (mr *MockInterfaceMockRecorder) OpenBrowser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenBrowser", reflect.TypeOf((*MockInterface)(nil).OpenBrowser), arg0)
}
