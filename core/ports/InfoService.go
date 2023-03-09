package ports

import "slimlink/core/entities"

type InfoService interface {
	GetInfo() *entities.Info
}
