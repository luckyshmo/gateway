package process

import (
	"encoding/json"
	"time"

	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
)

type LuaProcessService struct {
}

func NewLuaProcessService() *JsonProcessService {
	return &JsonProcessService{}
}

func (ps *LuaProcessService) SortData(chRaw <-chan models.RawData, chValid chan<- sensor.Sensor, chInValid chan<- models.RawData) error {
	for {
		select {
		case rawData := <-chRaw:
			var vPack sensor.Sensor
			json.Unmarshal(rawData.Data, &vPack)

			if len(vPack.Data) > 0 { //valid data
				vPack.FillPackage(rawData)
				chValid <- vPack
				continue
			}

			chInValid <- rawData
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}
