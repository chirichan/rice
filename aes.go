package rice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// 创建 AES-GCM 密钥和 nonce
func createCipher(key []byte) (cipher.AEAD, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, nil, err
	}
	return aesGCM, nonce, nil
}

func AESGCMEncryptText(key, plaintext string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	aesGCM, nonce, err := createCipher(keyBytes)
	if err != nil {
		return "", err
	}
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

func AESGCMDecryptText(key string, ciphertextHex string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func AESGCMEncryptFile(key, inputFile, outputFile string) error {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return err
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	aesGCM, nonce, err := createCipher(keyBytes)
	if err != nil {
		return err
	}

	// 将 nonce 写入输出文件
	_, err = outFile.Write(nonce)
	if err != nil {
		return err
	}

	buf := make([]byte, 4096)
	for {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		ciphertext := aesGCM.Seal(nil, nonce, buf[:n], nil)
		_, err = outFile.Write(ciphertext)
		if err != nil {
			return err
		}
	}

	return nil
}

func AESGCMDecryptFile(key, inputFile, outputFile string) error {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return err
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 获取 AES-GCM 密钥
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := aesGCM.NonceSize()
	nonce := make([]byte, nonceSize)
	_, err = io.ReadFull(inFile, nonce)
	if err != nil {
		return err
	}

	buf := make([]byte, 4096+aesGCM.Overhead())
	for {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		plaintext, err := aesGCM.Open(nil, nonce, buf[:n], nil)
		if err != nil {
			return err
		}
		_, err = outFile.Write(plaintext)
		if err != nil {
			return err
		}
	}

	return nil
}
