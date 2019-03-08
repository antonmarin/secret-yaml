package io

import "io/ioutil"

type File struct {
	path string
}

func (file File) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(file.path)
}
func (file File) AsString() (string, error) {
	bytes, err := ioutil.ReadFile(file.path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func NewFile(path string) *File {
	return &File{path: path}
}
