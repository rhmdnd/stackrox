// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/processwhitelist/index (interfaces: Indexer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
	reflect "reflect"
)

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddWhitelist mocks base method
func (m *MockIndexer) AddWhitelist(arg0 *storage.ProcessWhitelist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWhitelist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWhitelist indicates an expected call of AddWhitelist
func (mr *MockIndexerMockRecorder) AddWhitelist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWhitelist", reflect.TypeOf((*MockIndexer)(nil).AddWhitelist), arg0)
}

// AddWhitelists mocks base method
func (m *MockIndexer) AddWhitelists(arg0 []*storage.ProcessWhitelist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWhitelists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWhitelists indicates an expected call of AddWhitelists
func (mr *MockIndexerMockRecorder) AddWhitelists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWhitelists", reflect.TypeOf((*MockIndexer)(nil).AddWhitelists), arg0)
}

// DeleteWhitelist mocks base method
func (m *MockIndexer) DeleteWhitelist(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWhitelist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWhitelist indicates an expected call of DeleteWhitelist
func (mr *MockIndexerMockRecorder) DeleteWhitelist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWhitelist", reflect.TypeOf((*MockIndexer)(nil).DeleteWhitelist), arg0)
}

// DeleteWhitelists mocks base method
func (m *MockIndexer) DeleteWhitelists(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWhitelists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWhitelists indicates an expected call of DeleteWhitelists
func (mr *MockIndexerMockRecorder) DeleteWhitelists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWhitelists", reflect.TypeOf((*MockIndexer)(nil).DeleteWhitelists), arg0)
}

// GetTxnCount mocks base method
func (m *MockIndexer) GetTxnCount() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxnCount")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetTxnCount indicates an expected call of GetTxnCount
func (mr *MockIndexerMockRecorder) GetTxnCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxnCount", reflect.TypeOf((*MockIndexer)(nil).GetTxnCount))
}

// ResetIndex mocks base method
func (m *MockIndexer) ResetIndex() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetIndex")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetIndex indicates an expected call of ResetIndex
func (mr *MockIndexerMockRecorder) ResetIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetIndex", reflect.TypeOf((*MockIndexer)(nil).ResetIndex))
}

// Search mocks base method
func (m *MockIndexer) Search(arg0 *v1.Query, arg1 ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), varargs...)
}

// SetTxnCount mocks base method
func (m *MockIndexer) SetTxnCount(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTxnCount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTxnCount indicates an expected call of SetTxnCount
func (mr *MockIndexerMockRecorder) SetTxnCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTxnCount", reflect.TypeOf((*MockIndexer)(nil).SetTxnCount), arg0)
}
