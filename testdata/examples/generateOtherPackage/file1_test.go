package generateOtherPackage

//wrapgen:generate -template gomock -destination gomock_test.go -package generateOtherPackage_test
//wrapgen:generate -template gomock -destination subpackage/gomock_test.go
type _ interface {
	FromOtherFile
}

//wrapgen:generate -template prometheus -destination prometheus_test.go -package generateOtherPackage_test
//wrapgen:generate -template prometheus -destination subpackage/prometheus_test.go
type PrometheusTest interface {
	FromOtherFile
}
