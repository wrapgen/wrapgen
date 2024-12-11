package subp_test

import "github.com/wrapgen/wrapgen/testdata/examples/xtest"

//wrapgen:generate -template gomock -destination xtest_generate_test.go -package subp -name FooTest
type _ interface {
	Foo() xtest.X
}

//wrapgen:generate -template gomock -destination xtest_generate_x_test.go -package subp_test -name FooTest
type _ interface {
	Foo() xtest.X
}
