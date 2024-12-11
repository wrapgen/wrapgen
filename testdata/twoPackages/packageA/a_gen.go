// Code generated by WrapGen. DO NOT EDIT.
package packageA

import (
	"sync"

	"github.com/wrapgen/wrapgen/testdata/examples/twoPackages/packageB"
)


// TestAMock is a mock implementation of TestA.
type TestAMock struct {
	// FooFunc mocks the Foo method.
	FooFunc func() packageB.B

	// calls tracks calls to the methods.
	calls struct {
		// Foo holds details about calls to the Foo method.
		Foo []struct {
		}
	}
	lockFoo sync.RWMutex
}


// Foo calls FooFunc.
func (mock *TestAMock) Foo() packageB.B {
	if mock.FooFunc == nil {
		panic("TestAMock.FooFunc: method is nil but TestA.Foo was just called")
	}
	callInfo := struct {
	}{
	}
	mock.lockFoo.Lock()
	mock.calls.Foo = append(mock.calls.Foo, callInfo)
	mock.lockFoo.Unlock()
	return mock.FooFunc()
}
