package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/source"
)

type Writer interface {
	WriteData(ch <-chan models.Data) error
	WriteRawData(ch <-chan models.RawData) error
}

type Reader interface {
	ReadData(ch chan<- models.RawData) error
}

type Process interface {
	SortData(chRaw <-chan models.RawData, chValid chan<- models.Data, chInValid chan<- models.RawData) error
}
type Service struct {
	Writer
	Reader
	Process
}

func NewService(valid *repository.Repository, invalid *repository.Repository, dataSource *source.DataSourceObj) *Service {
	return &Service{
		Writer:  NewStorageService(valid.Storage, invalid),
		Reader:  NewDataSourceService(dataSource.DataSource),
		Process: NewProcessService(),
	}
}
