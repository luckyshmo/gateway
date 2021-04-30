package process

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
)

type Process interface {
	SortData(chRaw <-chan models.RawData, chValid chan<- sensor.Sensor, chInValid chan<- models.RawData) error
}

type ProcessService struct {
	Process
}

func NewProcessService(somePService Process) *ProcessService {
	return &ProcessService{
		Process: somePService,
	}
}
