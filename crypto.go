package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
)

// GetKey returns key for encrypt from file on given path.
func GetKey(fp string) ([]byte, error) {
	if _, err := os.Stat(fp); err != nil {
		return nil, err
	}
	keySrc, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return GenKey(keySrc), nil
}

// GenKey generates key for encrypt from src bytes.
func GenKey(src []byte) []byte {
	hash := sha256.Sum256(src)
	return hash[:]
}

// Encrypt text with AES-256 and encode with base64
func Encrypt(key []byte, data []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)
	return Encode(cipherText), nil
}

// Decrypt data with AES-256
func Decrypt(key []byte, encrypted string) ([]byte, error) {
	data, err := Decode(encrypted)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	src := data[aes.BlockSize:]
	dst := make([]byte, len(src))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(dst, src)
	return dst, nil
}

// Encode data with base64
func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode data with base64
func Decode(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}
