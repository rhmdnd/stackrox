// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	compliance "github.com/stackrox/rox/central/compliance"
	types "github.com/stackrox/rox/central/compliance/datastore/types"
	v1 "github.com/stackrox/rox/generated/api/v1"
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

// QueryControlResults mocks base method
func (m *MockDataStore) QueryControlResults(ctx context.Context, query *v1.Query) ([]*storage.ComplianceControlResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryControlResults", ctx, query)
	ret0, _ := ret[0].([]*storage.ComplianceControlResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryControlResults indicates an expected call of QueryControlResults
func (mr *MockDataStoreMockRecorder) QueryControlResults(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryControlResults", reflect.TypeOf((*MockDataStore)(nil).QueryControlResults), ctx, query)
}

// GetLatestRunResults mocks base method
func (m *MockDataStore) GetLatestRunResults(ctx context.Context, clusterID, standardID string, flags types.GetFlags) (types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunResults", ctx, clusterID, standardID, flags)
	ret0, _ := ret[0].(types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunResults indicates an expected call of GetLatestRunResults
func (mr *MockDataStoreMockRecorder) GetLatestRunResults(ctx, clusterID, standardID, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunResults", reflect.TypeOf((*MockDataStore)(nil).GetLatestRunResults), ctx, clusterID, standardID, flags)
}

// GetLatestRunResultsBatch mocks base method
func (m *MockDataStore) GetLatestRunResultsBatch(ctx context.Context, clusterIDs, standardIDs []string, flags types.GetFlags) (map[compliance.ClusterStandardPair]types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunResultsBatch", ctx, clusterIDs, standardIDs, flags)
	ret0, _ := ret[0].(map[compliance.ClusterStandardPair]types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunResultsBatch indicates an expected call of GetLatestRunResultsBatch
func (mr *MockDataStoreMockRecorder) GetLatestRunResultsBatch(ctx, clusterIDs, standardIDs, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunResultsBatch", reflect.TypeOf((*MockDataStore)(nil).GetLatestRunResultsBatch), ctx, clusterIDs, standardIDs, flags)
}

// GetLatestRunResultsFiltered mocks base method
func (m *MockDataStore) GetLatestRunResultsFiltered(ctx context.Context, clusterIDFilter, standardIDFilter func(string) bool, flags types.GetFlags) (map[compliance.ClusterStandardPair]types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunResultsFiltered", ctx, clusterIDFilter, standardIDFilter, flags)
	ret0, _ := ret[0].(map[compliance.ClusterStandardPair]types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunResultsFiltered indicates an expected call of GetLatestRunResultsFiltered
func (mr *MockDataStoreMockRecorder) GetLatestRunResultsFiltered(ctx, clusterIDFilter, standardIDFilter, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunResultsFiltered", reflect.TypeOf((*MockDataStore)(nil).GetLatestRunResultsFiltered), ctx, clusterIDFilter, standardIDFilter, flags)
}

// GetLatestRunMetadataBatch mocks base method
func (m *MockDataStore) GetLatestRunMetadataBatch(ctx context.Context, clusterID string, standardIDs []string) (map[compliance.ClusterStandardPair]types.ComplianceRunsMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunMetadataBatch", ctx, clusterID, standardIDs)
	ret0, _ := ret[0].(map[compliance.ClusterStandardPair]types.ComplianceRunsMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunMetadataBatch indicates an expected call of GetLatestRunMetadataBatch
func (mr *MockDataStoreMockRecorder) GetLatestRunMetadataBatch(ctx, clusterID, standardIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunMetadataBatch", reflect.TypeOf((*MockDataStore)(nil).GetLatestRunMetadataBatch), ctx, clusterID, standardIDs)
}

// IsComplianceRunSuccessfulOnCluster mocks base method
func (m *MockDataStore) IsComplianceRunSuccessfulOnCluster(ctx context.Context, clusterID string, standardIDs []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsComplianceRunSuccessfulOnCluster", ctx, clusterID, standardIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsComplianceRunSuccessfulOnCluster indicates an expected call of IsComplianceRunSuccessfulOnCluster
func (mr *MockDataStoreMockRecorder) IsComplianceRunSuccessfulOnCluster(ctx, clusterID, standardIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsComplianceRunSuccessfulOnCluster", reflect.TypeOf((*MockDataStore)(nil).IsComplianceRunSuccessfulOnCluster), ctx, clusterID, standardIDs)
}

// StoreRunResults mocks base method
func (m *MockDataStore) StoreRunResults(ctx context.Context, results *storage.ComplianceRunResults) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreRunResults", ctx, results)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreRunResults indicates an expected call of StoreRunResults
func (mr *MockDataStoreMockRecorder) StoreRunResults(ctx, results interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreRunResults", reflect.TypeOf((*MockDataStore)(nil).StoreRunResults), ctx, results)
}

// StoreFailure mocks base method
func (m *MockDataStore) StoreFailure(ctx context.Context, metadata *storage.ComplianceRunMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreFailure", ctx, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreFailure indicates an expected call of StoreFailure
func (mr *MockDataStoreMockRecorder) StoreFailure(ctx, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreFailure", reflect.TypeOf((*MockDataStore)(nil).StoreFailure), ctx, metadata)
}
