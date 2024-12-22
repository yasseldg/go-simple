package sCrypto

import (
	"testing"
)

type MockCrypto struct {
	key string
}

func (m *MockCrypto) Encrypt(plain string) (string, error) {
	return "encrypted:" + plain, nil
}

func (m *MockCrypto) Decrypt(encrypted string) (string, error) {
	return encrypted[len("encrypted:"):], nil
}

func TestEncryptDecrypt(t *testing.T) {
	mock := &MockCrypto{key: "testkey"}

	plainText := "Hello, World!"
	encryptedText, err := mock.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	decryptedText, err := mock.Decrypt(encryptedText)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be %s, but got %s", plainText, decryptedText)
	}
}

func TestEncryptEmptyString(t *testing.T) {
	mock := &MockCrypto{key: "testkey"}

	encryptedText, err := mock.Encrypt("")
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	if encryptedText != "encrypted:" {
		t.Errorf("Expected encrypted text to be 'encrypted:', but got %s", encryptedText)
	}
}

func TestDecryptEmptyString(t *testing.T) {
	mock := &MockCrypto{key: "testkey"}

	decryptedText, err := mock.Decrypt("encrypted:")
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedText != "" {
		t.Errorf("Expected decrypted text to be '', but got %s", decryptedText)
	}
}
