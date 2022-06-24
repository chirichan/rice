package rice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
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

func AESNewCBCDecrypter(keyString, cipherString string) (string, error) {
	key, _ := hex.DecodeString(keyString)
	ciphertext, _ := hex.DecodeString(cipherString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)

	return ByteString(ciphertext), nil
}

func AESNewCBCEncrypter(keyString, plainString string) (string, error) {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(plainString)

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)

	return hex.EncodeToString(ciphertext), nil
}

func AESNewCFBDecrypter() {}

func AESNewCFBEncrypter() {}

func AESNewCTREncrypter(keyString string, plainString string) string {

	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(plainString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	return hex.EncodeToString(plaintext2)
}

const (
	CBC = 1 + iota
	CFB
	CTR
	GCM
	OFB
)

type AES interface {
	Decrypter()
	Encrypter()
}
