package service

import (
	"encoding/json"
	"time"

	"github.com/luckyshmo/gateway/models"
)

type ProcessService struct {
}

func NewProcessService() *ProcessService {
	return &ProcessService{}
}

type Uid struct {
	Data   string
	DevEui string
}

func (ps *ProcessService) SortData(chRaw <-chan models.RawData, chValid chan<- models.Data, chInValid chan<- models.RawData) error {
	for {
		select {
		case rawData := <-chRaw:
			var p Uid

			json.Unmarshal(rawData.Data, &p)
			if /*p.DevEui != "" &&*/ len(p.Data) > 0 { //valid data
				chValid <- models.Data{
					Id:           rawData.Id,
					TimeCreated:  rawData.Time,
					DataComposed: rawData.Data,
				}
			}

			chInValid <- rawData
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}
