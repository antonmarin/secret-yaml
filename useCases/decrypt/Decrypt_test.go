package decrypt

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestDecrypt_Execute(t *testing.T) {
	useCase := NewDecrypt(new(DummyDecryptEncryptionService), new(FakeDecryptYamlDocumentManipulator))
	var resultOfEncrypt interface{}
	var err interface{}
	var expectedData interface{}

	simpleData := `key: encryptedValue`
	resultOfEncrypt, err = useCase.Execute(simpleData)
	if err != nil {
		t.Errorf("Should not fail if simpleData valid. Error: %s", err)
	}
	expectedData = "key: decryptedValue\n"
	if expectedData != resultOfEncrypt {
		t.Errorf("Should decrypt with service and manipulator.\nExpected:\n%s\nGot:\n%s", expectedData, resultOfEncrypt)
	}
}

type DummyDecryptEncryptionService struct {
}

func (DummyDecryptEncryptionService) Decrypt([]byte) (data []byte, err error) {
	return data, nil
}

type FakeDecryptYamlDocumentManipulator struct {
}

func (FakeDecryptYamlDocumentManipulator) ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error) {
	return yaml.MapSlice{yaml.MapItem{Key: "key", Value: "decryptedValue"}}, nil
}
