package generateSecretKey

import (
	"bytes"
	"testing"
)

type StubSecretKeyGenerator struct {
}

func (StubSecretKeyGenerator) GenerateSecretKey() ([]byte, error) {
	return []byte("12345"), nil
}

func TestGenerateSecretKey_Execute(t *testing.T) {
	useCase := New(new(StubSecretKeyGenerator))
	result, err := useCase.Execute()
	if err != nil {
		t.Errorf("Should not fail. Error: %s", err)
	}
	expectedResult := []byte("12345")
	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Should return generated value. Current: %s", result)
	}
}
