package kctl

import (
	"io"
)

type Config interface {
	Stdin() io.Reader  // standard input
	Stdout() io.Writer // standard output
	Stderr() io.Writer // standard error
}
