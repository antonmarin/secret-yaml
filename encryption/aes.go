package encryption

type AesEncryptionService struct {
	secret string
}

func (AesEncryptionService) Encrypt([]byte) ([]byte, error) {
	panic("implement me")
}

func NewAesEncryptionService(secret string) *AesEncryptionService {
	return &AesEncryptionService{secret: secret}
}
