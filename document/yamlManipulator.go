package document

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
)

type yamlManipulator struct {
}

func NewYamlManipulator() *yamlManipulator {
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
		isEncryptable, err := manipulator.isValueEncryptable(data)
		if err != nil {
			return nil, err
		}
		if isEncryptable == false {
			//log incoming value can not be converted to encryptable value
			return data, nil
		}

		dataBytes, err := manipulator.castValueToBytes(data)
		if err != nil {
			return nil, err
		}

		result, err := callback(dataBytes)
		if err != nil {
			return nil, err
		}

		return string(result), nil
	}
}

func (manipulator yamlManipulator) isValueEncryptable(value interface{}) (bool, error) {
	if fmt.Sprint(reflect.TypeOf(value)) == "bool" {
		return false, nil
	}
	return true, nil
}

func (manipulator yamlManipulator) castValueToBytes(value interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", value)), nil
}
