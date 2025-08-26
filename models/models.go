package models

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Id           uuid.UUID
	CreatedDate  time.Time
	LastMoveDate time.Time
	EndDate      time.Time
	PlayerOne    uuid.UUID
	PlayerTwo    uuid.UUID
	MoveNumber   int32
}

type Player struct {
	Id         uuid.UUID
	PlayerName string
}
