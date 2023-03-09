package services

import (
	"slimlink/core/entities"
	"slimlink/core/ports"
)

type InfoService struct {
	pageFooterText string
}

func NewInfoService(pageFooterText string) ports.InfoService {
	return &InfoService{pageFooterText}
}

func (infoService *InfoService) GetInfo() *entities.Info {
	return &entities.Info{PageFooterText: infoService.pageFooterText}
}
