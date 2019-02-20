package useCases

type Encrypt struct {
}

func NewEncrypt() *Encrypt {
	return &Encrypt{}
}

func (Encrypt) Execute(secret string, yaml string) (string, error) {
	// make yaml flat
	// encrypt each value
	// make yaml tree
	return yaml, nil
}
