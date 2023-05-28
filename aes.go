package rice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func AESEncrypt(keyString, plainString string) (string, error) {

	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}
	plaintext := []byte(plainString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func AESDecrypt(keyString, cipherString string) (string, error) {

	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(cipherString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]

	plaintext2 := make([]byte, len(ciphertext[aes.BlockSize:]))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	return string(plaintext2), nil
}
