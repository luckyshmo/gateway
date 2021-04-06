package models

import (
	"time"

	"github.com/google/uuid"
)

type RawData struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Time time.Time `json:"time" db:"time"`
	Data string    `json:"data" db:"data"`
}

type Data struct {
	Id           uuid.UUID `json:"id" db:"id"`
	TimeCreated  time.Time `json:"timeCr" db:"time"`
	Time1        time.Time `json:"time1" db:"time"`
	Time2        time.Time `json:"time2" db:"time"`
	DataComposed string    `json:"data" db:"data"`
}
