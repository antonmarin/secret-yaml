package encryption

import (
	"bytes"
	"testing"
)

func TestNewAesEncryptionService_ShouldCreateWith32byteKey(t *testing.T) {
	_, err := NewAesEncryptionService("fbd7908e8ccdf67d472995a5c54ea72f")
	if err != nil {
		t.Errorf("Should not error with valid secret. Error: %s", err)
	}
}

func TestNewAesEncryptionService_ShouldFailWithNot32byteKey(t *testing.T) {
	secret := "someBadSecret"
	_, err := NewAesEncryptionService(secret)
	if err == nil {
		t.Errorf("Should error on not valid secret. Secret was: %s", secret)
	}
}

func TestAesEncryptionService_DecryptedShouldEqualSource(t *testing.T) {
	//noinspection SpellCheckingInspection
	service, _ := NewAesEncryptionService("fbd7908e8ccdf67d472995a5c54ea72f")

	sourceData := []byte("someValue")

	encryptedData, err := service.Encrypt(sourceData)
	if err != nil {
		t.Errorf("Should not error on valid value. Error: %s", err)
	}
	if bytes.Equal(encryptedData, sourceData) {
		t.Errorf("Encrypted data can't equal source.\nEncrypted: %s\nSource: %s", encryptedData, sourceData)
	}

	decryptedData, err := service.Decrypt(encryptedData)
	if !bytes.Equal(sourceData, decryptedData) {
		t.Errorf("Descrypted should equal source.\nSource: %s\nDecrypted: %s", sourceData, decryptedData)
	}
}
