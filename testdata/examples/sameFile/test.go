package sameFile

import (
	"github.com/wrapgen/wrapgen/testdata/examples/sameFile/apackage"
	"github.com/wrapgen/wrapgen/testdata/examples/sameFile/bpackage"
)

//wrapgen:generate -template gomock -package sameFile_test -destination gen_test.go -name AuthAdapter
type Adapter interface {
	apackage.Adapter
	bpackage.Adapter
}
