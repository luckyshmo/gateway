package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/source"
)

type Writer interface {
	WriteData(ch <-chan models.ValidPackage) error
	WriteRawData(ch <-chan models.RawData) error
}

type Reader interface {
	ReadData(ch chan<- models.RawData) error
}

type Process interface {
	SortData(chRaw <-chan models.RawData, chValid chan<- models.ValidPackage, chInValid chan<- models.RawData) error
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

func (services *Service) Init() {
	chRaw := make(chan models.RawData)
	chValid := make(chan models.ValidPackage)
	chInvalid := make(chan models.RawData)
	go services.Reader.ReadData(chRaw)
	go services.Process.SortData(chRaw, chValid, chInvalid) //? no need to create extrenal and interface method //todo middleware
	go services.Writer.WriteData(chValid)
	go services.Writer.WriteRawData(chInvalid)
}
