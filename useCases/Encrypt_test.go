package useCases

import "testing"

func TestEncrypt_Execute(t *testing.T) {
	useCase := NewEncrypt()
	data := []byte(`---
some: value`)
	result, err := useCase.Execute(data)

}
