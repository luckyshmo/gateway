package service

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/source"
)

type DataSourceService struct {
	source source.DataSource
}

func (dss *DataSourceService) ReadData(ch chan<- models.RawData) error {
	dss.source.ReadData(ch)
	return nil
}

func NewDataSourceService(dataSource source.DataSource) *DataSourceService {
	return &DataSourceService{source: dataSource}
}
