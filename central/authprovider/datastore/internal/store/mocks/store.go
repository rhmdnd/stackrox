// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAuthProvider mocks base method.
func (m *MockStore) AddAuthProvider(authProvider *storage.AuthProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuthProvider", authProvider)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAuthProvider indicates an expected call of AddAuthProvider.
func (mr *MockStoreMockRecorder) AddAuthProvider(authProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuthProvider", reflect.TypeOf((*MockStore)(nil).AddAuthProvider), authProvider)
}

// GetAllAuthProviders mocks base method.
func (m *MockStore) GetAllAuthProviders() ([]*storage.AuthProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAuthProviders")
	ret0, _ := ret[0].([]*storage.AuthProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAuthProviders indicates an expected call of GetAllAuthProviders.
func (mr *MockStoreMockRecorder) GetAllAuthProviders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAuthProviders", reflect.TypeOf((*MockStore)(nil).GetAllAuthProviders))
}

// RemoveAuthProvider mocks base method.
func (m *MockStore) RemoveAuthProvider(d string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAuthProvider", d)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAuthProvider indicates an expected call of RemoveAuthProvider.
func (mr *MockStoreMockRecorder) RemoveAuthProvider(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAuthProvider", reflect.TypeOf((*MockStore)(nil).RemoveAuthProvider), d)
}

// UpdateAuthProvider mocks base method.
func (m *MockStore) UpdateAuthProvider(authProvider *storage.AuthProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthProvider", authProvider)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthProvider indicates an expected call of UpdateAuthProvider.
func (mr *MockStoreMockRecorder) UpdateAuthProvider(authProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthProvider", reflect.TypeOf((*MockStore)(nil).UpdateAuthProvider), authProvider)
}
