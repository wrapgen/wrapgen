//go:build tools
// +build tools

package examples

import (
	// This is just to pin the dependency to gomock.
	_ "go.uber.org/mock/mockgen"
)
