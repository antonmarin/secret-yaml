package decrypt

import (
	"github.com/antonmarin/secret-yaml/documentManipulator"
	"github.com/antonmarin/secret-yaml/encryption"
	"testing"
)

func TestDecrypt_IntegrationAes(t *testing.T) {
	var err interface{}

	encryptionService, err := encryption.NewAesEncryptionService("75b3703e27e300afffe2aa7ba047c930")
	if err != nil {
		t.Error(err)
	}
	useCase := NewDecrypt(encryptionService, documentManipulator.NewYamlManipulator())

	var resultOfEncrypt interface{}
	var expectedData interface{}

	data := `some: !!binary lu91bQLWkSHjiGyY+d5psHQjdoWUYCKxQcg/vAJw5bH1`
	expectedData = "some: value\n"
	resultOfEncrypt, err = useCase.Execute(data)
	if err != nil {
		t.Errorf("Should not fail if data valid. Error: %s", err)
	}
	if expectedData != resultOfEncrypt {
		t.Errorf("Should decrypt with service and manipulator.\nExpected:\n%s\nGot:\n%s", expectedData, resultOfEncrypt)
	}
}
