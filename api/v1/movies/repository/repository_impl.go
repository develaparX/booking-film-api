package repository

import (
	"bioskuy/api/v1/movies/entity"
	"database/sql"

	"github.com/google/uuid"
)

type movieRepositoryImpl struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepositoryImpl{db: db}
}

func (r *movieRepositoryImpl) Create(movie entity.Movie) (entity.Movie, error) {
	err := r.db.QueryRow(`INSERT INTO movies (title, description, price, duration, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		movie.Title, movie.Description, movie.Price, movie.Duration, movie.Status).Scan(&movie.ID)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}

func (r *movieRepositoryImpl) GetByID(id uuid.UUID) (entity.Movie, error) {
	var movie entity.Movie
	err := r.db.QueryRow(`SELECT id, title, description, price, duration, status FROM movies WHERE id = $1`, id).Scan(
		&movie.ID, &movie.Title, &movie.Description, &movie.Price, &movie.Duration, &movie.Status)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}

func (r *movieRepositoryImpl) Update(movie entity.Movie) (entity.Movie, error) {
	_, err := r.db.Exec(`UPDATE movies SET title = $1, description = $2, price = $3, duration = $4, status = $5 WHERE id = $6`,
		movie.Title, movie.Description, movie.Price, movie.Duration, movie.Status, movie.ID)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}

func (r *movieRepositoryImpl) Delete(id uuid.UUID) (entity.Movie, error) {
	var movie entity.Movie
	err := r.db.QueryRow(`DELETE FROM movies WHERE id = $1 RETURNING id, title, description, price, duration, status`, id).Scan(
		&movie.ID, &movie.Title, &movie.Description, &movie.Price, &movie.Duration, &movie.Status)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}
