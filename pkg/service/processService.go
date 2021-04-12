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

type UData struct {
	Ts     int64
	Data   string
	DevEui string
}

func (ps *ProcessService) SortData(chRaw <-chan models.RawData, chValid chan<- models.ValidPackage, chInValid chan<- models.RawData) error {
	for {
		select {
		case rawData := <-chRaw:
			var uData UData
			json.Unmarshal(rawData.Data, &uData)
			if uData.DevEui != "" && len(uData.Data) > 0 { //valid data
				chValid <- getValidPackage(rawData, uData)
			}

			chInValid <- rawData
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}

//valid package could contain 0 Time
func getValidPackage(rawData models.RawData, uData UData) models.ValidPackage {
	return models.ValidPackage{
		Id:          rawData.Id,
		DevEui:      uData.DevEui,
		TimeCreated: rawData.Time,
		TimePackage: time.Unix(uData.Ts/1000, 0),
		Data:        []byte(uData.Data),
		RawData:     rawData.Data,
	}
}
