package cAes

import (
	"testing"
)

func TestGCM_EncryptDecrypt(t *testing.T) {
	key := "6368616e676520746869732070617373"
	gcm := NewGCM(key)

	plainText := "This is a test message."

	encryptedText, err := gcm.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	decryptedText, err := gcm.Decrypt(encryptedText)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be %s, but got %s", plainText, decryptedText)
	}
}

func TestGCM_Encrypt_EmptyPlainText(t *testing.T) {
	key := "6368616e676520746869732070617373"
	gcm := NewGCM(key)

	_, err := gcm.Encrypt("")
	if err == nil {
		t.Fatal("Expected error when encrypting empty plain text, but got none")
	}
}

func TestGCM_Decrypt_EmptyEncryptedText(t *testing.T) {
	key := "6368616e676520746869732070617373"
	gcm := NewGCM(key)

	_, err := gcm.Decrypt("")
	if err == nil {
		t.Fatal("Expected error when decrypting empty encrypted text, but got none")
	}
}
