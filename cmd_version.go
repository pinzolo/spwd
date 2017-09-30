package main

import (
	"fmt"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "Print spwd version",
	Long:      `Print spwd version`,
}

// Version of spwd
const Version = "v1.1.0"

func runVersion(ctx context, args []string) error {
	fmt.Fprintln(ctx.out, fmt.Sprintf("spwd %s", Version))
	return nil
}
