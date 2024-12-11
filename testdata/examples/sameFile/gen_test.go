// Code generated by WrapGen. DO NOT EDIT.
package sameFile_test

import (
	"context"
	"reflect"

	"github.com/wrapgen/wrapgen/testdata/examples/sameFile/apackage"
	"github.com/wrapgen/wrapgen/testdata/examples/sameFile/bpackage"

	"go.uber.org/mock/gomock"
)

// MockAuthAdapter is a mock of AuthAdapter interface.
type MockAuthAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthAdapterMockRecorder
}

// MockAuthAdapterMockRecorder is the mock recorder for MockAuthAdapter.
type MockAuthAdapterMockRecorder struct {
	mock *MockAuthAdapter
}

// NewMockAuthAdapter creates a new mock instance.
func NewMockAuthAdapter(ctrl *gomock.Controller) *MockAuthAdapter {
	mock := &MockAuthAdapter{ctrl: ctrl}
	mock.recorder = &MockAuthAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthAdapter) EXPECT() *MockAuthAdapterMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockAuthAdapter) ISGOMOCK() struct{} {
	return struct{}{}
}

// ATest mocks base method.
func (m *MockAuthAdapter) ATest(ctx context.Context) (apackage.AdminUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ATest", ctx)
	ret0, _ := ret[0].(apackage.AdminUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ATest indicates an expected call of ATest.
func (mr *MockAuthAdapterMockRecorder) ATest(ctx any) *AuthAdapterATestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ATest", reflect.TypeOf((*MockAuthAdapter)(nil).ATest), ctx)
	return &AuthAdapterATestCall{Call: call}
}

// AuthAdapterATestCall wrap *gomock.Call
type AuthAdapterATestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AuthAdapterATestCall) Return(arg0 apackage.AdminUser, arg1 error) *AuthAdapterATestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AuthAdapterATestCall) Do(f func(context.Context) (apackage.AdminUser, error)) *AuthAdapterATestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AuthAdapterATestCall) DoAndReturn(f func(context.Context) (apackage.AdminUser, error)) *AuthAdapterATestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BTest mocks base method.
func (m *MockAuthAdapter) BTest(ctx context.Context) (bpackage.AdminUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTest", ctx)
	ret0, _ := ret[0].(bpackage.AdminUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTest indicates an expected call of BTest.
func (mr *MockAuthAdapterMockRecorder) BTest(ctx any) *AuthAdapterBTestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTest", reflect.TypeOf((*MockAuthAdapter)(nil).BTest), ctx)
	return &AuthAdapterBTestCall{Call: call}
}

// AuthAdapterBTestCall wrap *gomock.Call
type AuthAdapterBTestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AuthAdapterBTestCall) Return(arg0 bpackage.AdminUser, arg1 error) *AuthAdapterBTestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AuthAdapterBTestCall) Do(f func(context.Context) (bpackage.AdminUser, error)) *AuthAdapterBTestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AuthAdapterBTestCall) DoAndReturn(f func(context.Context) (bpackage.AdminUser, error)) *AuthAdapterBTestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
