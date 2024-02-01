package external

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	"github.com/google/uuid"
)

type ExternalServices struct{}

func NewExternalServices() *ExternalServices { return &ExternalServices{} }

func (e *ExternalServices) GenUuid() (uid string) {
	return uuid.New().String()
}

func (h *ExternalServices) HashString(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}

func (e *ExternalServices) Encrypt(plainText string, keyString string) (string, error) {
	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	cipherBlock := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherBlock[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBlock[aes.BlockSize:], plainTextBytes)

	return base64.StdEncoding.EncodeToString(cipherBlock), nil
}

func (e *ExternalServices) Decrypt(cipherText string, keyString string) (string, error) {
	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(cipherTextBytes) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}
	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}
