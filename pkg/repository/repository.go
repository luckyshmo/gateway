package repository

import "github.com/luckyshmo/gateway/models"

type Storage interface {
	WriteData(...models.Data) error
}

type Repository struct {
	Storage
}

func NewRepository(someStorage Storage) *Repository {
	return &Repository{
		Storage: someStorage,
	}
}
