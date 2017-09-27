package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInitialize(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(td)
	cfg := Config{
		DataFile: filepath.Join(td, "spwd", "data.yml"),
	}
	Initialize(cfg)
	if _, err = os.Stat(filepath.Dir(cfg.DataFile)); err != nil {
		t.Errorf("Initialize should create data dir, but got error: %+v", err)
	}
}
