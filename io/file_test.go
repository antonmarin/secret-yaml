package io

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestNewFile(t *testing.T) {
	file := NewFile("path/to/File")
	if file == nil {
		t.Errorf("Should create File")
	}
}
func TestFile_AsBytes(t *testing.T) {
	path := filepath.Join("../testdata", "simple.yml")
	file := NewFile(path)
	expectedData := []byte(`---
some: value
`)
	actualData, err := file.AsBytes()

	if err != nil {
		t.Errorf("Should not throw error on File without error. Error: %s", err)
	}
	if !bytes.Equal(expectedData, actualData) {
		t.Errorf("Should get same data as in File")
	}
}
func TestFile_AsString(t *testing.T) {
	path := filepath.Join("..", "testdata", "simple.yml")
	file := NewFile(path)
	expectedData := `---
some: value
`
	actualData, err := file.AsString()

	if err != nil {
		t.Errorf("Should not throw error on File without error. Error: %s", err)
	}
	if expectedData != actualData {
		t.Errorf("Should get same data as in File")
	}
}
