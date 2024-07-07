package entity

import (
	"github.com/google/uuid"
)

type Genre struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
