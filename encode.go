package main

import "encoding/base64"

// Encode data with base64
func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode data with base64
func Decode(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}
