package controllers

import (
	"encoding/json"
	"net/http"
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/interface/dto"
	"strings"
)

type LinkController struct {
	logger      ports.Logger
	linkService ports.LinkService
}

func NewLinkController(logger ports.Logger, linkService ports.LinkService) *LinkController {
	return &LinkController{logger, linkService}
}

func (linkController *LinkController) AddLink(w http.ResponseWriter, r *http.Request) {
	var addLinkRequestDTO *dto.AddLinkRequestDto
	err := json.NewDecoder(r.Body).Decode(&addLinkRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	link, err := linkController.linkService.AddLink(addLinkRequestDTO.Url)
	if err != nil {
		if exceptions.IsAppError[*exceptions.BadRequestError](err) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			linkController.logger.LogError(err, "LinkController.AddLink")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (linkController *LinkController) RunWithLink(w http.ResponseWriter, r *http.Request, inner func(*entities.Link, error)) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	link, err := linkController.linkService.GetLinkByID(urlPathSplit[2])
	if err != nil {
		if !exceptions.IsAppError[*exceptions.NotFoundError](err) {
			linkController.logger.LogError(err, "LinkController.RunWithLink")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	inner(link, err)
}

func (linkController *LinkController) GetLink(w http.ResponseWriter, r *http.Request) {
	linkController.RunWithLink(w, r, func(link *entities.Link, err error) {
		if link == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(link)
	})
}

func (linkController *LinkController) RedirectToLinkUrl(w http.ResponseWriter, r *http.Request) {
	linkController.RunWithLink(w, r, func(link *entities.Link, err error) {
		if err != nil {
			http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	})
}
