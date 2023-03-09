package routers

import (
	"net/http"
	"slimlink/interface/controllers"
	"strings"
)

type ApiRouter struct {
	infoController *controllers.InfoController
	linkController *controllers.LinkController
}

func NewApiRouter(infoController *controllers.InfoController, linkController *controllers.LinkController) *ApiRouter {
	return &ApiRouter{infoController, linkController}
}

func (apiRouter *ApiRouter) Route(w http.ResponseWriter, r *http.Request) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	n := len(urlPathSplit)
	if n > 1 {
		if urlPathSplit[1] == "links" {
			if n == 2 {
				if r.Method != http.MethodPost {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				apiRouter.linkController.AddLink(w, r)
				return
			}
			if n == 3 {
				if r.Method != http.MethodGet {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				apiRouter.linkController.GetLink(w, r)
				return
			}
			if n == 4 && urlPathSplit[3] == "redirect" {
				if r.Method != http.MethodGet {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				apiRouter.linkController.RedirectToLinkUrl(w, r)
				return
			}
		}
		if urlPathSplit[1] == "info" {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			apiRouter.infoController.GetInfo(w, r)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
