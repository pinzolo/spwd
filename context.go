package main

import "io"

type context struct {
	outWriter io.Writer
	errWriter io.Writer
}

func newContext(ow io.Writer, ew io.Writer) context {
	return context{
		outWriter: ow,
		errWriter: ew,
	}
}
