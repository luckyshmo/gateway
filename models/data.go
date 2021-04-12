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
	DevEui      string    `json:"devEui" db:"dev_eui"`
	TimeCreated time.Time `json:"timeCr" db:"time_cr"`
	TimePackage time.Time `json:"time" db:"time_p"`
	Data        []byte    `json:"data" db:"data_f"`
	RawData     []byte    `json:"rawData" db:"raw_data"`
}

func (vp *ValidPackage) isFirst() bool {
	return len(vp.Data) == 41
}
