package xtest_test

//wrapgen:generate -template gomock -destination xtest_generate_test.go -package xtest -name FooTest
type _ interface {
	// a reference to xtest_test package can't work from here.
	Foo()
}

type TestXType int

//wrapgen:generate -template gomock -destination xtest_generate_x_test.go -package xtest_test -name FooXTest
type _ interface {
	Foo(v1 TestXType)
}
