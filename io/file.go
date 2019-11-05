package io

import "io/ioutil"

//File represents some file in fs
type File struct {
	path string
}

//AsBytes returns file content as byte[]
func (file File) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(file.path)
}

//AsString returns file content as string
func (file File) AsString() (string, error) {
	bytes, err := ioutil.ReadFile(file.path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//NewFile creates new File input
func NewFile(path string) *File {
	return &File{path: path}
}
