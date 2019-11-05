package generateSecretKey

type generateSecretKey struct {
	generator SecretKeyGenerator
}

func NewGenerateSecretKey(generator SecretKeyGenerator) *generateSecretKey {
	return &generateSecretKey{generator: generator}
}

func (useCase generateSecretKey) Execute() ([]byte, error) {
	return useCase.generator.GenerateSecretKey()
}

//SecretKeyGenerator generates some random byte[]
type SecretKeyGenerator interface {
	GenerateSecretKey() ([]byte, error)
}
