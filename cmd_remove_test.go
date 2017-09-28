package main

import (
	"bytes"
	"testing"
)

func TestCmdRemove(t *testing.T) {
	teardown, err := setupTestData()
	if err != nil {
		if teardown != nil {
			teardown()
		}
	}
	defer teardown()
	out := &bytes.Buffer{}
	ctx := newContext(out)
	err = cmdRemove.Run(ctx, []string{"foo"})
	if err != nil {
		t.Error(err)
	}

	cfg, err := GetConfig()
	if err != nil {
		t.Error(err)
	}
	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		t.Error(err)
	}
	is, err := LoadItems(key, cfg.DataFile)
	if err != nil {
		t.Error(err)
	}

	if is.Find("foo") != nil {
		t.Error("remove failure")
	}
}

func TestCmdRemoveWithUnknownName(t *testing.T) {
	teardown, err := setupTestData()
	if err != nil {
		if teardown != nil {
			teardown()
		}
	}
	defer teardown()
	out := &bytes.Buffer{}
	ctx := newContext(out)
	err = cmdRemove.Run(ctx, []string{"baz"})
	if err == nil {
		t.Error("remove with unknown name should rase error")
	}
}
