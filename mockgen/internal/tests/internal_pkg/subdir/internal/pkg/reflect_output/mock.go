// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bitgaming/mock/mockgen/internal/tests/internal_pkg/subdir/internal/pkg (interfaces: Intf)

// Package mock_pkg is a generated GoMock package.
package mock_pkg

import (
	reflect "reflect"

	gomock "github.com/bitgaming/mock/gomock"
	pkg "github.com/bitgaming/mock/mockgen/internal/tests/internal_pkg/subdir/internal/pkg"
)

// MockIntf is a mock of Intf interface.
type MockIntf struct {
	ctrl     *gomock.Controller
	recorder *MockIntfMockRecorder
}

// MockIntfMockRecorder is the mock recorder for MockIntf.
type MockIntfMockRecorder struct {
	mock *MockIntf
}

// NewMockIntf creates a new mock instance.
func NewMockIntf(ctrl *gomock.Controller) *MockIntf {
	mock := &MockIntf{ctrl: ctrl}
	mock.recorder = &MockIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntf) EXPECT() *MockIntfMockRecorder {
	return m.recorder
}

// F mocks base method.
func (m *MockIntf) F() pkg.Arg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "F")
	ret0, _ := ret[0].(pkg.Arg)
	return ret0
}

// F indicates an expected call of F.
func (mr *MockIntfMockRecorder) F() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "F", reflect.TypeOf((*MockIntf)(nil).F))
}
