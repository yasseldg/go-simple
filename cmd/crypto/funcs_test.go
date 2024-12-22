package crypto

import (
	"testing"
)

func TestRun(t *testing.T) {
	passphrase := "your-secure-passphrase"
	originalText := "nada::dos::tres::::cuatro"

	back := NewBack(passphrase)

	encryptedText, err := back.Encrypt(originalText)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}

	decryptedText, err := back.Decrypt(encryptedText)
	if err != nil {
		t.Fatalf("Error decrypting: %v", err)
	}

	if decryptedText != originalText {
		t.Errorf("Expected %s, got %s", originalText, decryptedText)
	}
}
