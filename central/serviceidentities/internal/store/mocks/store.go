// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetServiceIdentities mocks base method
func (m *MockStore) GetServiceIdentities() ([]*storage.ServiceIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceIdentities")
	ret0, _ := ret[0].([]*storage.ServiceIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceIdentities indicates an expected call of GetServiceIdentities
func (mr *MockStoreMockRecorder) GetServiceIdentities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceIdentities", reflect.TypeOf((*MockStore)(nil).GetServiceIdentities))
}

// AddServiceIdentity mocks base method
func (m *MockStore) AddServiceIdentity(identity *storage.ServiceIdentity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServiceIdentity", identity)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddServiceIdentity indicates an expected call of AddServiceIdentity
func (mr *MockStoreMockRecorder) AddServiceIdentity(identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServiceIdentity", reflect.TypeOf((*MockStore)(nil).AddServiceIdentity), identity)
}
