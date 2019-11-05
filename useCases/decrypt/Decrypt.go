package decrypt

import "gopkg.in/yaml.v2"

/*
Decrypt use case decrypts data with `encryptionService` using `documentManipulator`.
`encryptionService` encapsulates logic of encrypting, and `documentManipulator`
handles data as document to encrypt all or part of data.
*/
type Decrypt struct {
	encryptionService   DecryptionService
	documentManipulator YamlManipulator
}

//NewDecrypt is a Decrypt usecase constructor. Use it only to create Decrypt
func NewDecrypt(encryptionService DecryptionService, documentManipulator YamlManipulator) *Decrypt {
	return &Decrypt{encryptionService: encryptionService, documentManipulator: documentManipulator}
}

//Execute starts Decrypt usecase
func (useCase Decrypt) Execute(dataYaml string) (string, error) {
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

//DecryptionService decrypts bytes
type DecryptionService interface {
	Decrypt([]byte) ([]byte, error)
}

//YamlManipulator applies callback to document.
//data types: yaml.MapSlice, yaml.MapItem, and Exact value
type YamlManipulator interface {
	ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error)
}
