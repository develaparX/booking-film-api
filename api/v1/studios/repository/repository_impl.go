package repository

import (
	"bioskuy/api/v1/studios/entity"
	"bioskuy/exception"
	"context"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type studioRepository struct {
}

// CreateStudio implements StudiosRepository.
func (s *studioRepository) CreateStudio(ctx context.Context, tx *sql.Tx, studios entity.Studios, c *gin.Context) (entity.Studios, error) {
	query := "INSERT INTO studios (name, capacity) VALUES ($1, $2) RETURNING id"

	err := tx.QueryRowContext(ctx, query, studios.Name, studios.Capacity).Scan(&studios.Id)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return studios, err
	}
	return studios, err
}

// Delete implements StudiosRepository.
func (s *studioRepository) Delete(ctx context.Context, tx *sql.Tx, id string, c *gin.Context) error {
	query := `DELETE FROM studios WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)

	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return err
	}

	return nil
}

// FindAll implements StudiosRepository.
func (s *studioRepository) FindAll(ctx context.Context, tx *sql.Tx, c *gin.Context) ([]entity.Studios, error) {
	query := "SELECT id, name, capacity FROM studios"

	studios := []entity.Studios{}
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return studios, err
	}
	defer rows.Close()
	
	for rows.Next() {
		studio := entity.Studios{}
		if err := rows.Scan(&studio.Id, &studio.Name, &studio.Capacity); err != nil {
			return nil, err
		}
		studios = append(studios, studio)
	}
	return studios, nil
}

// FindByName implements StudiosRepository.
func (s *studioRepository) FindByName(ctx context.Context, tx *sql.Tx, name string, c *gin.Context) (entity.Studios, error) {
	query := `SELECT id, name, capacity FROM studios WHERE name = $1`

	studios := entity.Studios{}
	rows, err := tx.QueryContext(ctx, query, name)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return studios, err
	}
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(&studios.Id, &studios.Name, &studios.Capacity)
		if err != nil {
			c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
			return studios, err
		}
		return studios, nil
	} else {
		return studios , errors.New("studios not found")
	}
}

// Update implements StudiosRepository.
func (s *studioRepository) Update(ctx context.Context, tx *sql.Tx, studios entity.Studios, c *gin.Context) (entity.Studios, error) {
	query := `UPDATE studios SET capacity = $1 WHERE id = $2`
	_, err := tx.ExecContext(ctx, query, studios.Capacity, studios.Id)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return studios, err
	}
	return studios, nil
}

func NewStudiosRepository() StudiosRepository {
	return &studioRepository{}
}
