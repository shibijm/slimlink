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

func (webUIRouter *WebUIRouter) Route(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	webUIRouter.webUIController.ServeContent(w, r)
}
