package service

import (
	"bioskuy/api/v1/genre/dto"
	"bioskuy/api/v1/movies/entity"

	"github.com/google/uuid"
)

type MovieService interface {
	GetAllMovies(page int, size int) ([]entity.Movie, dto.Paging, error)
	CreateMovie(movie entity.Movie) (entity.Movie, error)
	GetMovieByID(id uuid.UUID) (entity.Movie, error)
	UpdateMovie(movie entity.Movie) (entity.Movie, error)
	DeleteMovie(id uuid.UUID) (entity.Movie, error)
}
