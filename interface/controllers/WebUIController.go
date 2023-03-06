package controllers

import (
	"net/http"
	"path"
	"slimlink/core/ports"
	"strings"
)

type WebUIController struct {
	logger         ports.Logger
	httpFileSystem ports.HttpFileSystem
}

func NewWebUIController(logger ports.Logger, httpFileSystem ports.HttpFileSystem) *WebUIController {
	return &WebUIController{logger, httpFileSystem}
}

func (webUIController *WebUIController) ServeContent(w http.ResponseWriter, r *http.Request) {
	cleanedPath := path.Clean(r.URL.Path)
	file, err := webUIController.httpFileSystem.GetFile(cleanedPath)
	if err != nil {
		urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
		if len(urlPathSplit) == 1 {
			r.URL.Path = "/api/links/" + urlPathSplit[0] + "/redirect"
			return
		}
		file, err = webUIController.httpFileSystem.GetFile("/404.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	fileInfo, err := file.Stat()
	if err != nil {
		webUIController.logger.LogError(err, "WebUIController.ServeContent")
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
	}
	http.ServeContent(w, r, filename, fileInfo.ModTime(), file) // TODO: Fix superfluous response.WriteHeader call
	file.Close()
}
