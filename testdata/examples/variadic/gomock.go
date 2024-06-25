// Code generated by WrapGen. DO NOT EDIT.
package variadic

import (
	"net/http"
	"reflect"

	"go.uber.org/mock/gomock"
)

// MockVariadicParameters is a mock of VariadicParameters interface.
type MockVariadicParameters struct {
	ctrl     *gomock.Controller
	recorder *MockVariadicParametersMockRecorder
}

// MockVariadicParametersMockRecorder is the mock recorder for MockVariadicParameters.
type MockVariadicParametersMockRecorder struct {
	mock *MockVariadicParameters
}

// NewMockVariadicParameters creates a new mock instance.
func NewMockVariadicParameters(ctrl *gomock.Controller) *MockVariadicParameters {
	mock := &MockVariadicParameters{ctrl: ctrl}
	mock.recorder = &MockVariadicParametersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariadicParameters) EXPECT() *MockVariadicParametersMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockVariadicParameters) ISGOMOCK() struct{} {
	return struct{}{}
}

// Test1 mocks base method.
func (m *MockVariadicParameters) Test1(a ...int) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range a {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Test1", varargs...)
}

// Test1 indicates an expected call of Test1.
func (mr *MockVariadicParametersMockRecorder) Test1(a ...any) *VariadicParametersTest1Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{}, a...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test1", reflect.TypeOf((*MockVariadicParameters)(nil).Test1), varargs...)
	return &VariadicParametersTest1Call{Call: call}
}

// VariadicParametersTest1Call wrap *gomock.Call
type VariadicParametersTest1Call struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *VariadicParametersTest1Call) Return() *VariadicParametersTest1Call {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *VariadicParametersTest1Call) Do(f func(...int)) *VariadicParametersTest1Call {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *VariadicParametersTest1Call) DoAndReturn(f func(...int)) *VariadicParametersTest1Call {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Test2 mocks base method.
func (m *MockVariadicParameters) Test2(b ...B) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range b {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Test2", varargs...)
}

// Test2 indicates an expected call of Test2.
func (mr *MockVariadicParametersMockRecorder) Test2(b ...any) *VariadicParametersTest2Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{}, b...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test2", reflect.TypeOf((*MockVariadicParameters)(nil).Test2), varargs...)
	return &VariadicParametersTest2Call{Call: call}
}

// VariadicParametersTest2Call wrap *gomock.Call
type VariadicParametersTest2Call struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *VariadicParametersTest2Call) Return() *VariadicParametersTest2Call {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *VariadicParametersTest2Call) Do(f func(...B)) *VariadicParametersTest2Call {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *VariadicParametersTest2Call) DoAndReturn(f func(...B)) *VariadicParametersTest2Call {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Test3 mocks base method.
func (m *MockVariadicParameters) Test3(c ...http.Request) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range c {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Test3", varargs...)
}

// Test3 indicates an expected call of Test3.
func (mr *MockVariadicParametersMockRecorder) Test3(c ...any) *VariadicParametersTest3Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{}, c...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test3", reflect.TypeOf((*MockVariadicParameters)(nil).Test3), varargs...)
	return &VariadicParametersTest3Call{Call: call}
}

// VariadicParametersTest3Call wrap *gomock.Call
type VariadicParametersTest3Call struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *VariadicParametersTest3Call) Return() *VariadicParametersTest3Call {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *VariadicParametersTest3Call) Do(f func(...http.Request)) *VariadicParametersTest3Call {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *VariadicParametersTest3Call) DoAndReturn(f func(...http.Request)) *VariadicParametersTest3Call {
	c.Call = c.Call.DoAndReturn(f)
	return c
}