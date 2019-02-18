package useCases

import (
	"github.com/antonmarin/secret-yaml/io"
)

type Encrypt struct {
	input Input
}

func (Encrypt) Execute() ([]byte, error) {
	panic("implement me")
}

type Input interface {
	AsBytes() ([]byte, error)
}

var EncryptUseCase Encrypt

func init() {
	EncryptUseCase.input = new(io.File)
}
