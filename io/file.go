package io

import "io/ioutil"

type File struct {
	path string
}

func (file File) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(file.path)
}

func NewFile(path string) *File {
	return &File{path: path}
}
