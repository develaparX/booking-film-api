package repository

import (
	"bioskuy/api/v1/genre/entity"
	"bioskuy/exception"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type genreRepositoryImpl struct {
	DB *sql.DB
	C  *gin.Context
}

func NewGenreRepository(DB *sql.DB) GenreRepository {
	return &genreRepositoryImpl{
		DB: DB,
		C:  &gin.Context{},
	}
}

func (r *genreRepositoryImpl) Create(genre entity.Genre) (entity.Genre, error) {
	genre.ID = uuid.New()
	err := r.DB.QueryRow(
		"INSERT INTO genres (id, name, created_at) VALUES ($1, $2, $3) RETURNING id, name",
		genre.ID, genre.Name, time.Now(),
	).Scan(&genre.ID, &genre.Name)
	if err != nil {
		r.C.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return entity.Genre{}, err
	}
	return genre, nil
}

func (r *genreRepositoryImpl) GetByID(id uuid.UUID) (entity.Genre, error) {
	var genre entity.Genre
	err := r.DB.QueryRow(
		"SELECT id, name FROM genres WHERE id = $1",
		id,
	).Scan(&genre.ID, &genre.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			r.C.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
			return entity.Genre{}, nil
		}
		return entity.Genre{}, err
	}
	return genre, nil
}

func (r *genreRepositoryImpl) Update(genre entity.Genre) (entity.Genre, error) {
	err := r.DB.QueryRow(
		"UPDATE genres SET name = $1, updated_at = $2 WHERE id = $3 RETURNING id, name",
		genre.Name, time.Now(), genre.ID,
	).Scan(&genre.ID, &genre.Name)
	if err != nil {
		r.C.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return entity.Genre{}, err
	}
	return genre, nil
}

func (r *genreRepositoryImpl) Delete(id uuid.UUID) (entity.Genre, error) {
	var genre entity.Genre
	err := r.DB.QueryRow(
		"DELETE FROM genres WHERE id = $1 RETURNING id, name",
		id,
	).Scan(&genre.ID, &genre.Name)
	if err != nil {
		r.C.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return entity.Genre{}, err
	}
	return genre, nil
}
