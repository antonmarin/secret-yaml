package useCases

import (
	"github.com/antonmarin/secret-yaml/generateRandom"
)

type GenerateSecretKey struct {
	generator SecretKeyGenerator
}

func (useCase GenerateSecretKey) Execute() ([]byte, error) {
	return useCase.generator.GenerateSecretKey()
}

type SecretKeyGenerator interface {
	GenerateSecretKey() ([]byte, error)
}

var GenerateSecretKeyUseCase GenerateSecretKey

func init() {
	GenerateSecretKeyUseCase.generator = new(generateRandom.CryptoGeneratorService)
}
