// Code generated by WrapGen. DO NOT EDIT.
package subpackage

import (
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/generateOtherPackage"

	"go.uber.org/mock/gomock"
)

// Mock_ is a mock of _ interface.
type Mock_ struct {
	ctrl     *gomock.Controller
	recorder *Mock_MockRecorder
}

// Mock_MockRecorder is the mock recorder for Mock_.
type Mock_MockRecorder struct {
	mock *Mock_
}

// NewMock_ creates a new mock instance.
func NewMock_(ctrl *gomock.Controller) *Mock_ {
	mock := &Mock_{ctrl: ctrl}
	mock.recorder = &Mock_MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mock_) EXPECT() *Mock_MockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *Mock_) ISGOMOCK() struct{} {
	return struct{}{}
}

// Foo mocks base method.
func (m *Mock_) Foo(someVar generateOtherPackage.SomeType) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Foo", someVar)
	ret0, _ := ret[0].(int)
	return ret0
}

// Foo indicates an expected call of Foo.
func (mr *Mock_MockRecorder) Foo(someVar any) *_FooCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*Mock_)(nil).Foo), someVar)
	return &_FooCall{Call: call}
}

// _FooCall wrap *gomock.Call
type _FooCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *_FooCall) Return(arg0 int) *_FooCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *_FooCall) Do(f func(generateOtherPackage.SomeType) int) *_FooCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *_FooCall) DoAndReturn(f func(generateOtherPackage.SomeType) int) *_FooCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
