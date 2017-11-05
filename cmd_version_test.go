package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCmdVersion(t *testing.T) {
	out := &bytes.Buffer{}
	ctx := newContext(out, "version")
	err := cmdVersion.Run(ctx, []string{})
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(out.String(), Version) {
		t.Error("version should print version of spwd")
	}
}
