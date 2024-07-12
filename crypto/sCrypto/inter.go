package sCrypto

type Inter interface {
	Encrypt(plain string) (string, error)
	Decrypt(encrypted string) (string, error)
}

type Base struct {
	keyHex string
}

func NewBase(keyHex string) Base {
	return Base{keyHex: keyHex}
}

func (b *Base) KeyHex() string {
	return b.keyHex
}
