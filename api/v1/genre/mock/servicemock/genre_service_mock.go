package servicemock

import (
	"bioskuy/api/v1/genre/dto"
	"bioskuy/api/v1/genre/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGenreService struct {
	mock.Mock
}

func (m *MockGenreService) CreateGenre(genre entity.Genre) (entity.Genre, error) {
	args := m.Called(genre)
	return args.Get(0).(entity.Genre), args.Error(1)
}

func (m *MockGenreService) GetGenreByID(id uuid.UUID) (entity.Genre, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Genre), args.Error(1)
}

func (m *MockGenreService) GetAll(page int, size int) ([]entity.Genre, dto.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.Genre), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockGenreService) UpdateGenre(genre entity.Genre) (entity.Genre, error) {
	args := m.Called(genre)
	return args.Get(0).(entity.Genre), args.Error(1)
}

func (m *MockGenreService) DeleteGenre(id uuid.UUID) (entity.Genre, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Genre), args.Error(1)
}
