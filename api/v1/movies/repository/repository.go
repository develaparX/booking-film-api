package repository

import (
	"bioskuy/api/v1/movies/entity"

	"github.com/google/uuid"
)

type MovieRepository interface {
	Create(movie entity.Movie) (entity.Movie, error)
	GetByID(id uuid.UUID) (entity.Movie, error)
	Update(movie entity.Movie) (entity.Movie, error)
	Delete(id uuid.UUID) (entity.Movie, error)
}
