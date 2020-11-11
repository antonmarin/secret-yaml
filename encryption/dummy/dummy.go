package dummy

// EncryptionService implements no encryption algorithm
type EncryptionService struct {
}

// Encrypt returns string same as input
func (EncryptionService) Encrypt([]byte) (data []byte, err error) {
	return data, nil
}
