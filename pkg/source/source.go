package source

import (
	"github.com/luckyshmo/gateway/models"
)

type DataSource interface {
	ReadData(ch chan<- models.RawData) error
}

type DataSourceObj struct {
	DataSource
}

func NewDataSource(ds DataSource) *DataSourceObj {
	return &DataSourceObj{
		DataSource: ds,
	}
}
