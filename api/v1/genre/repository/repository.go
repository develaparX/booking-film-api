package repository

import (
	"bioskuy/api/v1/genre/entity"

	"github.com/google/uuid"
)

type GenreRepository interface {
	Create(genre entity.Genre) (entity.Genre, error)
	GetByID(id uuid.UUID) (entity.Genre, error)
	Update(genre entity.Genre) (entity.Genre, error)
	Delete(id uuid.UUID) (entity.Genre, error)
}
