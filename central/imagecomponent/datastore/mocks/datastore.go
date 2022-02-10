// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// Delete mocks base method.
func (m *MockDataStore) Delete(ctx context.Context, ids ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataStoreMockRecorder) Delete(ctx interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataStore)(nil).Delete), varargs...)
}

// Exists mocks base method.
func (m *MockDataStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDataStoreMockRecorder) Exists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDataStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockDataStore) Get(ctx context.Context, id string) (*storage.ImageComponent, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.ImageComponent)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockDataStoreMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataStore)(nil).Get), ctx, id)
}

// GetBatch mocks base method.
func (m *MockDataStore) GetBatch(ctx context.Context, id []string) ([]*storage.ImageComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBatch", ctx, id)
	ret0, _ := ret[0].([]*storage.ImageComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBatch indicates an expected call of GetBatch.
func (mr *MockDataStoreMockRecorder) GetBatch(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatch", reflect.TypeOf((*MockDataStore)(nil).GetBatch), ctx, id)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchImageComponents mocks base method.
func (m *MockDataStore) SearchImageComponents(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchImageComponents", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchImageComponents indicates an expected call of SearchImageComponents.
func (mr *MockDataStoreMockRecorder) SearchImageComponents(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchImageComponents", reflect.TypeOf((*MockDataStore)(nil).SearchImageComponents), ctx, q)
}

// SearchRawImageComponents mocks base method.
func (m *MockDataStore) SearchRawImageComponents(ctx context.Context, q *v1.Query) ([]*storage.ImageComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawImageComponents", ctx, q)
	ret0, _ := ret[0].([]*storage.ImageComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawImageComponents indicates an expected call of SearchRawImageComponents.
func (mr *MockDataStoreMockRecorder) SearchRawImageComponents(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawImageComponents", reflect.TypeOf((*MockDataStore)(nil).SearchRawImageComponents), ctx, q)
}

// Upsert mocks base method.
func (m *MockDataStore) Upsert(ctx context.Context, imagecomponents ...*storage.ImageComponent) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range imagecomponents {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upsert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDataStoreMockRecorder) Upsert(ctx interface{}, imagecomponents ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, imagecomponents...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDataStore)(nil).Upsert), varargs...)
}
