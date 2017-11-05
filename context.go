package main

import "io"

type context struct {
	out     io.Writer
	cmdName string
}

func newContext(ow io.Writer, cmdName string) context {
	return context{
		out:     ow,
		cmdName: cmdName,
	}
}
