// Code generated by WrapGen. DO NOT EDIT.
package xtest_test

import (
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/xtest"

	"go.uber.org/mock/gomock"
)

// MockBlaXTest is a mock of BlaXTest interface.
type MockBlaXTest struct {
	ctrl     *gomock.Controller
	recorder *MockBlaXTestMockRecorder
}

// MockBlaXTestMockRecorder is the mock recorder for MockBlaXTest.
type MockBlaXTestMockRecorder struct {
	mock *MockBlaXTest
}

// NewMockBlaXTest creates a new mock instance.
func NewMockBlaXTest(ctrl *gomock.Controller) *MockBlaXTest {
	mock := &MockBlaXTest{ctrl: ctrl}
	mock.recorder = &MockBlaXTestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlaXTest) EXPECT() *MockBlaXTestMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockBlaXTest) ISGOMOCK() struct{} {
	return struct{}{}
}

// Foo mocks base method.
func (m *MockBlaXTest) Foo(v1 xtest.X) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo", v1)
}

// Foo indicates an expected call of Foo.
func (mr *MockBlaXTestMockRecorder) Foo(v1 any) *BlaXTestFooCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockBlaXTest)(nil).Foo), v1)
	return &BlaXTestFooCall{Call: call}
}

// BlaXTestFooCall wrap *gomock.Call
type BlaXTestFooCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BlaXTestFooCall) Return() *BlaXTestFooCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BlaXTestFooCall) Do(f func(xtest.X)) *BlaXTestFooCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BlaXTestFooCall) DoAndReturn(f func(xtest.X)) *BlaXTestFooCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
