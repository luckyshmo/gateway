package models

import (
	"time"

	"github.com/google/uuid"
)

type RawData struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Time time.Time `json:"time" db:"time_cr"`
	Data []byte    `json:"data" db:"data_r"`
}

type ValidPackage struct {
	Id          uuid.UUID `json:"id" db:"id"`
	TimeCreated time.Time `json:"timeCr" db:"time_cr"`

	RawData []byte `json:"rawData" db:"raw_data"`

	AppEui    string  `json:"appEui" db:"app_eui"`
	Ack       bool    `json:"ack" db:"ack"` //!TODO 1/0 to bool
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

func (vp *ValidPackage) IsFirst() bool {
	return len(vp.Data) == 41
}
