package cAes

import (
	"bytes"
	"testing"
)

func TestPKCS7Pad(t *testing.T) {
	blockSize := 16
	data := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10")

	padded := pkcs7Pad(data, blockSize)
	if !bytes.Equal(padded, expected) {
		t.Errorf("Expected %v, got %v", expected, padded)
	}
}

func TestPKCS7Unpad(t *testing.T) {
	data := []byte("YELLOW SUBMARINE\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10")
	expected := []byte("YELLOW SUBMARINE")

	unpadded := pkcs7Unpad(data)
	if !bytes.Equal(unpadded, expected) {
		t.Errorf("Expected %v, got %v", expected, unpadded)
	}
}

func TestPKCS7Unpad_InvalidPadding(t *testing.T) {
	data := []byte("YELLOW SUBMARINE\x01\x02\x03\x04")
	expected := data

	unpadded := pkcs7Unpad(data)
	if !bytes.Equal(unpadded, expected) {
		t.Errorf("Expected %v, got %v", expected, unpadded)
	}
}
