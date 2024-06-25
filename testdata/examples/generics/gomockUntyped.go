// Code generated by WrapGen. DO NOT EDIT.
package generics

import (
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/gomock/generics/otherPackage"

	"go.uber.org/mock/gomock"
)

// MockTestUntyped is a mock of TestUntyped interface.
type MockTestUntyped[T comparable] struct {
	ctrl     *gomock.Controller
	recorder *MockTestUntypedMockRecorder[T]
}

// MockTestUntypedMockRecorder is the mock recorder for MockTestUntyped.
type MockTestUntypedMockRecorder[T comparable] struct {
	mock *MockTestUntyped[T]
}

// NewMockTestUntyped creates a new mock instance.
func NewMockTestUntyped[T comparable](ctrl *gomock.Controller) *MockTestUntyped[T] {
	mock := &MockTestUntyped[T]{ctrl: ctrl}
	mock.recorder = &MockTestUntypedMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestUntyped[T]) EXPECT() *MockTestUntypedMockRecorder[T] {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockTestUntyped[T]) ISGOMOCK() struct{} {
	return struct{}{}
}

// Equal mocks base method.
func (m *MockTestUntyped[T]) Equal(a T, b T) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", a, b)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockTestUntypedMockRecorder[T]) Equal(a, b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTestUntyped[T])(nil).Equal),  a, b)
}


// MockTestParamsUntyped is a mock of TestParamsUntyped interface.
type MockTestParamsUntyped[T comparable] struct {
	ctrl     *gomock.Controller
	recorder *MockTestParamsUntypedMockRecorder[T]
}

// MockTestParamsUntypedMockRecorder is the mock recorder for MockTestParamsUntyped.
type MockTestParamsUntypedMockRecorder[T comparable] struct {
	mock *MockTestParamsUntyped[T]
}

// NewMockTestParamsUntyped creates a new mock instance.
func NewMockTestParamsUntyped[T comparable](ctrl *gomock.Controller) *MockTestParamsUntyped[T] {
	mock := &MockTestParamsUntyped[T]{ctrl: ctrl}
	mock.recorder = &MockTestParamsUntypedMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestParamsUntyped[T]) EXPECT() *MockTestParamsUntypedMockRecorder[T] {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockTestParamsUntyped[T]) ISGOMOCK() struct{} {
	return struct{}{}
}

// Bar mocks base method.
func (m *MockTestParamsUntyped[T]) Bar(a otherPackage.TestStruct[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Bar", a)
	
}

// Bar indicates an expected call of Bar.
func (mr *MockTestParamsUntypedMockRecorder[T]) Bar(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Bar),  a)
}


// Baz mocks base method.
func (m *MockTestParamsUntyped[T]) Baz(a otherPackage.TestStruct[int]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Baz", a)
	
}

// Baz indicates an expected call of Baz.
func (mr *MockTestParamsUntypedMockRecorder[T]) Baz(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Baz", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Baz),  a)
}


// Foo mocks base method.
func (m *MockTestParamsUntyped[T]) Foo(a testInterface1[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo", a)
	
}

// Foo indicates an expected call of Foo.
func (mr *MockTestParamsUntypedMockRecorder[T]) Foo(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Foo),  a)
}


// Quuux mocks base method.
func (m *MockTestParamsUntyped[T]) Quuux() TestType[T, T, T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quuux")
	ret0, _ := ret[0].(TestType[T, T, T])
	return ret0
}

// Quuux indicates an expected call of Quuux.
func (mr *MockTestParamsUntypedMockRecorder[T]) Quuux() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quuux", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Quuux),  )
}


// Quux mocks base method.
func (m *MockTestParamsUntyped[T]) Quux() otherPackage.TestInterface[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quux")
	ret0, _ := ret[0].(otherPackage.TestInterface[T])
	return ret0
}

// Quux indicates an expected call of Quux.
func (mr *MockTestParamsUntypedMockRecorder[T]) Quux() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quux", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Quux),  )
}


// Qux mocks base method.
func (m *MockTestParamsUntyped[T]) Qux(a otherPackage.TestInterface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Qux", a)
	
}

// Qux indicates an expected call of Qux.
func (mr *MockTestParamsUntypedMockRecorder[T]) Qux(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qux", reflect.TypeOf((*MockTestParamsUntyped[T])(nil).Qux),  a)
}
