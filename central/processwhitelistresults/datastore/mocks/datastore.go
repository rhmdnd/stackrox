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

// UpsertWhitelistResults mocks base method
func (m *MockDataStore) UpsertWhitelistResults(ctx context.Context, results *storage.ProcessWhitelistResults) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWhitelistResults", ctx, results)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWhitelistResults indicates an expected call of UpsertWhitelistResults
func (mr *MockDataStoreMockRecorder) UpsertWhitelistResults(ctx, results interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWhitelistResults", reflect.TypeOf((*MockDataStore)(nil).UpsertWhitelistResults), ctx, results)
}

// GetWhitelistResults mocks base method
func (m *MockDataStore) GetWhitelistResults(ctx context.Context, deploymentID string) (*storage.ProcessWhitelistResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWhitelistResults", ctx, deploymentID)
	ret0, _ := ret[0].(*storage.ProcessWhitelistResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWhitelistResults indicates an expected call of GetWhitelistResults
func (mr *MockDataStoreMockRecorder) GetWhitelistResults(ctx, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWhitelistResults", reflect.TypeOf((*MockDataStore)(nil).GetWhitelistResults), ctx, deploymentID)
}
