package generateSecretKey

type generateSecretKey struct {
	generator SecretKeyGenerator
}

func New(generator SecretKeyGenerator) *generateSecretKey {
	return &generateSecretKey{generator: generator}
}

func (useCase generateSecretKey) Execute() ([]byte, error) {
	return useCase.generator.GenerateSecretKey()
}

type SecretKeyGenerator interface {
	GenerateSecretKey() ([]byte, error)
}
