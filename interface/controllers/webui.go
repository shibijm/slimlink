package controllers

import (
	"io"
	"net/http"
	"path"
	"slimlink/core/ports"
	"strings"
)

type WebUIController struct {
	logger     ports.Logger
	fileSystem http.FileSystem
}

func NewWebUIController(logger ports.Logger, fileSystem http.FileSystem) *WebUIController {
	return &WebUIController{logger, fileSystem}
}

func (c *WebUIController) ServeContent(w http.ResponseWriter, r *http.Request) {
	cleanedPath := path.Clean(r.URL.Path)
	file, err := c.getFile(cleanedPath)
	if err != nil {
		urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
		if len(urlPathSplit) == 1 {
			r.URL.Path = "/api/links/" + urlPathSplit[0] + "/redirect"
			return
		}
		file, err = c.getFile("/404.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		c.logger.LogError(err, "WebUIController.ServeContent")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if fileInfo.IsDir() {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	filename := fileInfo.Name()
	if filename == "404.html" {
		w.WriteHeader(http.StatusNotFound)
		io.Copy(w, file)
	} else {
		http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
	}
	// TODO: Cache-Control, ETag, Last-Modified and other headers
}

func (c *WebUIController) getFile(path string) (http.File, error) {
	file, err := c.fileSystem.Open(path)
	if err != nil {
		file2, err2 := c.fileSystem.Open(path + ".html")
		if err2 == nil {
			file, err = file2, err2
		}
	}
	return file, err
}
