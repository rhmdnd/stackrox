// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
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

// Search mocks base method
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchSecrets mocks base method
func (m *MockDataStore) SearchSecrets(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSecrets", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSecrets indicates an expected call of SearchSecrets
func (mr *MockDataStoreMockRecorder) SearchSecrets(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSecrets", reflect.TypeOf((*MockDataStore)(nil).SearchSecrets), ctx, q)
}

// SearchRawSecrets mocks base method
func (m *MockDataStore) SearchRawSecrets(ctx context.Context, q *v1.Query) ([]*storage.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawSecrets", ctx, q)
	ret0, _ := ret[0].([]*storage.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawSecrets indicates an expected call of SearchRawSecrets
func (mr *MockDataStoreMockRecorder) SearchRawSecrets(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawSecrets", reflect.TypeOf((*MockDataStore)(nil).SearchRawSecrets), ctx, q)
}

// SearchListSecrets mocks base method
func (m *MockDataStore) SearchListSecrets(ctx context.Context, q *v1.Query) ([]*storage.ListSecret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchListSecrets", ctx, q)
	ret0, _ := ret[0].([]*storage.ListSecret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListSecrets indicates an expected call of SearchListSecrets
func (mr *MockDataStoreMockRecorder) SearchListSecrets(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListSecrets", reflect.TypeOf((*MockDataStore)(nil).SearchListSecrets), ctx, q)
}

// CountSecrets mocks base method
func (m *MockDataStore) CountSecrets(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountSecrets", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountSecrets indicates an expected call of CountSecrets
func (mr *MockDataStoreMockRecorder) CountSecrets(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSecrets", reflect.TypeOf((*MockDataStore)(nil).CountSecrets), ctx)
}

// GetSecret mocks base method
func (m *MockDataStore) GetSecret(ctx context.Context, id string) (*storage.Secret, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, id)
	ret0, _ := ret[0].(*storage.Secret)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockDataStoreMockRecorder) GetSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockDataStore)(nil).GetSecret), ctx, id)
}

// UpsertSecret mocks base method
func (m *MockDataStore) UpsertSecret(ctx context.Context, request *storage.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSecret", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSecret indicates an expected call of UpsertSecret
func (mr *MockDataStoreMockRecorder) UpsertSecret(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSecret", reflect.TypeOf((*MockDataStore)(nil).UpsertSecret), ctx, request)
}

// RemoveSecret mocks base method
func (m *MockDataStore) RemoveSecret(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecret", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret
func (mr *MockDataStoreMockRecorder) RemoveSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecret", reflect.TypeOf((*MockDataStore)(nil).RemoveSecret), ctx, id)
}
