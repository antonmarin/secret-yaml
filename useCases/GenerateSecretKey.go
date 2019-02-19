package useCases

type GenerateSecretKey struct {
	generator SecretKeyGenerator
}

func NewGenerateSecretKey(generator SecretKeyGenerator) *GenerateSecretKey {
	return &GenerateSecretKey{generator: generator}
}

func (useCase GenerateSecretKey) Execute() ([]byte, error) {
	return useCase.generator.GenerateSecretKey()
}

type SecretKeyGenerator interface {
	GenerateSecretKey() ([]byte, error)
}
