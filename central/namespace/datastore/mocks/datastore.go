// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/namespace/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
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

// AddNamespace mocks base method
func (m *MockDataStore) AddNamespace(arg0 *storage.NamespaceMetadata) error {
	ret := m.ctrl.Call(m, "AddNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNamespace indicates an expected call of AddNamespace
func (mr *MockDataStoreMockRecorder) AddNamespace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNamespace", reflect.TypeOf((*MockDataStore)(nil).AddNamespace), arg0)
}

// GetNamespace mocks base method
func (m *MockDataStore) GetNamespace(arg0 string) (*storage.NamespaceMetadata, bool, error) {
	ret := m.ctrl.Call(m, "GetNamespace", arg0)
	ret0, _ := ret[0].(*storage.NamespaceMetadata)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockDataStoreMockRecorder) GetNamespace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockDataStore)(nil).GetNamespace), arg0)
}

// GetNamespaces mocks base method
func (m *MockDataStore) GetNamespaces() ([]*storage.NamespaceMetadata, error) {
	ret := m.ctrl.Call(m, "GetNamespaces")
	ret0, _ := ret[0].([]*storage.NamespaceMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaces indicates an expected call of GetNamespaces
func (mr *MockDataStoreMockRecorder) GetNamespaces() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaces", reflect.TypeOf((*MockDataStore)(nil).GetNamespaces))
}

// RemoveNamespace mocks base method
func (m *MockDataStore) RemoveNamespace(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNamespace indicates an expected call of RemoveNamespace
func (mr *MockDataStoreMockRecorder) RemoveNamespace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNamespace", reflect.TypeOf((*MockDataStore)(nil).RemoveNamespace), arg0)
}

// Search mocks base method
func (m *MockDataStore) Search(arg0 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), arg0)
}

// SearchNamespaces mocks base method
func (m *MockDataStore) SearchNamespaces(arg0 *v1.Query) ([]*storage.NamespaceMetadata, error) {
	ret := m.ctrl.Call(m, "SearchNamespaces", arg0)
	ret0, _ := ret[0].([]*storage.NamespaceMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNamespaces indicates an expected call of SearchNamespaces
func (mr *MockDataStoreMockRecorder) SearchNamespaces(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNamespaces", reflect.TypeOf((*MockDataStore)(nil).SearchNamespaces), arg0)
}

// UpdateNamespace mocks base method
func (m *MockDataStore) UpdateNamespace(arg0 *storage.NamespaceMetadata) error {
	ret := m.ctrl.Call(m, "UpdateNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNamespace indicates an expected call of UpdateNamespace
func (mr *MockDataStoreMockRecorder) UpdateNamespace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespace", reflect.TypeOf((*MockDataStore)(nil).UpdateNamespace), arg0)
}
