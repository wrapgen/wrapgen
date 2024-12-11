// Code generated by WrapGen. DO NOT EDIT.
package generics

import (
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/generics/otherPackage"

	"go.uber.org/mock/gomock"
)

// MockTest is a mock of Test interface.
type MockTest[T comparable] struct {
	ctrl     *gomock.Controller
	recorder *MockTestMockRecorder[T]
}

// MockTestMockRecorder is the mock recorder for MockTest.
type MockTestMockRecorder[T comparable] struct {
	mock *MockTest[T]
}

// NewMockTest creates a new mock instance.
func NewMockTest[T comparable](ctrl *gomock.Controller) *MockTest[T] {
	mock := &MockTest[T]{ctrl: ctrl}
	mock.recorder = &MockTestMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTest[T]) EXPECT() *MockTestMockRecorder[T] {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockTest[T]) ISGOMOCK() struct{} {
	return struct{}{}
}

// Equal mocks base method.
func (m *MockTest[T]) Equal(a T, b T) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", a, b)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockTestMockRecorder[T]) Equal(a, b any) *TestEqualCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTest[T])(nil).Equal), a, b)
	return &TestEqualCall[T]{Call: call}
}

// TestEqualCall wrap *gomock.Call
type TestEqualCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestEqualCall[T]) Return(arg0 bool) *TestEqualCall[T] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestEqualCall[T]) Do(f func(T, T) bool) *TestEqualCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestEqualCall[T]) DoAndReturn(f func(T, T) bool) *TestEqualCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockTestParams is a mock of TestParams interface.
type MockTestParams[T comparable] struct {
	ctrl     *gomock.Controller
	recorder *MockTestParamsMockRecorder[T]
}

// MockTestParamsMockRecorder is the mock recorder for MockTestParams.
type MockTestParamsMockRecorder[T comparable] struct {
	mock *MockTestParams[T]
}

// NewMockTestParams creates a new mock instance.
func NewMockTestParams[T comparable](ctrl *gomock.Controller) *MockTestParams[T] {
	mock := &MockTestParams[T]{ctrl: ctrl}
	mock.recorder = &MockTestParamsMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestParams[T]) EXPECT() *MockTestParamsMockRecorder[T] {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockTestParams[T]) ISGOMOCK() struct{} {
	return struct{}{}
}

// Bar mocks base method.
func (m *MockTestParams[T]) Bar(a otherPackage.TestStruct[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Bar", a)
}

// Bar indicates an expected call of Bar.
func (mr *MockTestParamsMockRecorder[T]) Bar(a any) *TestParamsBarCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockTestParams[T])(nil).Bar), a)
	return &TestParamsBarCall[T]{Call: call}
}

// TestParamsBarCall wrap *gomock.Call
type TestParamsBarCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsBarCall[T]) Return() *TestParamsBarCall[T] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsBarCall[T]) Do(f func(otherPackage.TestStruct[T])) *TestParamsBarCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsBarCall[T]) DoAndReturn(f func(otherPackage.TestStruct[T])) *TestParamsBarCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Baz mocks base method.
func (m *MockTestParams[T]) Baz(a otherPackage.TestStruct[int]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Baz", a)
}

// Baz indicates an expected call of Baz.
func (mr *MockTestParamsMockRecorder[T]) Baz(a any) *TestParamsBazCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Baz", reflect.TypeOf((*MockTestParams[T])(nil).Baz), a)
	return &TestParamsBazCall[T]{Call: call}
}

// TestParamsBazCall wrap *gomock.Call
type TestParamsBazCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsBazCall[T]) Return() *TestParamsBazCall[T] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsBazCall[T]) Do(f func(otherPackage.TestStruct[int])) *TestParamsBazCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsBazCall[T]) DoAndReturn(f func(otherPackage.TestStruct[int])) *TestParamsBazCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Foo mocks base method.
func (m *MockTestParams[T]) Foo(a testInterface1[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo", a)
}

// Foo indicates an expected call of Foo.
func (mr *MockTestParamsMockRecorder[T]) Foo(a any) *TestParamsFooCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockTestParams[T])(nil).Foo), a)
	return &TestParamsFooCall[T]{Call: call}
}

// TestParamsFooCall wrap *gomock.Call
type TestParamsFooCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsFooCall[T]) Return() *TestParamsFooCall[T] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsFooCall[T]) Do(f func(testInterface1[T])) *TestParamsFooCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsFooCall[T]) DoAndReturn(f func(testInterface1[T])) *TestParamsFooCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Quuux mocks base method.
func (m *MockTestParams[T]) Quuux() TestType[T, T, T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quuux")
	ret0, _ := ret[0].(TestType[T, T, T])
	return ret0
}

// Quuux indicates an expected call of Quuux.
func (mr *MockTestParamsMockRecorder[T]) Quuux() *TestParamsQuuuxCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quuux", reflect.TypeOf((*MockTestParams[T])(nil).Quuux))
	return &TestParamsQuuuxCall[T]{Call: call}
}

// TestParamsQuuuxCall wrap *gomock.Call
type TestParamsQuuuxCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsQuuuxCall[T]) Return(arg0 TestType[T, T, T]) *TestParamsQuuuxCall[T] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsQuuuxCall[T]) Do(f func() TestType[T, T, T]) *TestParamsQuuuxCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsQuuuxCall[T]) DoAndReturn(f func() TestType[T, T, T]) *TestParamsQuuuxCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Quux mocks base method.
func (m *MockTestParams[T]) Quux() otherPackage.TestInterface[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quux")
	ret0, _ := ret[0].(otherPackage.TestInterface[T])
	return ret0
}

// Quux indicates an expected call of Quux.
func (mr *MockTestParamsMockRecorder[T]) Quux() *TestParamsQuuxCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quux", reflect.TypeOf((*MockTestParams[T])(nil).Quux))
	return &TestParamsQuuxCall[T]{Call: call}
}

// TestParamsQuuxCall wrap *gomock.Call
type TestParamsQuuxCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsQuuxCall[T]) Return(arg0 otherPackage.TestInterface[T]) *TestParamsQuuxCall[T] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsQuuxCall[T]) Do(f func() otherPackage.TestInterface[T]) *TestParamsQuuxCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsQuuxCall[T]) DoAndReturn(f func() otherPackage.TestInterface[T]) *TestParamsQuuxCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Qux mocks base method.
func (m *MockTestParams[T]) Qux(a otherPackage.TestInterface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Qux", a)
}

// Qux indicates an expected call of Qux.
func (mr *MockTestParamsMockRecorder[T]) Qux(a any) *TestParamsQuxCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qux", reflect.TypeOf((*MockTestParams[T])(nil).Qux), a)
	return &TestParamsQuxCall[T]{Call: call}
}

// TestParamsQuxCall wrap *gomock.Call
type TestParamsQuxCall[T comparable] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TestParamsQuxCall[T]) Return() *TestParamsQuxCall[T] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TestParamsQuxCall[T]) Do(f func(otherPackage.TestInterface[T])) *TestParamsQuxCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TestParamsQuxCall[T]) DoAndReturn(f func(otherPackage.TestInterface[T])) *TestParamsQuxCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
