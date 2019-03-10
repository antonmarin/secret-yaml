package encrypt

import "gopkg.in/yaml.v2"

type encrypt struct {
	encryptionService   EncryptionService
	documentManipulator YamlManipulator
}

//NewEncrypt constructor of encrypt use case
func NewEncrypt(encryptionService EncryptionService, documentManipulator YamlManipulator) *encrypt {
	return &encrypt{encryptionService: encryptionService, documentManipulator: documentManipulator}
}

func (useCase encrypt) Execute(dataYaml string) (string, error) {
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
