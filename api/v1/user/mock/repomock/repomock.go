package repomock

import (
	"bioskuy/api/v1/user/entity"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (r *UserRepository) Save(ctx context.Context, tx *sql.Tx, user entity.User, c *gin.Context) (entity.User, error) {
	args := r.Called(ctx, tx, user, c)
	return args.Get(0).(entity.User), args.Error(1)
}

func (r *UserRepository) FindByEmail(ctx context.Context, tx *sql.Tx, email string, c *gin.Context) (entity.User, error) {
	args := r.Called(ctx, tx, email, c)
	return args.Get(0).(entity.User), args.Error(1)
}

func (r *UserRepository) FindByID(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) (entity.User, error) {
	args := r.Called(ctx, tx, id, c)
	return args.Get(0).(entity.User), args.Error(1)
}

func (r *UserRepository) FindAll(ctx context.Context, tx *sql.Tx, c *gin.Context) ([]entity.User, error) {
	args := r.Called(ctx, tx, c)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (r *UserRepository) UpdateToken(ctx context.Context, tx *sql.Tx, user entity.User, c *gin.Context) (entity.User, error) {
	args := r.Called(ctx, tx, user, c)
	return args.Get(0).(entity.User), args.Error(1)
}

func (r *UserRepository) Update(ctx context.Context, tx *sql.Tx, user entity.User, c *gin.Context) (entity.User, error) {
	args := r.Called(ctx, tx, user, c)
	return args.Get(0).(entity.User), args.Error(1)
}

func (r *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) error {
	args := r.Called(ctx, tx, id, c)
	return args.Error(0)
}
