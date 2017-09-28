package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdMigrate(t *testing.T) {
	teardown, err := setupTestData()
	if err != nil {
		if teardown != nil {
			teardown()
		}
		t.Error(err)
	}
	defer func() {
		teardown()
		if _, serr := os.Stat(MigrateFileName); serr == nil {
			os.Remove(MigrateFileName)
		}
	}()
	out := &bytes.Buffer{}
	err = cmdMigrate.Run(newContext(out), []string{"testdata/other_key"})
	if err != nil {
		t.Error(err)
	}
	if _, err = os.Stat(MigrateFileName); err != nil {
		t.Errorf("Migrate should create %s: %v", MigrateFileName, err)
	}
	src, err := ioutil.ReadFile("testdata/data.yml")
	if err != nil {
		t.Error(err)
	}
	dst, err := ioutil.ReadFile(MigrateFileName)
	if err != nil {
		t.Error(err)
	}
	if string(src) == string(dst) {
		t.Error("Migrate should create other data")
	}
}

func TestCmdMigrateWithNotExistOtherKeyFile(t *testing.T) {
	teardown, err := setupTestData()
	if err != nil {
		if teardown != nil {
			teardown()
		}
		t.Error(err)
	}
	defer func() {
		teardown()
		if _, serr := os.Stat(MigrateFileName); serr == nil {
			os.Remove(MigrateFileName)
		}
	}()
	out := &bytes.Buffer{}
	err = cmdMigrate.Run(newContext(out), []string{"testdata/not_exist"})
	if err == nil {
		t.Error("Migrate with not exist other key file should raise error")
	}
}

func TestCmdMigrateWithoutOtherKeyFile(t *testing.T) {
	teardown, err := setupTestData()
	if err != nil {
		if teardown != nil {
			teardown()
		}
		t.Error(err)
	}
	defer func() {
		teardown()
		if _, serr := os.Stat(MigrateFileName); serr == nil {
			os.Remove(MigrateFileName)
		}
	}()
	out := &bytes.Buffer{}
	err = cmdMigrate.Run(newContext(out), []string{})
	if err == nil {
		t.Error("Migrate without other key file should raise error")
	}
}
