package ports

import "slimlink/core/entities"

type LinkRepo interface {
	Add(link *entities.Link) error
	GetByID(id string) (*entities.Link, error)
}
