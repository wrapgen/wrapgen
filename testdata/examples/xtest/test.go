package xtest

type X int

//wrapgen:generate -template gomock -destination test_generate_test.go -package xtest -name BlaTest
type _ interface {
	Foo(v1 X)
}

//wrapgen:generate -template gomock -destination test_generate_x_test.go -package xtest_test -name BlaXTest
type _ interface {
	Foo(v1 X)
}
