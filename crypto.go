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
	"io"
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

// BCryptGenerateFromPassword generate hash from password
func BCryptGenerateFromPassword(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword(StringByteUnsafe(pwd), bcrypt.DefaultCost)
	return ByteString(password), err
}

// BCryptCompareHashAndPassword true or false
func BCryptCompareHashAndPassword(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword(StringByteUnsafe(hash), StringByteUnsafe(pwd)) == nil
}
