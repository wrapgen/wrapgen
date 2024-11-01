// Code generated by WrapGen. DO NOT EDIT.
package subp

import (
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/gomock/xtest"

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
func (m *MockFooTest) Foo() xtest.X {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Foo")
	ret0, _ := ret[0].(xtest.X)
	return ret0
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
func (c *FooTestFooCall) Return(arg0 xtest.X) *FooTestFooCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *FooTestFooCall) Do(f func() xtest.X) *FooTestFooCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *FooTestFooCall) DoAndReturn(f func() xtest.X) *FooTestFooCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
