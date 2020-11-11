package encrypt

import (
	encryption "github.com/antonmarin/secret-yaml/encryption/dummy"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestEncrypt_Execute(t *testing.T) {
	useCase := NewEncrypt(new(encryption.EncryptionService), new(FakeYamlDocumentManipulator))
	var resultOfEncrypt interface{}
	var err interface{}
	var expectedData interface{}

	simpleData := `key: value`
	resultOfEncrypt, err = useCase.Execute(simpleData)
	if err != nil {
		t.Errorf("Should not fail if simpleData valid. Error: %s", err)
	}
	expectedData = "key: valueEncryptedWithSecretAsdf\n"
	if expectedData != resultOfEncrypt {
		t.Errorf("Should encrypt with service and manipulator.\nExpected:\n%s\nGot:\n%s", expectedData, resultOfEncrypt)
	}
}

type FakeYamlDocumentManipulator struct {
}

func (FakeYamlDocumentManipulator) ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error) {
	return yaml.MapSlice{yaml.MapItem{Key: "key", Value: "valueEncryptedWithSecretAsdf"}}, nil
}
