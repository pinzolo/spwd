package main

import "io"

type context struct {
	out io.Writer
}

func newContext(ow io.Writer) context {
	return context{
		out: ow,
	}
}
