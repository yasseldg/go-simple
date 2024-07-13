package cAes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/yasseldg/go-simple/crypto/sCrypto"
)

type CBC struct {
	sCrypto.Base
}

func NewCBC(keyHex string) *CBC {
	return &CBC{Base: sCrypto.NewBase(keyHex)}
}

// encrypt

func (cr *CBC) Encrypt(plain string) (string, error) {

	if len(plain) == 0 {
		return "", errors.New("empty plain text")
	}

	// Decode hexadecimal key to bytes
	key, err := hex.DecodeString(cr.KeyHex())
	if err != nil {
		return "", fmt.Errorf("error decoding key: %w", err)
	}

	// Create the cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	// Create the IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("error generating IV: %w", err)
	}

	// Apply PKCS#7 padding
	paddedPlain := pkcs7Pad([]byte(plain), aes.BlockSize)

	// Create the CBC encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(paddedPlain))
	mode.CryptBlocks(ciphertext, paddedPlain)

	// Combine IV and ciphertext
	finalCiphertext := append(iv, ciphertext...)

	// Encode the result in Base64
	encryptedBase64 := base64.StdEncoding.EncodeToString(finalCiphertext)

	return encryptedBase64, nil
}

// decrypt

func (cr *CBC) Decrypt(encryptedBase64 string) (string, error) {

	if len(encryptedBase64) == 0 {
		return "", errors.New("empty encrypted text")
	}

	// Decode hexadecimal key to bytes
	key, err := hex.DecodeString(cr.KeyHex())
	if err != nil {
		return "", fmt.Errorf("error decoding key: %w", err)
	}

	// Decode the Base64 encrypted text to bytes
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("error decoding Base64: %w", err)
	}

	// Extract the IV from the first 16 bytes
	iv := encryptedData[:16]
	ciphertext := encryptedData[16:]

	// Create the cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	// Create the CBC decrypter
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the text
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove the padding
	plaintext = pkcs7Unpad(plaintext)

	return string(plaintext), nil
}
