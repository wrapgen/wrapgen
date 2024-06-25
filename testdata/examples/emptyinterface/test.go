package emptyinterface

//wrapgen:generate -template gomock -destination gomock.go -name EmptyInterface
//wrapgen:generate -template gomock_untyped -destination gomockUntyped.go -name EmptyInterfaceUntyped
//wrapgen:generate -template prometheus -destination prometheus.go
type EmptyInterface interface {
}
