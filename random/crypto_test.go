package random

import "testing"

func TestGenerateSecretKeyShouldReturnBytes32(t *testing.T) {
	service := new(CryptoGeneratorService)
	result, err := service.GenerateSecretKey()
	if err != nil {
		t.Errorf("Should not be error while simple generation. Error: %s", err)
	}
	if len(result) != 32 {
		t.Errorf("Generated secret key length should be 32, current %d(%s)", len(result), result)
	}
}
