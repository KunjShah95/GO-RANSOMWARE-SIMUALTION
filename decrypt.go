package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func DecryptFileInDir(filename string, key []byte) error {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	nonce := ciphertext[:12]
	ciphertext = ciphertext[12:]

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	originalFilename := filename[:len(filename)-4] // Remove .enc extension
	err = os.WriteFile(originalFilename, plaintext, 0644)
	if err != nil {
		return err
	}

	LogAction("Decrypted: " + originalFilename)
	fmt.Println("Decrypted:", originalFilename)
	os.Remove(filename) // Remove encrypted file
	return nil
}

func DecryptDirectoryInDir(dir string, key []byte) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".enc" {
			return DecryptFileInDir(path, key)
		}
		return nil
	})
}
