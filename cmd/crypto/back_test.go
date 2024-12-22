package crypto

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "your-secure-passphrase"
	plainText := "nada::dos::tres::::cuatro"

	back := NewBack(key)

	encryptedText, err := back.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}

	decryptedText, err := back.Decrypt(encryptedText)
	if err != nil {
		t.Fatalf("Error decrypting: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Expected %s, got %s", plainText, decryptedText)
	}
}

func TestEncryptDecryptWithDifferentKey(t *testing.T) {
	key1 := "your-secure-passphrase"
	key2 := "another-secure-passphrase"
	plainText := "nada::dos::tres::::cuatro"

	back1 := NewBack(key1)
	back2 := NewBack(key2)

	encryptedText, err := back1.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}

	_, err = back2.Decrypt(encryptedText)
	if err == nil {
		t.Fatalf("Expected error decrypting with different key, got nil")
	}
}
