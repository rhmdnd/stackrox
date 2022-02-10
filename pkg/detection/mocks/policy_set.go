// Code generated by MockGen. DO NOT EDIT.
// Source: policy_set.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	detection "github.com/stackrox/rox/pkg/detection"
)

// MockPolicySet is a mock of PolicySet interface.
type MockPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockPolicySetMockRecorder
}

// MockPolicySetMockRecorder is the mock recorder for MockPolicySet.
type MockPolicySetMockRecorder struct {
	mock *MockPolicySet
}

// NewMockPolicySet creates a new mock instance.
func NewMockPolicySet(ctrl *gomock.Controller) *MockPolicySet {
	mock := &MockPolicySet{ctrl: ctrl}
	mock.recorder = &MockPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicySet) EXPECT() *MockPolicySetMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockPolicySet) Exists(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockPolicySetMockRecorder) Exists(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockPolicySet)(nil).Exists), id)
}

// ForEach mocks base method.
func (m *MockPolicySet) ForEach(arg0 func(detection.CompiledPolicy) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEach", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEach indicates an expected call of ForEach.
func (mr *MockPolicySetMockRecorder) ForEach(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEach", reflect.TypeOf((*MockPolicySet)(nil).ForEach), arg0)
}

// ForOne mocks base method.
func (m *MockPolicySet) ForOne(policyID string, f func(detection.CompiledPolicy) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForOne", policyID, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForOne indicates an expected call of ForOne.
func (mr *MockPolicySetMockRecorder) ForOne(policyID, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForOne", reflect.TypeOf((*MockPolicySet)(nil).ForOne), policyID, f)
}

// GetCompiledPolicies mocks base method.
func (m *MockPolicySet) GetCompiledPolicies() map[string]detection.CompiledPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompiledPolicies")
	ret0, _ := ret[0].(map[string]detection.CompiledPolicy)
	return ret0
}

// GetCompiledPolicies indicates an expected call of GetCompiledPolicies.
func (mr *MockPolicySetMockRecorder) GetCompiledPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompiledPolicies", reflect.TypeOf((*MockPolicySet)(nil).GetCompiledPolicies))
}

// RemovePolicy mocks base method.
func (m *MockPolicySet) RemovePolicy(policyID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemovePolicy", policyID)
}

// RemovePolicy indicates an expected call of RemovePolicy.
func (mr *MockPolicySetMockRecorder) RemovePolicy(policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePolicy", reflect.TypeOf((*MockPolicySet)(nil).RemovePolicy), policyID)
}

// UpsertPolicy mocks base method.
func (m *MockPolicySet) UpsertPolicy(arg0 *storage.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPolicy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPolicy indicates an expected call of UpsertPolicy.
func (mr *MockPolicySetMockRecorder) UpsertPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPolicy", reflect.TypeOf((*MockPolicySet)(nil).UpsertPolicy), arg0)
}
