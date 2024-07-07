package service

import (
	"bioskuy/api/v1/movies/entity"

	"github.com/google/uuid"
)

type MovieService interface {
	CreateMovie(movie entity.Movie) (entity.Movie, error)
	GetMovieByID(id uuid.UUID) (entity.Movie, error)
	UpdateMovie(movie entity.Movie) (entity.Movie, error)
	DeleteMovie(id uuid.UUID) (entity.Movie, error)
}
