package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
)

type StorageService struct {
	repo repository.Storage
}

func (ss *StorageService) WriteData(rawData ...models.RawData) error {
	err := ss.repo.WriteData(rawData...)
	if err != nil {
		return err
	}
	return nil
}

func NewStorageService(repo repository.Storage) *StorageService {
	return &StorageService{repo: repo}
}
