package cAes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/yasseldg/go-simple/crypto/sCrypto"
)

type GCM struct {
	sCrypto.Base
}

func NewGCM(keyHex string) *GCM {
	return &GCM{Base: sCrypto.NewBase(keyHex)}
}

func (b *GCM) Encrypt(plain string) (string, error) {

	if len(plain) == 0 {
		return "", errors.New("empty plain text")
	}

	block, err := aes.NewCipher(b.key())
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error generating nonce: %w", err)
	}

	cipherText := aead.Seal(nonce, nonce, []byte(plain), nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (b *GCM) Decrypt(encrypted string) (string, error) {

	if len(encrypted) == 0 {
		return "", errors.New("empty encrypted text")
	}

	block, err := aes.NewCipher(b.key())
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM: %w", err)
	}

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("error decoding base64: %w", err)
	}

	nonceSize := aead.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// private methods

// generateKey generates a 256-bit key from a secret phrase using SHA-256
func (b *GCM) key() []byte {
	hash := sha256.Sum256([]byte(b.Base.KeyHex()))
	return hash[:]
}
