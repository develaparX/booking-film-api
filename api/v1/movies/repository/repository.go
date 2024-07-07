package repository

import (
	"bioskuy/api/v1/genre/dto"
	"bioskuy/api/v1/movies/entity"

	"github.com/google/uuid"
)

type MovieRepository interface {
	GetAll(page int, size int) ([]entity.Movie, dto.Paging, error)
	Create(movie entity.Movie) (entity.Movie, error)
	GetByID(id uuid.UUID) (entity.Movie, error)
	Update(movie entity.Movie) (entity.Movie, error)
	Delete(id uuid.UUID) (entity.Movie, error)
}
