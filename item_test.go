package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var testItems = Items{
	Item{
		Name:        "foo",
		Description: "1",
		Encrypted:   "AAA",
	},
	Item{
		Name:        "bar",
		Description: "2",
		Encrypted:   "BBB",
	},
}

func TestFindOnNameMatched(t *testing.T) {
	i := testItems.Find("foo")

	if i == nil {
		t.Error("item not found")
	}

	if i.Description != "1" {
		t.Errorf("unexpected item found: %+v", i)
	}
}

func TestFindOnNameUnmatched(t *testing.T) {
	i := testItems.Find("qux")

	if i != nil {
		t.Error("find with unmatched name should return nil")
	}
}

func TestSave(t *testing.T) {
	tf, _ := ioutil.TempFile("", "")
	defer func() {
		tf.Close()
		os.Remove(tf.Name())
	}()

	testItems.Save(tf.Name())
	p, err := ioutil.ReadAll(tf)
	if err != nil {
		t.Error(err)
	}
	yml := `- name: foo
  description: "1"
  encrypted: AAA
- name: bar
  description: "2"
  encrypted: BBB
`
	if string(p) != yml {
		t.Errorf("saved yaml text is invalid: %s", string(p))
	}
}

func TestLoadItems(t *testing.T) {
	tf, _ := ioutil.TempFile("", "")
	defer func() {
		tf.Close()
		os.Remove(tf.Name())
	}()
	tf.WriteString(`- name: foo
  description: "1"
  encrypted: AAA
- name: bar
  description: "2"
  encrypted: BBB
`)

	is, err := LoadItems(tf.Name())
	if err != nil {
		t.Error(err)
	}
	i := is.Find("foo")

	if i.Description != "1" {
		t.Error("LoadItems is failure")
	}
}

func TestLoadItemsWithNotExistFile(t *testing.T) {
	tf, _ := ioutil.TempFile("", "")
	tf.Close()
	os.Remove(tf.Name())

	is, err := LoadItems(tf.Name())
	if err != nil {
		t.Error(err)
	}
	if len(is) != 0 {
		t.Error("LoadItems should return empty items when file does not exist")
	}
}
