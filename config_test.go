package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMerge(t *testing.T) {
	a := Config{
		KeyFile:  "aid",
		DataFile: "adata",
	}
	b := Config{
		KeyFile:  "bid",
		DataFile: "bdata",
	}
	x := a.Merge(b)
	if x.KeyFile != "bid" {
		t.Errorf("KeyFile should be 'bid', but got %s", x.KeyFile)
	}
	if x.DataFile != "bdata" {
		t.Errorf("DataFile should be 'bdata', but got %s", x.DataFile)
	}
}

func TestMergeNotOverwriteWithEmpty(t *testing.T) {
	a := Config{
		KeyFile:  "aid",
		DataFile: "adata",
	}
	b := Config{
		KeyFile: "bid",
	}
	c := Config{
		DataFile: "cdata",
	}
	d := Config{}

	x := a.Merge(b)
	if x.KeyFile != "bid" {
		t.Errorf("KeyFile should be 'bid', but got %s", x.KeyFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
	x = a.Merge(c)
	if x.KeyFile != "aid" {
		t.Errorf("KeyFile should not be empty, but got %s", x.KeyFile)
	}
	if x.DataFile != "cdata" {
		t.Errorf("DataFile should be 'cdata', but got %s", x.DataFile)
	}
	x = a.Merge(d)
	if x.KeyFile != "aid" {
		t.Errorf("KeyFile should not be empty, but got %s", x.KeyFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
}

func TestMergeNotOverwriteSelf(t *testing.T) {
	a := Config{
		KeyFile:  "aid",
		DataFile: "adata",
	}
	b := Config{
		KeyFile:  "bid",
		DataFile: "bdata",
	}
	a.Merge(b)
	if a.KeyFile != "aid" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.KeyFile)
	}
	if a.DataFile != "adata" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.DataFile)
	}
}

func TestMergeNotOverwriteOther(t *testing.T) {
	a := Config{
		KeyFile:  "aid",
		DataFile: "adata",
	}
	b := Config{
		KeyFile:  "bid",
		DataFile: "bdata",
	}
	a.Merge(b)
	if b.KeyFile != "bid" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.KeyFile)
	}
	if b.DataFile != "bdata" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.DataFile)
	}
}

func TestDefaultConfig(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}

	os.Setenv("HOME", td)
	os.Setenv("XDG_DATA_HOME", td)
	cfg, err := DefaultConfig()
	if err != nil {
		t.Error(err)
	}
	if cfg.KeyFile != filepath.Join(td, ".ssh", "id_rsa") {
		t.Errorf("default key file path is invalid, got %s", cfg.KeyFile)
	}
	if cfg.DataFile != filepath.Join(td, "spwd", "data.yml") {
		t.Errorf("default data file path is invalid, got %s", cfg.KeyFile)
	}
}

func TestFileConfigOnFileExist(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}
	os.Setenv("XDG_CONFIG_HOME", td)
	fp := filepath.Join(td, "spwd", "config.yml")
	err = os.Mkdir(filepath.Dir(fp), 0755)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create(fp)
	if err != nil {
		t.Error(err)
	}
	f.WriteString(`key_file: /tmp/config.yml
data_file: /tmp/data.yml
`)
	cfg, ok := FileConfig()
	if !ok {
		t.Error("FileConfig should return true when config file exists")
	}
	if cfg.KeyFile != "/tmp/config.yml" {
		t.Error("FileConfig load invalid file.")
	}
}

func TestFileConfigOnFileNotExist(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}
	os.Setenv("XDG_CONFIG_HOME", td)
	_, ok := FileConfig()
	if ok {
		t.Error("FileConfig should return false when config file exists")
	}
}
