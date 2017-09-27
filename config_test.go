package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMerge(t *testing.T) {
	a := Config{
		IdentityFile: "aid",
		DataFile:     "adata",
	}
	b := Config{
		IdentityFile: "bid",
		DataFile:     "bdata",
	}
	x := a.Merge(b)
	if x.IdentityFile != "bid" {
		t.Errorf("IdentityFile should be 'bid', but got %s", x.IdentityFile)
	}
	if x.DataFile != "bdata" {
		t.Errorf("DataFile should be 'bdata', but got %s", x.DataFile)
	}
}

func TestMergeNotOverwriteWithEmpty(t *testing.T) {
	a := Config{
		IdentityFile: "aid",
		DataFile:     "adata",
	}
	b := Config{
		IdentityFile: "bid",
	}
	c := Config{
		DataFile: "cdata",
	}
	d := Config{}

	x := a.Merge(b)
	if x.IdentityFile != "bid" {
		t.Errorf("IdentityFile should be 'bid', but got %s", x.IdentityFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
	x = a.Merge(c)
	if x.IdentityFile != "aid" {
		t.Errorf("IdentityFile should not be empty, but got %s", x.IdentityFile)
	}
	if x.DataFile != "cdata" {
		t.Errorf("DataFile should be 'cdata', but got %s", x.DataFile)
	}
	x = a.Merge(d)
	if x.IdentityFile != "aid" {
		t.Errorf("IdentityFile should not be empty, but got %s", x.IdentityFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
}

func TestMergeNotOverwriteSelf(t *testing.T) {
	a := Config{
		IdentityFile: "aid",
		DataFile:     "adata",
	}
	b := Config{
		IdentityFile: "bid",
		DataFile:     "bdata",
	}
	a.Merge(b)
	if a.IdentityFile != "aid" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.IdentityFile)
	}
	if a.DataFile != "adata" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.DataFile)
	}
}

func TestMergeNotOverwriteOther(t *testing.T) {
	a := Config{
		IdentityFile: "aid",
		DataFile:     "adata",
	}
	b := Config{
		IdentityFile: "bid",
		DataFile:     "bdata",
	}
	a.Merge(b)
	if b.IdentityFile != "bid" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.IdentityFile)
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
	if cfg.IdentityFile != filepath.Join(td, ".ssh", "id_rsa") {
		t.Errorf("default identity file path is invalid, got %s", cfg.IdentityFile)
	}
	if cfg.DataFile != filepath.Join(td, "spwd", "data.yml") {
		t.Errorf("default data file path is invalid, got %s", cfg.IdentityFile)
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
	f.WriteString(`identity_file: /tmp/config.yml
data_file: /tmp/data.yml
`)
	cfg, ok := FileConfig()
	if !ok {
		t.Error("FileConfig should return true when config file exists")
	}
	if cfg.IdentityFile != "/tmp/config.yml" {
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
