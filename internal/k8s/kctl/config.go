package kctl

import (
	"io"
)

// Config is a configuration abstraction for kctl
type Config interface {
	Stdin() io.Reader  // standard input
	Stdout() io.Writer // standard output
	Stderr() io.Writer // standard error
	UseK3s() bool
}
