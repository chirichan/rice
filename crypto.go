package rice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// BCryptGenerateFromPassword generate hash from password
func BCryptGenerateFromPassword(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword(StringByte(pwd), bcrypt.DefaultCost)
	return ByteString(password), err
}

// BCryptCompareHashAndPassword true or false
func BCryptCompareHashAndPassword(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword(StringByte(pwd), StringByte(hash)) == nil
}

func AESEncrypt(key []byte, s string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	dst := make([]byte, len(s))
	block.Encrypt(dst, StringByte(s))

	return hex.EncodeToString(dst), nil
}

func AESDecrypt(key []byte, s string) (string, error) {
	decodeBytes, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	dst := make([]byte, len(s))
	block.Decrypt(dst, decodeBytes)

	return ByteString(dst), nil
}

func AESNewGCMEncrypt(keyString, plaintext string) (string, error) {

	key, _ := hex.DecodeString(keyString)
	//b, _ := hex.DecodeString(plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	fmt.Println(hex.EncodeToString(nonce))

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//nonceSize := aesgcm.NonceSize()

	ciphertext := aesgcm.Seal(nil, nonce, StringByte(plaintext), nil)

	return hex.EncodeToString(ciphertext), nil
}

func AESNewGCMDecrypt(keyString, nonceString, ciphertext string) (string, error) {
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(nonceString)
	if err != nil {
		return "", err
	}

	cipherByte, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, cipherByte, nil)
	if err != nil {
		return "", err
	}
	return ByteString(plaintext), nil
}

func AESNewCBCDecrypter() {}

func AESNewCBCEncrypter() {

}

func AESNewCFBDecrypter() {}

func AESNewCFBEncrypter() {}
