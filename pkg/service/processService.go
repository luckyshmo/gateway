package service

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/gateway/models"
)

type ProcessService struct {
}

func NewProcessService() *ProcessService {
	return &ProcessService{}
}

func fillPackage(vPack *models.ValidPackage, rawData models.RawData) {
	vPack.Id = uuid.New()
	vPack.TimeCreated = rawData.Time
	vPack.RawData = rawData.Data

	switch len(vPack.Data) {
	case 102:
		vPack.PackageType = models.First
	case 82:
		vPack.PackageType = models.Second
	default:
		vPack.PackageType = models.Over
	}
}

func (ps *ProcessService) SortData(chRaw <-chan models.RawData, chValid chan<- models.ValidPackage, chInValid chan<- models.RawData) error {
	for {
		select {
		case rawData := <-chRaw:
			var vPack models.ValidPackage
			json.Unmarshal(rawData.Data, &vPack)

			if vPack.DevEui != "" && len(vPack.Data) > 0 { //valid data

				fillPackage(&vPack, rawData)
				// tools.Validate(vPack)
				chValid <- vPack
			}

			chInValid <- rawData
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}
