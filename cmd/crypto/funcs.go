package crypto

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sStrings"
)

func Run() {
	passphrase := "your-secure-passphrase"
	originalText := "nada::dos::tres::::cuatro"

	back := NewBack(passphrase)

	encryptedText, err := back.Encrypt(originalText)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted:", encryptedText)

	decryptedText, err := back.Decrypt(encryptedText)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", decryptedText)

	values := sStrings.SplitString(decryptedText, "::")

	last := values[len(values)-1]
	keys := values[1 : len(values)-1]

	sLog.Warn("keys: %v  ..  last: %s", keys, last)

}
