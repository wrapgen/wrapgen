// Code generated by WrapGen. DO NOT EDIT.
package emptyinterface

import (

	"go.uber.org/mock/gomock"
)

// MockEmptyInterface is a mock of EmptyInterface interface.
type MockEmptyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEmptyInterfaceMockRecorder
}

// MockEmptyInterfaceMockRecorder is the mock recorder for MockEmptyInterface.
type MockEmptyInterfaceMockRecorder struct {
	mock *MockEmptyInterface
}

// NewMockEmptyInterface creates a new mock instance.
func NewMockEmptyInterface(ctrl *gomock.Controller) *MockEmptyInterface {
	mock := &MockEmptyInterface{ctrl: ctrl}
	mock.recorder = &MockEmptyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmptyInterface) EXPECT() *MockEmptyInterfaceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockEmptyInterface) ISGOMOCK() struct{} {
	return struct{}{}
}
