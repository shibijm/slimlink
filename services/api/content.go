package api

import (
	"net/http"
)

var fileSystem http.FileSystem

func InitContentApi(fileSystem_ http.FileSystem) {
	fileSystem = fileSystem_
}

func GetFile(path string) (http.File, bool) {
	is404 := false
	file, err := fileSystem.Open(path)
	if err != nil {
		file, err = fileSystem.Open(path + ".html")
		if err != nil {
			is404 = true
			file, err = fileSystem.Open("/404.html")
			if err != nil {
				return nil, is404
			}
		}
	}
	return file, is404
}
