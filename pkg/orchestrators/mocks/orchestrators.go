// Code generated by MockGen. DO NOT EDIT.
// Source: orchestrators.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	orchestrators "github.com/stackrox/rox/pkg/orchestrators"
	reflect "reflect"
	time "time"
)

// MockOrchestrator is a mock of Orchestrator interface
type MockOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *MockOrchestratorMockRecorder
}

// MockOrchestratorMockRecorder is the mock recorder for MockOrchestrator
type MockOrchestratorMockRecorder struct {
	mock *MockOrchestrator
}

// NewMockOrchestrator creates a new mock instance
func NewMockOrchestrator(ctrl *gomock.Controller) *MockOrchestrator {
	mock := &MockOrchestrator{ctrl: ctrl}
	mock.recorder = &MockOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrchestrator) EXPECT() *MockOrchestratorMockRecorder {
	return m.recorder
}

// LaunchDaemonSet mocks base method
func (m *MockOrchestrator) LaunchDaemonSet(service orchestrators.SystemService) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchDaemonSet", service)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LaunchDaemonSet indicates an expected call of LaunchDaemonSet
func (mr *MockOrchestratorMockRecorder) LaunchDaemonSet(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchDaemonSet", reflect.TypeOf((*MockOrchestrator)(nil).LaunchDaemonSet), service)
}

// Kill mocks base method
func (m *MockOrchestrator) Kill(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kill", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kill indicates an expected call of Kill
func (mr *MockOrchestratorMockRecorder) Kill(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockOrchestrator)(nil).Kill), id)
}

// WaitForCompletion mocks base method
func (m *MockOrchestrator) WaitForCompletion(service string, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCompletion", service, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForCompletion indicates an expected call of WaitForCompletion
func (mr *MockOrchestratorMockRecorder) WaitForCompletion(service, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCompletion", reflect.TypeOf((*MockOrchestrator)(nil).WaitForCompletion), service, timeout)
}

// CleanUp mocks base method
func (m *MockOrchestrator) CleanUp(ownedByThisInstance bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUp", ownedByThisInstance)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUp indicates an expected call of CleanUp
func (mr *MockOrchestratorMockRecorder) CleanUp(ownedByThisInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockOrchestrator)(nil).CleanUp), ownedByThisInstance)
}
