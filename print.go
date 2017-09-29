package main

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// PrintSuccess prints out success message to given writer.
func PrintSuccess(w io.Writer, format string, a ...interface{}) {
	color.New(color.FgGreen).Fprintln(w, fmt.Sprintf(format, a...))
}

// PrintError prints out error message to given writer.
func PrintError(w io.Writer, err error) {
	color.New(color.FgRed).Fprintln(w, err)
}
