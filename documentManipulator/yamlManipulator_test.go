package documentManipulator

import "testing"

//noinspection GoUnusedParameter
func returnStatic([]byte) ([]byte, error) {
	return []byte("staticValue"), nil
}
func TestYamlManipulator_ApplyToLeafs(t *testing.T) {
	manipulator := NewYamlManipulator()
	var expectedData interface{}
	var outputData interface{}
	var err error

	simpleData := `---
someKey: someValue`
	expectedData = `---
someKey: staticValue`
	outputData, err = manipulator.ApplyToLeafs(returnStatic, simpleData)
	if err != nil {
		t.Errorf("Should not error on valid yaml. Error: %s", err)
	}
	if expectedData != outputData {
		t.Errorf("Should apply function to leafs of yaml tree.\nExpected: %s\nActual: %s", expectedData, outputData)
	}

	nestedData := `---
someKey: 
  nestedKey: someValue`
	expectedData = `---
someKey: 
  nestedKey: staticValue`
	outputData, err = manipulator.ApplyToLeafs(returnStatic, nestedData)
	if err != nil {
		t.Errorf("Should not error on valid yaml. Error: %s", err)
	}
	if expectedData != outputData {
		t.Errorf("Should apply function to leafs of yaml tree.\nExpected: %s\nActual: %s", expectedData, outputData)
	}
}
