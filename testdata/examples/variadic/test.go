package variadic

import "net/http"

type B string

//wrapgen:generate -template gomock -destination gomock.go -name VariadicParameters
//wrapgen:generate -template gomock_untyped -destination gomockUntyped.go -name VariadicParametersUntyped
//wrapgen:generate -template prometheus -destination prometheus.go
type VariadicParameters interface {
	Test1(a ...int)
	Test2(b ...B)
	Test3(c ...http.Request)
}
