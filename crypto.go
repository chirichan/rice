package rice

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Signer interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, signature []byte) error
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
