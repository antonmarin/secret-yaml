package io

import "io/ioutil"

type file struct {
	path string
}

func (file file) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(file.path)
}
func (file file) AsString() (string, error) {
	bytes, err := ioutil.ReadFile(file.path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func New(path string) *file {
	return &file{path: path}
}
