package otherPackage

type TestStruct[T comparable] struct {
	Value T
}

type TestInterface[T comparable] interface {
	Test(T)
}
