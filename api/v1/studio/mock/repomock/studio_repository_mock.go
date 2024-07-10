package repomock

import (
	"bioskuy/api/v1/studio/entity"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type StudioRepositoryMock struct {
	mock.Mock
}

func (m *StudioRepositoryMock) Save(ctx context.Context, tx *sql.Tx, user entity.Studio, c *gin.Context) (entity.Studio, error) {
	args := m.Called(ctx, tx, user, c)
	return args.Get(0).(entity.Studio), args.Error(1)
}

func (m *StudioRepositoryMock) FindByID(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) (entity.Studio, error) {
	args := m.Called(ctx, tx, id, c)
	return args.Get(0).(entity.Studio), args.Error(1)
}

func (m *StudioRepositoryMock) Update(ctx context.Context, tx *sql.Tx, studio entity.Studio, c *gin.Context) (entity.Studio, error) {
	args := m.Called(ctx, tx, studio, c)
	return args.Get(0).(entity.Studio), args.Error(1)
}

func (m *StudioRepositoryMock) FindAll(ctx context.Context, id string, tx *sql.Tx, c *gin.Context) ([]entity.Studio, error) {
	args := m.Called(ctx, id, tx, c)
	return args.Get(0).([]entity.Studio), args.Error(1)
}

func (m *StudioRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) error {
	args := m.Called(ctx, tx, id, c)
	return args.Error(0)
}
