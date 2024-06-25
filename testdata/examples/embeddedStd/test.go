package embeddedStd

import (
	"net/http"
)

//wrapgen:generate -template gomock -destination gomock1.go -name Intf1
//wrapgen:generate -template gomock -destination gomock2.go -name Intf2
type X interface {
	http.ResponseWriter
}

//wrapgen:generate -template prometheus -destination prometheus1.go -name Intf1
//wrapgen:generate -template prometheus -destination prometheus2.go -name Intf2
type testInterface interface {
	X
}
