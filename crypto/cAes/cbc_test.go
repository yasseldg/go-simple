package cAes

import (
	"testing"
)

func TestCBC_EncryptDecrypt(t *testing.T) {
	key := "6368616e676520746869732070617373"
	cbc := NewCBC(key)

	plainText := "This is a test message."

	encryptedText, err := cbc.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	decryptedText, err := cbc.Decrypt(encryptedText)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be %s, but got %s", plainText, decryptedText)
	}
}

func TestCBC_Encrypt_EmptyPlainText(t *testing.T) {
	key := "6368616e676520746869732070617373"
	cbc := NewCBC(key)

	_, err := cbc.Encrypt("")
	if err == nil {
		t.Fatal("Expected error when encrypting empty plain text, but got none")
	}
}

func TestCBC_Decrypt_EmptyEncryptedText(t *testing.T) {
	key := "6368616e676520746869732070617373"
	cbc := NewCBC(key)

	_, err := cbc.Decrypt("")
	if err == nil {
		t.Fatal("Expected error when decrypting empty encrypted text, but got none")
	}
}
