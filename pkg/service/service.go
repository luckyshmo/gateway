package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/source"
)

type Writer interface {
	WriteData(ch <-chan models.Data) error
}

type Reader interface {
	ReadData(ch chan<- models.RawData) error
}

type Process interface {
	ProcessData(chRaw <-chan models.RawData, chData chan<- models.Data) error
}
type Service struct {
	Writer
	Reader
	Process
}

func NewService(repos *repository.Repository, dataSource *source.DataSourceObj) *Service {
	return &Service{
		Writer:  NewStorageService(repos.Storage),
		Reader:  NewSourceService(dataSource.DataSource),
		Process: NewProcessService(),
	}
}
