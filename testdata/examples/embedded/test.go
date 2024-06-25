package embedded

import "github.com/wrapgen/wrapgen/testdata/examples/gomock/embedded/otherPkg"

type LocalIntf interface {
	Abcdef() float64
}

//wrapgen:generate -template gomock -destination gomock.go
//wrapgen:generate -template prometheus -destination prometheus.go
type intf interface {
	otherPkg.Intf
	LocalIntf
}
