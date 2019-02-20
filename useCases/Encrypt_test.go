package useCases

import (
	"testing"
)

func TestEncrypt_Execute(t *testing.T) {
	secret := "asdf"
	useCase := NewEncrypt()
	data := `---
some: value`

	result, err := useCase.Execute(secret, data)
	if err != nil {
		t.Errorf("Should not fail if data valid. Error: %s", err)
	}

	expectedData := `---
some: 123`
	if expectedData != result {
		t.Errorf("Should encrypt. Got: %s", result)
	}
}
