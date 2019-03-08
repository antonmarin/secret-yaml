package encrypt

import (
	"gopkg.in/yaml.v2"
)

type Encrypt struct {
	encryptionService   EncryptionService
	documentManipulator YamlManipulator
}

func NewEncrypt(encryptionService EncryptionService, documentManipulator YamlManipulator) *Encrypt {
	return &Encrypt{encryptionService: encryptionService, documentManipulator: documentManipulator}
}

func (useCase Encrypt) Execute(dataYaml string) (string, error) {
	document := make(yaml.MapSlice, 0)
	err := yaml.Unmarshal([]byte(dataYaml), &document)
	if err != nil {
		return "", err
	}
	encryptedDocument, err := useCase.documentManipulator.ApplyToLeafs(useCase.encryptionService.Encrypt, document)
	if err != nil {
		return "", err
	}

	encryptedYaml, err := yaml.Marshal(encryptedDocument)
	if err != nil {
		return "", err
	}

	return string(encryptedYaml), nil
}

type EncryptionService interface {
	Encrypt([]byte) ([]byte, error)
}

//data types: yaml.MapSlice, yaml.MapItem, and Exact value
type YamlManipulator interface {
	ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error)
}
