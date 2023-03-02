package api

import (
	"errors"
	"slimlink/models"
	"slimlink/services/data"
	"slimlink/utils"
)

var linkIDLength int

func InitLinkApi(linkIDLength_ int) {
	linkIDLength = linkIDLength_
}

func AddLink(linkRequestDTO *models.LinkRequestDTO) (*models.Link, error) {
	if !utils.IsValidUrl(linkRequestDTO.Url) {
		return nil, errors.New("invalid URL")
	}
	var id string
	var url string
	var err error
	for tries := 0; tries < 10; tries++ {
		id = utils.GenerateBase62String(linkIDLength)
		url, err = data.RedisGet(id)
		if url == "" && err == nil {
			break
		}
	}
	if url != "" {
		return nil, errors.New("maximum tries exceeded for link ID generation")
	}
	err = data.RedisSet(id, linkRequestDTO.Url)
	if err != nil {
		return nil, err
	}
	return &models.Link{Id: id, Url: linkRequestDTO.Url}, nil
}

func GetLinkByID(id string) (*models.Link, error) {
	url, err := data.RedisGet(id)
	if err != nil {
		return nil, errors.New("failed to fetch data from Redis")
	}
	if url == "" {
		return nil, errors.New("link not found")
	}
	return &models.Link{Id: id, Url: url}, nil
}
