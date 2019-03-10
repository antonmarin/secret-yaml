/*
Package decrypt use case decrypts data with `encryptionService` using `documentManipulator`.
`encryptionService` encapsulates logic of encrypting, and `documentManipulator`
handles data as document to encrypt all or part of data.
*/
package decrypt

import "gopkg.in/yaml.v2"

type decrypt struct {
	encryptionService   DecryptionService
	documentManipulator YamlManipulator
}

func NewDecrypt(encryptionService DecryptionService, documentManipulator YamlManipulator) *decrypt {
	return &decrypt{encryptionService: encryptionService, documentManipulator: documentManipulator}
}

func (useCase decrypt) Execute(dataYaml string) (string, error) {
	document := make(yaml.MapSlice, 0)
	err := yaml.Unmarshal([]byte(dataYaml), &document)
	if err != nil {
		return "", err
	}
	encryptedDocument, err := useCase.documentManipulator.ApplyToLeafs(useCase.encryptionService.Decrypt, document)
	if err != nil {
		return "", err
	}

	encryptedYaml, err := yaml.Marshal(encryptedDocument)
	if err != nil {
		return "", err
	}

	return string(encryptedYaml), nil
}

type DecryptionService interface {
	Decrypt([]byte) ([]byte, error)
}

//data types: yaml.MapSlice, yaml.MapItem, and Exact value
type YamlManipulator interface {
	ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error)
}
