package documentManipulator

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
)

type YamlManipulator struct {
}

func NewYamlManipulator() *YamlManipulator {
	return &YamlManipulator{}
}

func (manipulator YamlManipulator) ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error) {
	switch data.(type) {
	case yaml.MapSlice:
		result := make(yaml.MapSlice, len(data.(yaml.MapSlice)))
		for ind, elm := range data.(yaml.MapSlice) {
			result[ind].Key = elm.Key
			resultValue, err := manipulator.ApplyToLeafs(callback, elm.Value)
			if err != nil {
				return nil, err
			}

			result[ind].Value = resultValue
		}

		return result, nil
	case yaml.MapItem:
		var result yaml.MapItem

		result.Key = data.(yaml.MapItem).Key

		resultValue, err := manipulator.ApplyToLeafs(callback, data.(yaml.MapItem).Value)
		if err != nil {
			return nil, err
		}

		result.Value = resultValue

		return result, nil
	case []interface{}:
		var result []interface{}
		for _, elm := range data.([]interface{}) {
			resultElm, err := manipulator.ApplyToLeafs(callback, elm)
			if err != nil {
				return nil, err
			}

			result = append(result, resultElm)
		}

		return result, nil
	default:
		//genData, err := manipulator.parseValueOfLeaf([]byte(fmt.Sprintf("%v", data)))
		//if genData == nil {
		//	return nil, err
		//}
		result, err := callback([]byte(fmt.Sprintf("%v", data)))
		if err != nil {
			return nil, err
		}

		return string(result), nil
	}
}

func (manipulator YamlManipulator) parseValueOfLeaf([]byte) ([]byte, error) {
	return []byte{}, nil
}

func pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func Generate(data []byte) ([]byte, error) {
	dataToEncrypt := pad([]byte(data))

	cipherData := make([]byte, aes.BlockSize+len(dataToEncrypt))
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	//mode := cipher.NewCBCEncrypter(cipher.Block, iv)
	//mode.CryptBlocks(cipherData[aes.BlockSize:], dataToEncrypt)

	ivSize := make([]byte, 2)
	binary.LittleEndian.PutUint16(ivSize, aes.BlockSize)

	var args []byte
	args = append(args, ivSize...)
	args = append(args, cipherData...)

	result := make([]byte, hex.EncodedLen(len(args)))
	hex.Encode(result, args)

	return result, nil
}

//func (s *AesSecret) Extract(data []byte) ([]byte, error) {
//	if len(data) == 0 {
//		return data, nil
//	}
//
//	dataToExtract, err := hexToBinary(data)
//	if err != nil {
//		return nil, err
//	}
//
//	ivLengthInfoSize := 2
//	ivSize := aes.BlockSize
//	paddingMaxSize := aes.BlockSize
//	minimalDataBinarySize := ivLengthInfoSize + ivSize + paddingMaxSize
//	minimalDataSize := minimalDataBinarySize * 2
//	if len(dataToExtract) < minimalDataBinarySize { // iv + padding
//		return nil, fmt.Errorf("minimum required data length: '%v'", minimalDataSize)
//	}
//
//	iv := dataToExtract[ivLengthInfoSize : ivLengthInfoSize+ivSize]
//	cipherText := dataToExtract[ivLengthInfoSize+ivSize:]
//
//	if len(cipherText)%aes.BlockSize != 0 {
//		return nil, fmt.Errorf("data isn't a multiple of the block size")
//	}
//
//	mode := cipher.NewCBCDecrypter(s.CipherBlock, iv)
//	mode.CryptBlocks(cipherText, cipherText)
//
//	result, err := unpad(cipherText)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func unpad(data []byte) ([]byte, error) {
//	length := len(data)
//	unpadding := int(data[length-1])
//
//	if unpadding > length {
//		return nil, fmt.Errorf("inconsistent data, unpad failed")
//	}
//
//	return data[:(length - unpadding)], nil
//}
//func hexToBinary(data []byte) ([]byte, error) {
//	result := make([]byte, hex.DecodedLen(len(data)))
//	if _, err := hex.Decode(result, data); err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func IsExtractDataError(err error) bool {
//	dataErrorPrefixs := []string{
//		"minimum required data length",
//		"encoding/hex: odd length hex string",
//	}
//
//	for _, prefix := range dataErrorPrefixs {
//		if strings.HasPrefix(err.Error(), prefix) {
//			return true
//		}
//	}
//
//	return false
//}
