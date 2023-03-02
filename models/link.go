package models

type Link struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type LinkRequestDTO struct {
	Url string `json:"url"`
}
