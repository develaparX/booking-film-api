package service

import (
	"bioskuy/api/v1/genre/dto"
	"bioskuy/api/v1/movies/entity"
	"bioskuy/api/v1/movies/repository"

	"github.com/google/uuid"
)

type movieServiceImpl struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieServiceImpl{repo: repo}
}
func (s *movieServiceImpl) GetAllMovies(page int, size int) ([]entity.Movie, dto.Paging, error) {
	return s.repo.GetAll(page, size)
}

func (s *movieServiceImpl) CreateMovie(movie entity.Movie) (entity.Movie, error) {
	return s.repo.Create(movie)
}

func (s *movieServiceImpl) GetMovieByID(id uuid.UUID) (entity.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *movieServiceImpl) UpdateMovie(movie entity.Movie) (entity.Movie, error) {
	return s.repo.Update(movie)
}

func (s *movieServiceImpl) DeleteMovie(id uuid.UUID) (entity.Movie, error) {
	return s.repo.Delete(id)
}
