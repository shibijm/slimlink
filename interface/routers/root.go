package routers

import (
	"net/http"
	"strings"
)

type RootRouter struct {
	webUIRouter *WebUIRouter
	apiRouter   *ApiRouter
}

func NewRootRouter(webUIRouter *WebUIRouter, apiRouter *ApiRouter) *RootRouter {
	return &RootRouter{webUIRouter, apiRouter}
}

func (rt *RootRouter) Route(w http.ResponseWriter, r *http.Request) {
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
	urlPreRouting := r.URL.Path
	if urlPathSplit[0] == "api" {
		rt.apiRouter.Route(w, r)
	} else {
		rt.webUIRouter.Route(w, r)
	}
	urlPostRouting := r.URL.Path
	if urlPreRouting != urlPostRouting {
		rt.Route(w, r)
	}
}
