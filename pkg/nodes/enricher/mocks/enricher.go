// Code generated by MockGen. DO NOT EDIT.
// Source: enricher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	types "github.com/stackrox/rox/pkg/scanners/types"
)

// MockNodeEnricher is a mock of NodeEnricher interface.
type MockNodeEnricher struct {
	ctrl     *gomock.Controller
	recorder *MockNodeEnricherMockRecorder
}

// MockNodeEnricherMockRecorder is the mock recorder for MockNodeEnricher.
type MockNodeEnricherMockRecorder struct {
	mock *MockNodeEnricher
}

// NewMockNodeEnricher creates a new mock instance.
func NewMockNodeEnricher(ctrl *gomock.Controller) *MockNodeEnricher {
	mock := &MockNodeEnricher{ctrl: ctrl}
	mock.recorder = &MockNodeEnricherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeEnricher) EXPECT() *MockNodeEnricherMockRecorder {
	return m.recorder
}

// CreateNodeScanner mocks base method.
func (m *MockNodeEnricher) CreateNodeScanner(integration *storage.NodeIntegration) (types.NodeScannerWithDataSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNodeScanner", integration)
	ret0, _ := ret[0].(types.NodeScannerWithDataSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNodeScanner indicates an expected call of CreateNodeScanner.
func (mr *MockNodeEnricherMockRecorder) CreateNodeScanner(integration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNodeScanner", reflect.TypeOf((*MockNodeEnricher)(nil).CreateNodeScanner), integration)
}

// EnrichNode mocks base method.
func (m *MockNodeEnricher) EnrichNode(node *storage.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrichNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnrichNode indicates an expected call of EnrichNode.
func (mr *MockNodeEnricherMockRecorder) EnrichNode(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichNode", reflect.TypeOf((*MockNodeEnricher)(nil).EnrichNode), node)
}

// RemoveNodeIntegration mocks base method.
func (m *MockNodeEnricher) RemoveNodeIntegration(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveNodeIntegration", id)
}

// RemoveNodeIntegration indicates an expected call of RemoveNodeIntegration.
func (mr *MockNodeEnricherMockRecorder) RemoveNodeIntegration(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNodeIntegration", reflect.TypeOf((*MockNodeEnricher)(nil).RemoveNodeIntegration), id)
}

// UpsertNodeIntegration mocks base method.
func (m *MockNodeEnricher) UpsertNodeIntegration(integration *storage.NodeIntegration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNodeIntegration", integration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNodeIntegration indicates an expected call of UpsertNodeIntegration.
func (mr *MockNodeEnricherMockRecorder) UpsertNodeIntegration(integration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNodeIntegration", reflect.TypeOf((*MockNodeEnricher)(nil).UpsertNodeIntegration), integration)
}

// MockcveSuppressor is a mock of cveSuppressor interface.
type MockcveSuppressor struct {
	ctrl     *gomock.Controller
	recorder *MockcveSuppressorMockRecorder
}

// MockcveSuppressorMockRecorder is the mock recorder for MockcveSuppressor.
type MockcveSuppressorMockRecorder struct {
	mock *MockcveSuppressor
}

// NewMockcveSuppressor creates a new mock instance.
func NewMockcveSuppressor(ctrl *gomock.Controller) *MockcveSuppressor {
	mock := &MockcveSuppressor{ctrl: ctrl}
	mock.recorder = &MockcveSuppressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcveSuppressor) EXPECT() *MockcveSuppressorMockRecorder {
	return m.recorder
}

// EnrichNodeWithSuppressedCVEs mocks base method.
func (m *MockcveSuppressor) EnrichNodeWithSuppressedCVEs(node *storage.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnrichNodeWithSuppressedCVEs", node)
}

// EnrichNodeWithSuppressedCVEs indicates an expected call of EnrichNodeWithSuppressedCVEs.
func (mr *MockcveSuppressorMockRecorder) EnrichNodeWithSuppressedCVEs(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichNodeWithSuppressedCVEs", reflect.TypeOf((*MockcveSuppressor)(nil).EnrichNodeWithSuppressedCVEs), node)
}
