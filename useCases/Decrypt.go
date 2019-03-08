package useCases

type Decrypt struct {
	encryptionService   EncryptionService
	documentManipulator YamlManipulator
}

func NewDecrypt(encryptionService EncryptionService, documentManipulator YamlManipulator) *Encrypt {
	return &Encrypt{encryptionService: encryptionService, documentManipulator: documentManipulator}
}
