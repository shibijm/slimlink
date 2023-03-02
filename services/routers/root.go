package routers

import (
	"net/http"
	"slimlink/services/controllers"
	"strings"
)

func RootRouter(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/") {
		r.URL.Path = "/" + r.URL.Path
	}
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}
	if strings.HasSuffix(r.URL.Path, "/") {
		http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
		return
	}
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	if urlPathSplit[0] == "api" {
		ApiRouter(w, r)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	controllers.ServeContentHandler(w, r)
}
