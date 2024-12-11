package packageB

type B string

//wrapgen:generate -template moq -destination b_gen.go
type TestB interface {
	Foo() B
}
