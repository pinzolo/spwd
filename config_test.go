package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMerge(t *testing.T) {
	a := Config{
		KeyFile:              "aid",
		DataFile:             "adata",
		FilteringCommand:     "afil",
		UnprotectiveCommands: []string{"new", "migrate"},
	}
	b := Config{
		KeyFile:              "bid",
		DataFile:             "bdata",
		FilteringCommand:     "bfil",
		UnprotectiveCommands: []string{"remove"},
	}
	x := a.Merge(b)
	if x.KeyFile != "bid" {
		t.Errorf("KeyFile should be 'bid', but got %s", x.KeyFile)
	}
	if x.DataFile != "bdata" {
		t.Errorf("DataFile should be 'bdata', but got %s", x.DataFile)
	}
	if x.FilteringCommand != "bfil" {
		t.Errorf("FilteringCommand should be 'bfil', but got %s", x.FilteringCommand)
	}
	if len(x.UnprotectiveCommands) == 1 && x.UnprotectiveCommands[0] != "remove" {
		t.Errorf("UnprotectiveCommands should be overridden, but got %s", x.UnprotectiveCommands)
	}
}

func TestMergeNotOverwriteWithEmpty(t *testing.T) {
	a := Config{
		KeyFile:          "aid",
		DataFile:         "adata",
		FilteringCommand: "afil",
	}
	b := Config{
		KeyFile: "bid",
	}
	c := Config{
		DataFile: "cdata",
	}
	d := Config{
		FilteringCommand: "dfil",
	}
	z := Config{}

	x := a.Merge(b)
	if x.KeyFile != "bid" {
		t.Errorf("KeyFile should be 'bid', but got %s", x.KeyFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
	if x.FilteringCommand != "afil" {
		t.Errorf("FilteringCommand should not be empty, but got %s", x.FilteringCommand)
	}
	x = a.Merge(c)
	if x.KeyFile != "aid" {
		t.Errorf("KeyFile should not be empty, but got %s", x.KeyFile)
	}
	if x.DataFile != "cdata" {
		t.Errorf("DataFile should be 'cdata', but got %s", x.DataFile)
	}
	if x.FilteringCommand != "afil" {
		t.Errorf("FilteringCommand should not be empty, but got %s", x.FilteringCommand)
	}
	x = a.Merge(d)
	if x.KeyFile != "aid" {
		t.Errorf("KeyFile should not be empty, but got %s", x.KeyFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
	if x.FilteringCommand != "dfil" {
		t.Errorf("FilteringCommand should be 'dfil', but got %s", x.FilteringCommand)
	}
	x = a.Merge(z)
	if x.KeyFile != "aid" {
		t.Errorf("KeyFile should not be empty, but got %s", x.KeyFile)
	}
	if x.DataFile != "adata" {
		t.Errorf("DataFile should not be empty, but got %s", x.DataFile)
	}
	if x.FilteringCommand != "afil" {
		t.Errorf("FilteringCommand should not be empty, but got %s", x.FilteringCommand)
	}
}

func TestMergeNotOverwriteSelf(t *testing.T) {
	a := Config{
		KeyFile:          "aid",
		DataFile:         "adata",
		FilteringCommand: "afil",
	}
	b := Config{
		KeyFile:          "bid",
		DataFile:         "bdata",
		FilteringCommand: "bfil",
	}
	a.Merge(b)
	if a.KeyFile != "aid" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.KeyFile)
	}
	if a.DataFile != "adata" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.DataFile)
	}
	if a.FilteringCommand != "afil" {
		t.Errorf("self value should not be overwritten with other, but got %s", a.FilteringCommand)
	}
}

func TestMergeNotOverwriteOther(t *testing.T) {
	a := Config{
		KeyFile:          "aid",
		DataFile:         "adata",
		FilteringCommand: "afil",
	}
	b := Config{
		KeyFile:          "bid",
		DataFile:         "bdata",
		FilteringCommand: "bfil",
	}
	a.Merge(b)
	if b.KeyFile != "bid" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.KeyFile)
	}
	if b.DataFile != "bdata" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.DataFile)
	}
	if b.FilteringCommand != "bfil" {
		t.Errorf("other value should not be overwritten with other, but got %s", b.FilteringCommand)
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
	if cfg.DataFile != filepath.Join(td, "spwd", "data.dat") {
		t.Errorf("default data file path is invalid, got %s", cfg.KeyFile)
	}
	if cfg.FilteringCommand != "peco" {
		t.Errorf("default filtering command is invalid, got %s", cfg.FilteringCommand)
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
data_file: /tmp/data.dat
filtering_command: fzf
`)
	cfg, ok := FileConfig()
	if !ok {
		t.Error("FileConfig should return true when config file exists")
	}
	if cfg.KeyFile != "/tmp/config.yml" {
		t.Error("FileConfig load invalid file.")
	}
	if cfg.DataFile != "/tmp/data.dat" {
		t.Error("FileConfig load invalid file.")
	}
	if cfg.FilteringCommand != "fzf" {
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

func TestIsProtective(t *testing.T) {
	cfg := Config{
		KeyFile:              "aid",
		DataFile:             "adata",
		FilteringCommand:     "afil",
		UnprotectiveCommands: []string{"copy", "search", "new", "migrate"},
	}
	if cfg.IsProtective("new") {
		t.Errorf("new should not be protected")
	}
	if !cfg.IsProtective("remove") {
		t.Errorf("remove should be protected")
	}
	if !cfg.IsProtective("copy") {
		t.Errorf("copy should be always protected")
	}
	if !cfg.IsProtective("search") {
		t.Errorf("search should be always protected")
	}
}
