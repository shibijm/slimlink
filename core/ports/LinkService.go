package ports

import "slimlink/core/entities"

type LinkService interface {
	AddLink(url string) (*entities.Link, error)
	GetLinkByID(id string) (*entities.Link, error)
}
