package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
	process "github.com/luckyshmo/gateway/pkg/processors"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/source"
)

type Writer interface {
	WriteData(ch <-chan sensor.Sensor) error
	WriteRawData(ch <-chan models.RawData) error
}

type Reader interface {
	ReadData(ch chan<- models.RawData) error
}

type ProcessData interface {
	SortData(chRaw <-chan models.RawData, chValid chan<- sensor.Sensor, chInValid chan<- models.RawData) error
}

type Service struct {
	Writer
	Reader
	ProcessData
}

func NewService(
	valid repository.Storage,
	invalid repository.Storage,
	dataSource *source.DataSourceObj,
	pService *process.ProcessService,
) *Service {
	return &Service{
		Writer:      NewStorageService(valid, invalid),
		Reader:      NewDataSourceService(dataSource.DataSource),
		ProcessData: *pService,
	}
}

func (services *Service) Init() {
	chRaw := make(chan models.RawData)
	chValid := make(chan sensor.Sensor)
	chInvalid := make(chan models.RawData)
	go services.Reader.ReadData(chRaw)
	go services.ProcessData.SortData(chRaw, chValid, chInvalid) //todo middleware
	go services.Writer.WriteData(chValid)
	go services.Writer.WriteRawData(chInvalid)
}
