// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GetServiceIdentities mocks base method
func (m *MockDataStore) GetServiceIdentities(arg0 context.Context) ([]*storage.ServiceIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceIdentities", arg0)
	ret0, _ := ret[0].([]*storage.ServiceIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceIdentities indicates an expected call of GetServiceIdentities
func (mr *MockDataStoreMockRecorder) GetServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceIdentities", reflect.TypeOf((*MockDataStore)(nil).GetServiceIdentities), arg0)
}

// AddServiceIdentity mocks base method
func (m *MockDataStore) AddServiceIdentity(ctx context.Context, identity *storage.ServiceIdentity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServiceIdentity", ctx, identity)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddServiceIdentity indicates an expected call of AddServiceIdentity
func (mr *MockDataStoreMockRecorder) AddServiceIdentity(ctx, identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServiceIdentity", reflect.TypeOf((*MockDataStore)(nil).AddServiceIdentity), ctx, identity)
}
