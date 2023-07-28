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

func (c *LinkController) AddLink(w http.ResponseWriter, r *http.Request) {
	addLinkRequestDto := &dto.AddLinkRequestDto{}
	err := json.NewDecoder(r.Body).Decode(addLinkRequestDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	link, err := c.linkService.CreateLink(addLinkRequestDto.Url)
	if err != nil {
		if exceptions.IsAppError[*exceptions.BadRequestError](err) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			c.logger.LogError(err, "LinkController.AddLink")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (c *LinkController) runWithLink(w http.ResponseWriter, r *http.Request, callback func(*entities.Link, error)) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	link, err := c.linkService.GetLinkByID(urlPathSplit[2])
	if err != nil && !exceptions.IsAppError[*exceptions.NotFoundError](err) {
		c.logger.LogError(err, "LinkController.RunWithLink")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	callback(link, err)
}

func (c *LinkController) GetLink(w http.ResponseWriter, r *http.Request) {
	c.runWithLink(w, r, func(link *entities.Link, err error) {
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(link)
	})
}

func (c *LinkController) RedirectToLinkUrl(w http.ResponseWriter, r *http.Request) {
	c.runWithLink(w, r, func(link *entities.Link, err error) {
		if err != nil {
			http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	})
}
