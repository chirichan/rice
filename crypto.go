package rice

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	_ AESCrypter = &CTRCrypt{}
	_ Signer     = &RSASign{}
)

type AESCrypter interface {
	Encrypt(keyString, plainString string) (string, error)
	Decrypt(keyString, cipherString string) (string, error)
}

type Signer interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, signature []byte) error
}

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

const (
	_defaultLength          = 16
	_defaultNumLowerLetters = 4
	_defaultNumUpperLetters = 4
	_defaultNumDigits       = 4
	_defaultNumSymbols      = 4
)

type FullPasswordConf struct {
	Length          int
	NumLowerLetters int
	NumUpperLetters int
	NumDigits       int
	NumSymbols      int
}

func SetLevel(level, length int) FullPasswordConf {

	var fullConf FullPasswordConf

	if level == 1 {
		fullConf.NumDigits = length
	} else if level == 2 {
		fullConf.NumLowerLetters = length / 2
		fullConf.NumDigits = length - fullConf.NumLowerLetters
	} else if level == 3 {
		fullConf.NumDigits = length / 3
		fullConf.NumUpperLetters = (length - fullConf.NumDigits) / 2
		fullConf.NumLowerLetters = length - fullConf.NumDigits - fullConf.NumUpperLetters
	} else if level == 4 {
		fullConf.NumDigits = length / 5
		fullConf.NumUpperLetters = length / 4
		fullConf.NumLowerLetters = length / 4
		fullConf.NumSymbols = length - fullConf.NumDigits - fullConf.NumUpperLetters - fullConf.NumLowerLetters
	} else {
		fullConf.Length = _defaultLength
		fullConf.NumLowerLetters = _defaultNumLowerLetters
		fullConf.NumUpperLetters = _defaultNumUpperLetters
		fullConf.NumDigits = _defaultNumDigits
		fullConf.NumSymbols = _defaultNumSymbols
	}
	return fullConf
}

func FullPassword(level, length int) (string, error) {

	if length < 6 {
		return "", errors.New("length must >= 6")
	} else if length > 2048 {
		return "", errors.New("length too long")
	}

	if length < 1 || length > 4 {
		return "", errors.New("level must range 1-4")
	}

	var (
		result string
		read   = rand.Reader
	)

	var fullConf = SetLevel(level, length)

	// Characters
	for i := 0; i < fullConf.NumLowerLetters; i++ {
		ch, err := randomElement(read, LowerLetters)
		if err != nil {
			return "", err
		}

		result, err = randomInsert(read, result, ch)
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < fullConf.NumUpperLetters; i++ {
		ch, err := randomElement(read, UpperLetters)
		if err != nil {
			return "", err
		}

		result, err = randomInsert(read, result, ch)
		if err != nil {
			return "", err
		}
	}

	// Digits
	for i := 0; i < fullConf.NumDigits; i++ {
		d, err := randomElement(read, Digits)
		if err != nil {
			return "", err
		}

		result, err = randomInsert(read, result, d)
		if err != nil {
			return "", err
		}
	}

	// Symbols
	for i := 0; i < fullConf.NumSymbols; i++ {
		sym, err := randomElement(read, Symbols)
		if err != nil {
			return "", err
		}

		result, err = randomInsert(read, result, sym)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

// randomInsert randomly inserts the given value into the given string.
func randomInsert(reader io.Reader, s, val string) (string, error) {
	if s == "" {
		return val, nil
	}

	n, err := rand.Int(reader, big.NewInt(int64(len(s)+1)))
	if err != nil {
		return "", err
	}
	i := n.Int64()
	return s[0:i] + val + s[i:], nil
}

// randomElement extracts a random element from the given string.
func randomElement(reader io.Reader, s string) (string, error) {
	n, err := rand.Int(reader, big.NewInt(int64(len(s))))
	if err != nil {
		return "", err
	}
	return string(s[n.Int64()]), nil
}

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
	password, err := bcrypt.GenerateFromPassword(StringByteUnsafe(pwd), bcrypt.DefaultCost)
	return ByteString(password), err
}

// BCryptCompareHashAndPassword true or false
func BCryptCompareHashAndPassword(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword(StringByteUnsafe(hash), StringByteUnsafe(pwd)) == nil
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

	plaintext2 := make([]byte, len(ciphertext[aes.BlockSize:]))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	return string(plaintext2), nil
}

type RSASign struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewRSASignFromFile(privateKeyFile, publicKeyFile string) (Signer, error) {

	privateKeyData, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	publicKeyData, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}

	return NewRSASignFromBytes(privateKeyData, publicKeyData)
}

func NewRSASignFromBytes(privateKeyData, publicKeyData []byte) (Signer, error) {

	privatePEM, _ := pem.Decode(privateKeyData)
	publicPEM, _ := pem.Decode(publicKeyData)

	privateKey, err := x509.ParsePKCS1PrivateKey(privatePEM.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicPEM.Bytes)
	if err != nil {
		return nil, err
	}

	return &RSASign{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

func (r *RSASign) Sign(data []byte) ([]byte, error) {
	h := sha256.New()

	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}

	hashSum := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, hashSum)
}

func (r *RSASign) Verify(data []byte, signature []byte) error {
	h := sha256.New()

	_, err := h.Write(data)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, h.Sum(nil), signature)
}
