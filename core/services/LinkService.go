package services

import (
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/core/utils"
)

type LinkService struct {
	linkRepo     ports.LinkRepo
	linkIDLength int
}

func NewLinkService(linkRepo ports.LinkRepo, linkIDLength int) ports.LinkService {
	return &LinkService{linkRepo, linkIDLength}
}

func (linkService *LinkService) AddLink(url string) (*entities.Link, error) {
	if !utils.IsValidUrl(url) {
		return nil, exceptions.NewAppError[*exceptions.BadRequestError]("invalid URL")
	}
	var id string
	for attempt := 1; ; attempt++ {
		id = utils.GenerateBase62String(linkService.linkIDLength)
		_, err := linkService.linkRepo.GetByID(id)
		if err != nil {
			if exceptions.IsAppError[*exceptions.NotFoundError](err) {
				break
			}
			return nil, err
		}
		if attempt == 10 {
			err := exceptions.NewAppError[*exceptions.UnexpectedError]("maximum tries exceeded for link ID generation")
			return nil, err
		}
	}
	link := &entities.Link{ID: id, Url: url}
	err := linkService.linkRepo.Add(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (linkService *LinkService) GetLinkByID(id string) (*entities.Link, error) {
	return linkService.linkRepo.GetByID(id)
}
