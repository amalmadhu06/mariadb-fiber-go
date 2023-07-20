// Code generated by MockGen. DO NOT EDIT.
// Source: internal/memory/interfaces/offer.go

// Package mockMemory is a generated GoMock package.
package mockMemory

import (
	context "context"
	reflect "reflect"

	entities "github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockOfferMemory is a mock of OfferMemory interface.
type MockOfferMemory struct {
	ctrl     *gomock.Controller
	recorder *MockOfferMemoryMockRecorder
}

// MockOfferMemoryMockRecorder is the mock recorder for MockOfferMemory.
type MockOfferMemoryMockRecorder struct {
	mock *MockOfferMemory
}

// NewMockOfferMemory creates a new mock instance.
func NewMockOfferMemory(ctrl *gomock.Controller) *MockOfferMemory {
	mock := &MockOfferMemory{ctrl: ctrl}
	mock.recorder = &MockOfferMemoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfferMemory) EXPECT() *MockOfferMemoryMockRecorder {
	return m.recorder
}

// GetOffer mocks base method.
func (m *MockOfferMemory) GetOffer(ctx context.Context, country string) ([]entities.OfferCompany, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffer", ctx, country)
	ret0, _ := ret[0].([]entities.OfferCompany)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffer indicates an expected call of GetOffer.
func (mr *MockOfferMemoryMockRecorder) GetOffer(ctx, country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffer", reflect.TypeOf((*MockOfferMemory)(nil).GetOffer), ctx, country)
}

// StartCacheUpdater mocks base method.
func (m *MockOfferMemory) StartCacheUpdater() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartCacheUpdater")
}

// StartCacheUpdater indicates an expected call of StartCacheUpdater.
func (mr *MockOfferMemoryMockRecorder) StartCacheUpdater() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCacheUpdater", reflect.TypeOf((*MockOfferMemory)(nil).StartCacheUpdater))
}
