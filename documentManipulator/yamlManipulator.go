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

type yamlManipulator struct {
}

func New() *yamlManipulator {
	return &yamlManipulator{}
}

func (manipulator yamlManipulator) ApplyToLeafs(callback func([]byte) ([]byte, error), data interface{}) (interface{}, error) {
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

func (manipulator yamlManipulator) parseValueOfLeaf([]byte) ([]byte, error) {
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

	ivSize := make([]byte, 2)
	binary.LittleEndian.PutUint16(ivSize, aes.BlockSize)

	var args []byte
	args = append(args, ivSize...)
	args = append(args, cipherData...)

	result := make([]byte, hex.EncodedLen(len(args)))
	hex.Encode(result, args)

	return result, nil
}
