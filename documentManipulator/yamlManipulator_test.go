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
func TestYamlManipulator_ApplyToLeafs_ShouldApplyOnlyToLeafsOfMapItem(t *testing.T) {
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
}

func TestYamlManipulator_ApplyToLeafs_ShouldApplyOnlyToLeafsOfMapSlice(t *testing.T) {
	manipulator := NewYamlManipulator()
	nestedData := `---
someKey:
  nestedKey: someValue`
	expectedData := []byte("someKey:\n  nestedKey: staticValue\n")

	document := make(yaml.MapSlice, 0)
	err := yaml.Unmarshal([]byte(nestedData), &document)
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
}
