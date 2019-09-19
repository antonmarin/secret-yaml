package random

import (
	"crypto/rand"
	"encoding/hex"
)

type CryptoGeneratorService struct {
}

func (service CryptoGeneratorService) GenerateSecretKey() ([]byte, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, err
	}

	result := []byte(hex.EncodeToString(randomBytes))

	return result, nil
}
