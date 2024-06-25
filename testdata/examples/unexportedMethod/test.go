package unexportedMethod

type unexportedType int

// The first should just work:
// The second should work but creates invalid code since the un-exported symbols are not accessible from another package:
//
//wrapgen:generate -template gomock -destination gomock.go
//wrapgen:generate -template gomock -destination gomock_test.go -package unexportedMethod_test
type example interface {
	unexportedMethod(test unexportedType)
}
