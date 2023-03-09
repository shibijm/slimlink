package controllers

import (
	"encoding/json"
	"net/http"
	"slimlink/core/ports"
)

type InfoController struct {
	infoService ports.InfoService
}

func NewInfoController(infoService ports.InfoService) *InfoController {
	return &InfoController{infoService}
}

func (infoController *InfoController) GetInfo(w http.ResponseWriter, r *http.Request) {
	info := infoController.infoService.GetInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
