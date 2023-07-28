package services

import (
	"slimlink/core/entities"
	"slimlink/core/ports"
)

type infoService struct {
	pageFooterText string
}

func NewInfoService(pageFooterText string) ports.InfoService {
	return &infoService{pageFooterText}
}

func (s *infoService) GetInfo() *entities.Info {
	return &entities.Info{PageFooterText: s.pageFooterText}
}
