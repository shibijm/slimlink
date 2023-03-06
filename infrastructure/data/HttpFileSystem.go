package data

import (
	"net/http"
	"slimlink/core/ports"
)

type HttpFileSystem struct {
	fileSystem http.FileSystem
}

func NewHttpFileSystem(fileSystem http.FileSystem) ports.HttpFileSystem {
	return &HttpFileSystem{fileSystem}
}

func (httpFileSystem *HttpFileSystem) GetFile(path string) (http.File, error) {
	file, err := httpFileSystem.fileSystem.Open(path)
	if err != nil {
		file2, err2 := httpFileSystem.fileSystem.Open(path + ".html")
		if err2 == nil {
			file, err = file2, err2
		}
	}
	return file, err
}
