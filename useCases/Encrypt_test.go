package useCases

import (
	"testing"
)

func TestEncrypt_Execute(t *testing.T) {
	secret := "asdf"
	useCase := NewEncrypt()
	var resultOfEncrypt interface{}
	var err interface{}
	var expectedData interface{}

	simpleData := `---
key: value`
	resultOfEncrypt, err = useCase.Execute(secret, simpleData)
	if err != nil {
		t.Errorf("Should not fail if simpleData valid. Error: %s", err)
	}
	expectedData = `---
key: valueEncryptedWithSecretAsdf`
	if expectedData != resultOfEncrypt {
		t.Errorf("Should encrypt. Got: %s", resultOfEncrypt)
	}

	nestedData := `---
key:
  nestedKey: valueEncryptedWithSecretAsdf`
	resultOfEncrypt, err = useCase.Execute(secret, nestedData)
	if err != nil {
		t.Errorf("Should not fail if nestedData valid. Error: %s", err)
	}
	expectedData = `---
key: 
  nestedKey: valueEncryptedWithSecretAsdf`
	if expectedData != resultOfEncrypt {
		t.Errorf("Should encrypt only values of nestedData.\nExpected: %s\nGot: %s", expectedData, resultOfEncrypt)
	}

	//	Should return error if yaml invalid
}
