package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slimlink/models"
	"slimlink/services/api"
	"strings"
)

func AddLinkHandler(w http.ResponseWriter, r *http.Request) {
	var linkRequestDTO models.LinkRequestDTO
	err := json.NewDecoder(r.Body).Decode(&linkRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	link, err := api.AddLink(&linkRequestDTO)
	if err != nil {
		if err.Error() == "invalid URL" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	link, err := api.GetLinkByID(urlPathSplit[2])
	if err != nil {
		if err.Error() == "link not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func RedirectToLinkUrlHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSplit := strings.Split(r.URL.Path, "/")[1:]
	link, err := api.GetLinkByID(urlPathSplit[2])
	if err != nil {
		if err.Error() == "link not found" {
			http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
}
