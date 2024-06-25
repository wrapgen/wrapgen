package fancyTypes

//wrapgen:generate -template gomock -destination gomock.go
type Test interface {
	TestFunc(f func(a int32) func(error) int)
	TestChan(a <-chan struct{}, b chan<- struct{}, c chan struct{}) error
	TestPointer(a *int32) *struct{}
	TestSlice([][][]map[string]int)
}
