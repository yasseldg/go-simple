package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Back struct {
	key []byte
}

// NewBack crea una nueva instancia de Back con una clave proporcionada en formato hexadecimal
func NewBack(keyHex string) *Back {
	key := generateKey(keyHex)

	sLog.Info("Key:", key)

	return &Back{key: key}
}

// generateKey genera una clave de 256 bits a partir de una frase secreta utilizando SHA-256
func generateKey(passphrase string) []byte {
	hash := sha256.Sum256([]byte(passphrase))
	return hash[:]
}

// Encrypt cifra un texto plano usando AES-GCM y devuelve el texto cifrado en formato base64
func (b *Back) Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aead.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt descifra un texto cifrado en formato base64 usando AES-GCM
func (b *Back) Decrypt(encryptedText string) (string, error) {
	block, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
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
