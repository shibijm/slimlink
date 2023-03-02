package routers

import (
	"net/http"
	"slimlink/services/controllers"
	"strings"
)

func ApiRouter(w http.ResponseWriter, r *http.Request) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	n := len(urlPathSplit)
	if n > 1 && urlPathSplit[1] == "links" {
		if n == 2 {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			controllers.AddLinkHandler(w, r)
			return
		}
		if n == 3 {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			controllers.GetLinkHandler(w, r)
			return
		}
		if n == 4 && urlPathSplit[3] == "redirect" {
			if r.Method != http.MethodGet && r.Method != http.MethodHead {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			controllers.RedirectToLinkUrlHandler(w, r)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
