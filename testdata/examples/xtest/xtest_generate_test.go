// Code generated by WrapGen. DO NOT EDIT.
package xtest

import (
	"reflect"

	"go.uber.org/mock/gomock"
)

// MockFooTest is a mock of FooTest interface.
type MockFooTest struct {
	ctrl     *gomock.Controller
	recorder *MockFooTestMockRecorder
}

// MockFooTestMockRecorder is the mock recorder for MockFooTest.
type MockFooTestMockRecorder struct {
	mock *MockFooTest
}

// NewMockFooTest creates a new mock instance.
func NewMockFooTest(ctrl *gomock.Controller) *MockFooTest {
	mock := &MockFooTest{ctrl: ctrl}
	mock.recorder = &MockFooTestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFooTest) EXPECT() *MockFooTestMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockFooTest) ISGOMOCK() struct{} {
	return struct{}{}
}

// Foo mocks base method.
func (m *MockFooTest) Foo() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo")
}

// Foo indicates an expected call of Foo.
func (mr *MockFooTestMockRecorder) Foo() *FooTestFooCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockFooTest)(nil).Foo))
	return &FooTestFooCall{Call: call}
}

// FooTestFooCall wrap *gomock.Call
type FooTestFooCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *FooTestFooCall) Return() *FooTestFooCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *FooTestFooCall) Do(f func()) *FooTestFooCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *FooTestFooCall) DoAndReturn(f func()) *FooTestFooCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
