package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
)

type Storage interface {
	WriteData(...models.RawData) error
}

type Service struct {
	Storage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Storage: NewStorageService(repos.Storage),
	}
}
