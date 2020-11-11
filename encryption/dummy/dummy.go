package dummy

import "fmt"

//EncryptionService implements no encryption algorithm
type EncryptionService struct {
}

//NewEncryptionService constructs valid EncryptionService
func NewEncryptionService(secret string) (*EncryptionService, error) {
	// todo: do smth with `32`
	if len(secret) != 32 {
		return nil, fmt.Errorf("secret should be 32 byte length")
	}

	return &EncryptionService{}, nil
}

//Encrypt returns string same as input
func (EncryptionService) Encrypt([]byte) (data []byte, err error) {
	return data, nil
}

//Decrypt returns string same as input
func (service EncryptionService) Decrypt(data []byte) ([]byte, error) {
	return data, nil
}
