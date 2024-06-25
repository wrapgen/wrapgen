package constArrayLength

import "math"

const C = 2

//wrapgen:generate -template gomock -destination gomock.go
//wrapgen:generate -template prometheus -destination prometheus.go
type I interface {
	Foo() [C]int
	Bar() [2]int
	Baz() [math.MaxInt8]int
	Qux() [1 + 2]int
	Quux() [(1 + 2)]int
	Corge() [math.MaxInt8 - 120]int
}
