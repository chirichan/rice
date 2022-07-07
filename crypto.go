package rice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const (
	CBC = 1 + iota
	CFB
	CTR
	GCM
	OFB
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func CheckPassword(pwd string) error {
	if len(pwd) < 16 {
		return errors.New("至少16位")
	}

	hasStrFunc := func(s rune, str string) bool {
		for _, i := range str {
			if i == s {
				return true
			}
		}
		return false
	}

	var pwdmap, lowermap, uppermap, digitsmap, symbolsmap, othermap = make(map[rune]struct{}),
		make(map[rune]struct{}),
		make(map[rune]struct{}),
		make(map[rune]struct{}),
		make(map[rune]struct{}),
		make(map[rune]struct{})

	for _, v := range pwd {
		pwdmap[v] = struct{}{}

		if hasStrFunc(v, LowerLetters) {
			lowermap[v] = struct{}{}
		} else if hasStrFunc(v, UpperLetters) {
			uppermap[v] = struct{}{}
		} else if hasStrFunc(v, Digits) {
			digitsmap[v] = struct{}{}
		} else if hasStrFunc(v, Symbols) {
			symbolsmap[v] = struct{}{}
		} else {
			othermap[v] = struct{}{}
		}
	}

	if len(pwdmap) < 10 {
		return errors.New("去重后至少10位")
	}

	if len(lowermap) < 4 {
		return errors.New("小写字母至少4个")
	}

	if len(uppermap) < 4 {
		return errors.New("大写字母至少4个")
	}

	if len(digitsmap) > 6 {
		return errors.New("数字不能超过6个")
	}

	if len(symbolsmap) < 4 {
		return errors.New("特殊符号至少4个")
	}

	return nil
}

// BCryptGenerateFromPassword generate hash from password
func BCryptGenerateFromPassword(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword(StringByteUnsafe(pwd), 14)
	return ByteString(password), err
}

// BCryptCompareHashAndPassword true or false
func BCryptCompareHashAndPassword(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword(StringByteUnsafe(hash), StringByteUnsafe(pwd)) == nil
}

type AESCrypter interface {
	Encrypt(keyString, plainString string) (string, error)
	Decrypt(keyString, cipherString string) (string, error)
}

type CTRCrypt struct{}

func NewCTRCrypt() AESCrypter { return &CTRCrypt{} }

func (*CTRCrypt) Encrypt(keyString, plainString string) (string, error) {

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

func (*CTRCrypt) Decrypt(keyString, cipherString string) (string, error) {

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

	plaintext2 := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	return string(plaintext2), nil
}
