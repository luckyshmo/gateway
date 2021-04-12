package service

import (
	"time"

	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
)

type StorageService struct {
	valid   repository.Storage
	invalid repository.Storage
}

func NewStorageService(valid repository.Storage, invalid repository.Storage) *StorageService {
	return &StorageService{
		valid:   valid,
		invalid: invalid,
	}
}

func (ss StorageService) WriteData(ch <-chan models.ValidPackage) error {
	for {
		select {
		case data := <-ch:
			ss.valid.WriteData(data)
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}

func (ss StorageService) WriteRawData(ch <-chan models.RawData) error {
	for {
		select {
		case data := <-ch:
			ss.invalid.WriteRawData(data)
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}
