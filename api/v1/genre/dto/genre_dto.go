package dto

import "github.com/google/uuid"

type CreateGenreDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateGenreDTO struct {
	ID   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}

type GenreResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
