package domain

import (
	"crypto/rand"
	"encoding/hex"
)

type AesEncryptionService struct {
}

func (service AesEncryptionService) GenerateSecretKey() ([]byte, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, err
	}

	result := []byte(hex.EncodeToString(randomBytes))

	return result, nil
}
