package main

import (
	"testing"
)

func TestKeyLength(t *testing.T) {
	k1 := GenKey([]byte("short string"))
	k2 := GenKey([]byte("abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqrstuvwxyz0123456789"))

	if len(k1) != 32 {
		t.Errorf("GenKey should returns 32 byte, but got %v", len(k1))
	}
	if len(k2) != 32 {
		t.Errorf("GenKey should returns 32 byte, but got %v", len(k1))
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	key := GenKey([]byte("this is crypto key"))
	pwd := "password"

	e, err := Encrypt(key, pwd)
	if err != nil {
		t.Error(err)
	}
	d, err := Decrypt(key, e)
	if err != nil {
		t.Error(err)
	}

	if d != pwd {
		t.Errorf("Decrypt failure: %s", d)
	}
}

func TestEncryptWithInvalidKey(t *testing.T) {
	_, err := Encrypt([]byte("foobar"), "password")
	if err == nil {
		t.Error("Encrypt with invalid key should raise error")
	}
}

func TestDecryptWithInvalidKey(t *testing.T) {
	_, err := Decrypt([]byte("foobar"), []byte("password"))
	if err == nil {
		t.Error("Decrypt with invalid key should raise error")
	}
}

func TestCannotDecryptWithOtherKey(t *testing.T) {
	k1 := GenKey([]byte("this is crypto key"))
	k2 := GenKey([]byte("this is other key"))
	pwd := "password"

	e, err := Encrypt(k1, pwd)
	if err != nil {
		t.Error(err)
	}
	d, err := Decrypt(k2, e)
	if err != nil {
		t.Error(err)
	}

	if d == pwd {
		t.Errorf("Decrypt with other key should fail")
	}
}
