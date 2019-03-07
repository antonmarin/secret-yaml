package documentManipulator

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"testing"
)

//noinspection GoUnusedParameter
func returnStatic([]byte) ([]byte, error) {
	return []byte("staticValue"), nil
}
func TestYamlManipulator_ApplyToLeafs(t *testing.T) {
	manipulator := NewYamlManipulator()
	var expectedData []byte
	var err error

	simpleData := `someKey: someValue`
	expectedData = []byte("someKey: staticValue\n")

	document := make(yaml.MapSlice, 0)
	err = yaml.Unmarshal([]byte(simpleData), &document)
	if err != nil {
		t.Errorf("Should not error on valid yaml. Error: %s", err)
	}

	outputData, err := manipulator.ApplyToLeafs(returnStatic, document)
	if err != nil {
		t.Errorf("Should not error on valid yaml. Error: %s", err)
	}
	outputDataBytes, err := yaml.Marshal(outputData)
	if err != nil {
		t.Errorf("Should not error on valid yaml. Error: %s", err)
	}

	if !bytes.Equal(expectedData, outputDataBytes) {
		t.Errorf("Should apply function to leafs of yaml tree.\nExpected:\n%s\nActual:\n%s", expectedData, outputData)
	}

	//	nestedData := `---
	//someKey:
	//  nestedKey: someValue`
	//	expectedData = `---
	//someKey:
	//  nestedKey: staticValue`
	//	outputData, err = manipulator.ApplyToLeafs(returnStatic, nestedData)
	//	if err != nil {
	//		t.Errorf("Should not error on valid yaml. Error: %s", err)
	//	}
	//	if expectedData != outputData {
	//		t.Errorf("Should apply function to leafs of yaml tree.\nExpected: %s\nActual: %s", expectedData, outputData)
	//	}

	//	nestedData := `---
	//key:
	//  nestedKey: valueEncryptedWithSecretAsdf`
	//	resultOfEncrypt, err = useCase.Execute(secret, nestedData)
	//	if err != nil {
	//		t.Errorf("Should not fail if nestedData valid. Error: %s", err)
	//	}
	//	expectedData = `---
	//key:
	//  nestedKey: valueEncryptedWithSecretAsdf`
	//	if expectedData != resultOfEncrypt {
	//		t.Errorf("Should encrypt only values of nestedData.\nExpected: %s\nGot: %s", expectedData, resultOfEncrypt)
	//	}

	//	Should return error if yaml invalid

}
