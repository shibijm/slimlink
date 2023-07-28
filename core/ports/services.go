package ports

import "slimlink/core/entities"

type InfoService interface {
	GetInfo() *entities.Info
}

type LinkService interface {
	CreateLink(url string) (*entities.Link, error)
	GetLinkByID(id string) (*entities.Link, error)
}
