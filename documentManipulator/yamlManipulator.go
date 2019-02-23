package documentManipulator

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"gopkg.in/yaml.v2"
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
		result, err := callback([]byte(fmt.Sprintf("%v", data)))
		if err != nil {
			return nil, err
		}

		return string(result), nil
	}
}

func pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func unpad(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])

	if unpadding > length {
		return nil, fmt.Errorf("inconsistent data, unpad failed")
	}

	return data[:(length - unpadding)], nil
}
