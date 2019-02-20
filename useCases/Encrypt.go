package useCases

type Encrypt struct {
}

func NewEncrypt() *Encrypt {
	return &Encrypt{}
}

func (Encrypt) Execute(secret string, yaml []byte) ([]byte, error) {
	panic("implement me")
}
