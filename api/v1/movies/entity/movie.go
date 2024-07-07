package entity

import "github.com/google/uuid"

type Movie struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Duration    int       `json:"duration"`
	Status      string    `json:"status"`
}
