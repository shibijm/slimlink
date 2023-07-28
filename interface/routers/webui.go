package routers

import (
	"net/http"
	"slimlink/interface/controllers"
)

type WebUIRouter struct {
	webUIController *controllers.WebUIController
}

func NewWebUIRouter(webUIController *controllers.WebUIController) *WebUIRouter {
	return &WebUIRouter{webUIController}
}

func (rt *WebUIRouter) Route(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rt.webUIController.ServeContent(w, r)
}
