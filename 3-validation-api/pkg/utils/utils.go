package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHash() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	hash := hex.EncodeToString(b)

	return hash, nil
}
