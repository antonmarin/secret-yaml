package encryption

type AesEncryptionService struct {
}

func (AesEncryptionService) Encrypt([]byte) ([]byte, error) {
	panic("implement me")
}

func NewAesEncryptionService() *AesEncryptionService {
	return &AesEncryptionService{}
}
