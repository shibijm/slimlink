package controllers

import (
	"fmt"
	"net/http"
	"path"
	"slimlink/services/api"
	"strings"
)

func ServeContentHandler(w http.ResponseWriter, r *http.Request) {
	cleanedPath := path.Clean(r.URL.Path)
	file, is404 := api.GetFile(cleanedPath)
	if is404 {
		urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
		if len(urlPathSplit) == 1 {
			r.URL.Path = "/api/links/" + urlPathSplit[0] + "/redirect"
			RedirectToLinkUrlHandler(w, r)
			return
		}
	}
	if file == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
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
