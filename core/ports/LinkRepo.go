package ports

import "slimlink/core/entities"

type LinkRepo interface {
	Add(*entities.Link) error
	GetByID(string) (*entities.Link, error)
}
