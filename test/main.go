package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	// key := "your-32-byte-long-key-here------" // Must be 32 bytes for AES-256
	key := "testKeySizeEqualTo32Bitsfortest." // Must be 32 bytes for AES-256
	plainText := "Hello, Gopher!"

	fmt.Printf("key size: %d\n", len(key))

	encrypted, err := encrypt(plainText, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", encrypted, "lenght:", len(encrypted))

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", decrypted)
}

func decrypt(cipherText string, keyString string) (string, error) {
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

func encrypt(plainText string, keyString string) (string, error) {
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
