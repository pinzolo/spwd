package main

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	data := []byte("password")
	e := Encode(data)
	d, err := Decode(e)
	if err != nil {
		t.Error(err)
	}
	if string(d) != "password" {
		t.Error("Decorde failure")
	}
}
