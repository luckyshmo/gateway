package repository

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
)

type Storage interface {
	WriteData(...sensor.Sensor) error
	WriteRawData(...models.RawData) error
}

type Repository struct {
	Storage
}

func NewRepository(someStorage Storage) *Repository {
	return &Repository{
		Storage: someStorage,
	}
}
