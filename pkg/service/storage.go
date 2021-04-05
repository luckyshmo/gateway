package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
)

type StorageService struct {
	repo repository.Storage
}

func (ss *StorageService) WriteData(...models.RawData) error {
	return nil
}

func NewStorageService(repo repository.Storage) *StorageService {
	return &StorageService{repo: repo}
}
