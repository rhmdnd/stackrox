// Code generated by MockGen. DO NOT EDIT.
// Source: clienv.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	common "github.com/stackrox/rox/roxctl/common"
	environment "github.com/stackrox/rox/roxctl/common/environment"
	grpc "google.golang.org/grpc"
)

// MockEnvironment is a mock of Environment interface.
type MockEnvironment struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentMockRecorder
}

// MockEnvironmentMockRecorder is the mock recorder for MockEnvironment.
type MockEnvironmentMockRecorder struct {
	mock *MockEnvironment
}

// NewMockEnvironment creates a new mock instance.
func NewMockEnvironment(ctrl *gomock.Controller) *MockEnvironment {
	mock := &MockEnvironment{ctrl: ctrl}
	mock.recorder = &MockEnvironmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironment) EXPECT() *MockEnvironmentMockRecorder {
	return m.recorder
}

// ColorWriter mocks base method.
func (m *MockEnvironment) ColorWriter() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColorWriter")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// ColorWriter indicates an expected call of ColorWriter.
func (mr *MockEnvironmentMockRecorder) ColorWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColorWriter", reflect.TypeOf((*MockEnvironment)(nil).ColorWriter))
}

// ConnectNames mocks base method.
func (m *MockEnvironment) ConnectNames() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectNames")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ConnectNames indicates an expected call of ConnectNames.
func (mr *MockEnvironmentMockRecorder) ConnectNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectNames", reflect.TypeOf((*MockEnvironment)(nil).ConnectNames))
}

// GRPCConnection mocks base method.
func (m *MockEnvironment) GRPCConnection() (*grpc.ClientConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GRPCConnection")
	ret0, _ := ret[0].(*grpc.ClientConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GRPCConnection indicates an expected call of GRPCConnection.
func (mr *MockEnvironmentMockRecorder) GRPCConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GRPCConnection", reflect.TypeOf((*MockEnvironment)(nil).GRPCConnection))
}

// HTTPClient mocks base method.
func (m *MockEnvironment) HTTPClient(timeout time.Duration) (common.RoxctlHTTPClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient", timeout)
	ret0, _ := ret[0].(common.RoxctlHTTPClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockEnvironmentMockRecorder) HTTPClient(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockEnvironment)(nil).HTTPClient), timeout)
}

// InputOutput mocks base method.
func (m *MockEnvironment) InputOutput() environment.IO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputOutput")
	ret0, _ := ret[0].(environment.IO)
	return ret0
}

// InputOutput indicates an expected call of InputOutput.
func (mr *MockEnvironmentMockRecorder) InputOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputOutput", reflect.TypeOf((*MockEnvironment)(nil).InputOutput))
}

// Logger mocks base method.
func (m *MockEnvironment) Logger() environment.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(environment.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockEnvironmentMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockEnvironment)(nil).Logger))
}
