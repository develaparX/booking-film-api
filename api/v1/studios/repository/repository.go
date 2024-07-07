package repository

import (
	"bioskuy/api/v1/studios/entity"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type StudiosRepository interface {
	CreateStudio(ctx context.Context, tx *sql.Tx, studios entity.Studios, c *gin.Context) (entity.Studios, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string, c *gin.Context) (entity.Studios, error)
	FindAll(ctx context.Context, tx *sql.Tx, c *gin.Context) ([]entity.Studios, error)
	Update(ctx context.Context, tx *sql.Tx, studios entity.Studios, c *gin.Context) (entity.Studios, error)
	Delete(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) error
}