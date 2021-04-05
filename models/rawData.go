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
