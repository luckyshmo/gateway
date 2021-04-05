package source

import "github.com/luckyshmo/gateway/pkg/service"

type DataSource struct {
	services *service.Service
}

func NewDataSource(services *service.Service) *DataSource {
	return &DataSource{services: services}
}

func (ds *DataSource) Init() error {
	return nil
}
