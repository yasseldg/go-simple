package cAes

import "bytes"

// private functions

// Function to apply PKCS#7 padding
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// Function to remove PKCS#7 padding
func pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}

	padding := data[len(data)-1]
	if int(padding) > len(data) {
		return data
	}
	return data[:len(data)-int(padding)]
}
