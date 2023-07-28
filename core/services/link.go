package services

import (
	"slimlink/core/entities"
	"slimlink/core/exceptions"
	"slimlink/core/ports"
	"slimlink/core/utils"
)

type linkService struct {
	linkRepo     ports.LinkRepo
	linkIDLength int
}

func NewLinkService(linkRepo ports.LinkRepo, linkIDLength int) ports.LinkService {
	return &linkService{linkRepo, linkIDLength}
}

func (s *linkService) CreateLink(url string) (*entities.Link, error) {
	if !utils.IsValidUrl(url) {
		return nil, exceptions.NewAppError[*exceptions.BadRequestError]("invalid URL")
	}
	var id string
	for attempt := 1; ; attempt++ {
		id = utils.GenerateBase62String(s.linkIDLength)
		_, err := s.linkRepo.GetByID(id)
		if err != nil {
			if exceptions.IsAppError[*exceptions.NotFoundError](err) {
				break
			}
			return nil, err
		}
		if attempt == 10 {
			return nil, exceptions.NewAppError[*exceptions.UnexpectedError]("maximum tries exceeded for link ID generation")
		}
	}
	link := &entities.Link{ID: id, Url: url}
	err := s.linkRepo.Add(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (svc *linkService) GetLinkByID(id string) (*entities.Link, error) {
	return svc.linkRepo.GetByID(id)
}
