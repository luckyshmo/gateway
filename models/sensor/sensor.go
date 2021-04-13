package sensor

import (
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/gateway/models"
)

type PackageType int

const (
	FirstPressure PackageType = iota
	SecondPressure
	ConcreteTemp
	Over
)

func (pt PackageType) String() string {
	return [...]string{"FirstPressure", "SecondPressure", "ConcreteTemp", "Over"}[pt]
}

type Sensor struct {
	Id          uuid.UUID `json:"id" db:"id"`
	TimeCreated time.Time `json:"timeCr" db:"time_cr"`

	RawData []byte `json:"rawData" db:"raw_data"`

	PackageType PackageType `json:"packageType" db:"package_type"`

	AppEui    string  `json:"appEui" db:"app_eui"`
	Ack       bool    `json:"ack" db:"ack"` //TODO 1/0 to bool
	Data      string  `json:"data" db:"data_f"`
	Dr        string  `json:"dr" db:"dr"`
	Fcnt      int     `json:"fcnt" db:"fcnt"`
	Freq      int     `json:"freq" db:"freq"`
	GatewayId string  `json:"gatewayId" db:"gateway_id"`
	Port      int     `json:"port" db:"port"`
	Rssi      int     `json:"rssi" db:"rssi"`
	Snr       float64 `json:"snr" db:"snr"` //? 32/64
	TimeStamp int64   `json:"ts" db:"time_stamp_"`
	Type      string  `json:"type" db:"type_"`
	DevEui    string  `json:"devEui" db:"dev_eui"`
}

func (vPack *Sensor) FillPackage(rawData models.RawData) {
	vPack.Id = rawData.Id
	vPack.TimeCreated = rawData.Time
	vPack.RawData = rawData.Data

	switch len(vPack.Data) {
	case 102:
		vPack.PackageType = FirstPressure
	case 82:
		vPack.PackageType = SecondPressure
	case 26:
		vPack.PackageType = ConcreteTemp
	default:
		vPack.PackageType = Over
	}
}
