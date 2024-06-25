package generics

import "github.com/wrapgen/wrapgen/testdata/examples/gomock/generics/otherPackage"

//wrapgen:generate -template gomock -destination gomock.go -name Test
//wrapgen:generate -template gomock_untyped -destination gomockUntyped.go -name TestUntyped
//wrapgen:generate -template prometheus -destination prometheus.go -name Test123
type testInterface1[T comparable] interface {
	Equal(a, b T) bool
}

type TestType[A, B, C comparable] interface {
	Do(A, B, C)
}

//wrapgen:generate -template gomock -destination gomock.go -name TestParams
//wrapgen:generate -template gomock_untyped -destination gomockUntyped.go -name TestParamsUntyped
//wrapgen:generate -template prometheus -destination prometheus.go
type testInterfaceGenericParameters[T comparable] interface {
	Foo(a testInterface1[T])
	Bar(a otherPackage.TestStruct[T])
	Baz(a otherPackage.TestStruct[int])
	Qux(a otherPackage.TestInterface[T])
	Quux() otherPackage.TestInterface[T]
	Quuux() TestType[T, T, T]
}
