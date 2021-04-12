package models

import (
	"time"

	"github.com/google/uuid"
)

type RawData struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Time time.Time `json:"time" db:"time"`
	Data []byte    `json:"data" db:"data"`
}

type ValidPackage struct {
	Id          uuid.UUID `json:"id" db:"id"`
	DevEui      string    `json:"devEui" db:"devEui"`
	TimeCreated time.Time `json:"timeCr" db:"timeCr"`
	TimePackage time.Time `json:"time" db:"time"`
	Data        []byte    `json:"data" db:"data"`
	RawData     []byte    `json:"rawData" db:"rawData"`
}

func (vp *ValidPackage) isFirst() bool {
	return len(vp.Data) == 41
}
