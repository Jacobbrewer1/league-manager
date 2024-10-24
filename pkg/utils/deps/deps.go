//go:build deps

package deps

import (
	_ "github.com/jacobbrewer1/pagefilter"
	_ "github.com/vektra/mockery/v2" // Mockery is a tool for generating mocks for interfaces in Go. This prevents the tool from being removed when running go mod tidy.
)
