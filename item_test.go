package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testItems = Items{
	Item{
		Name:        "foo",
		Description: "This is test password",
		Password:    "foopassword",
	},
	Item{
		Name:        "bar",
		Description: "This is another password",
		Password:    "barpassword",
	},
}

func TestFindOnNameMatched(t *testing.T) {
	i := testItems.Find("foo")

	if i == nil {
		t.Error("item not found")
	}

	if i.Description != "This is test password" {
		t.Errorf("unexpected item found: %+v", i)
	}
	// key, _ := GetKey(filepath.Join("testdata", "key_file"))
	// testItems.Save(key, "testdata/data.dat")
}

func TestFindOnNameUnmatched(t *testing.T) {
	i := testItems.Find("qux")

	if i != nil {
		t.Error("find with unmatched name should return nil")
	}
}

func TestSave(t *testing.T) {
	td, _ := ioutil.TempDir("", "")
	defer func() {
		os.Remove(td)
	}()

	key, err := GetKey(filepath.Join("testdata", "key_file"))
	if err != nil {
		t.Error(err)
	}
	path := filepath.Join(td, "data.dat")
	testItems.Save(key, path)
	if _, err := os.Stat(path); err != nil {
		t.Error(err)
	}
}

func TestLoadItems(t *testing.T) {
	key, err := GetKey(filepath.Join("testdata", "key_file"))
	if err != nil {
		t.Error(err)
	}
	is, err := LoadItems(key, filepath.Join("testdata", "data.dat"))
	if err != nil {
		t.Error(err)
	}
	i := is.Find("foo")

	if i.Description != "This is test password" {
		t.Error("LoadItems is failure")
	}
}

func TestLoadItemsWithNotExistFile(t *testing.T) {
	key, err := GetKey(filepath.Join("testdata", "key_file"))
	if err != nil {
		t.Error(err)
	}
	is, err := LoadItems(key, filepath.Join("testdata", "not_exist.dat"))
	if err != nil {
		t.Error(err)
	}
	if len(is) != 0 {
		t.Error("LoadItems should return empty items when file does not exist")
	}
}
