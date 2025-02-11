package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

const keyFile = "encryption.key"

func GenerateKey() []byte {
	key := make([]byte, 32) // AES-256 requires a 32-byte key
	rand.Read(key)
	return key
}

func SaveEncryptionKey(key []byte) {
	keyHex := hex.EncodeToString(key)
	os.WriteFile(keyFile, []byte(keyHex), 0644)
	fmt.Println("Encryption key saved.")
}

func LoadEncryptionKey() ([]byte, error) {
	data, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	return key, nil
}
