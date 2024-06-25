package otherPkg

type Intf interface {
	Foo(x Bar) error
}

type Bar interface {
	Test() int
}
