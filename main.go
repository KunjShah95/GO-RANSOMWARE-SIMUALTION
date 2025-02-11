package main

import (
	"fmt"
	"os"
	"path/filepath"
	"crypto/aes"
	"crypto/cipher"
)


func EncryptDirectory(directory string, key []byte) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = EncryptFile(path, key)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func DecryptDirectory(directory string, key []byte) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = DecryptFile(path, key)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func EncryptFile(filename string, key []byte) error {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return os.WriteFile(filename, ciphertext, 0777)
}

func DecryptFile(filename string, key []byte) error {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, plaintext, 0777)
}

func GenerateEncryptionKey() []byte {
	// Implement key generation logic here
	// For example, return a dummy key for now
	return []byte("dummykeydummykeydummykeydummykey")
}

func SaveKey(key []byte) error {
	return os.WriteFile("encryption.key", key, 0644)
}

func LoadKey() ([]byte, error) {
	return os.ReadFile("encryption.key")
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run Main.go <encrypt|decrypt> <directory>")
		return
	}
	action := os.Args[1]
	directory := os.Args[2]

	key := GenerateEncryptionKey() // Generate AES-256 key

	if action == "encrypt" {
		err := EncryptDirectory(directory, key)
		if err != nil {
			fmt.Println("Error encrypting files:", err)
			return
		}
		SaveKey(key)
		fmt.Println("Encryption completed.")
	} else if action == "decrypt" {
		key, err := LoadKey()
		if err != nil {
			fmt.Println("Error loading key:", err)
			return
		}
		err = DecryptDirectory(directory, key)
		if err != nil {
			fmt.Println("Error decrypting files:", err)
			return
		}
		fmt.Println("Decryption completed.")
	} else {
		fmt.Println("Invalid command. Use encrypt or decrypt.")
	}
}
