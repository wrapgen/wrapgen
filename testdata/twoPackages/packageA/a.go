package packageA

import "github.com/wrapgen/wrapgen/testdata/twoPackages/packageB"

// TestA embeds TestB but must not trigger source generation on TestB.
//
//wrapgen:generate -template moq -destination a_gen.go
type TestA interface {
	packageB.TestB
}
