package repository

import (
	"bioskuy/api/v1/genretomovie/entity"
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GenreToMovieRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    GenreToMovieRepository
}

var mockingGenreToMovie = entity.GenreToMovie{
	ID:      uuid.New().String(),
	GenreID: uuid.New().String(),
	MovieID: uuid.New().String(),
}

func (suite *GenreToMovieRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewGenreToMovieRepository()
}

func TestGenreToMovieRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GenreToMovieRepositoryTestSuite))
}

func (suite *GenreToMovieRepositoryTestSuite) TestSave_Success() {
	ctx := context.Background()
	c := gin.Context{}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`INSERT INTO genre_to_movies \(genre_id, movie_id\) VALUES \(\$1, \$2\) RETURNING id`).
		WithArgs(mockingGenreToMovie.GenreID, mockingGenreToMovie.MovieID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockingGenreToMovie.ID))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.Save(ctx, tx, mockingGenreToMovie, &c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingGenreToMovie.ID, result.ID)
}

func (suite *GenreToMovieRepositoryTestSuite) TestSave_Failed() {
	ctx := context.Background()
	c := gin.Context{}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`INSERT INTO genre_to_movies \(genre_id, movie_id\) VALUES \(\$1, \$2\) RETURNING id`).
		WithArgs(mockingGenreToMovie.GenreID, mockingGenreToMovie.MovieID).
		WillReturnError(errors.New("Insert GenreToMovie Failed"))
	suite.mockSql.ExpectRollback()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Save(ctx, tx, mockingGenreToMovie, &c)
	assert.Error(suite.T(), err)
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindByID_Success() {
	ctx := context.Background()
	c := gin.Context{}
	id := mockingGenreToMovie.ID

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "genre_id", "movie_id"}).
			AddRow(mockingGenreToMovie.ID, mockingGenreToMovie.GenreID, mockingGenreToMovie.MovieID))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindByID(ctx, tx, id, &c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingGenreToMovie, result)
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindByID_NoRows() {
	ctx := context.Background()
	c := gin.Context{}
	id := uuid.New().String()

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "genre_id", "movie_id"})) // No rows
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindByID(ctx, tx, id, &c)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.GenreToMovie{}, result)
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindByID_QueryContextError() {
	ctx := context.Background()
	c := gin.Context{}
	id := uuid.New().String()

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnError(errors.New("QueryContext error"))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindByID(ctx, tx, id, &c)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.GenreToMovie{}, result)
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindByID_NotFound() {
	ctx := context.Background()
	c := gin.Context{}
	id := uuid.New().String()

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindByID(ctx, tx, id, &c)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.GenreToMovie{}, result)
}
func (suite *GenreToMovieRepositoryTestSuite) TestFindAll_Success() {
	ctx := context.Background()
	c := gin.Context{}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "genre_id", "movie_id"}).
			AddRow(mockingGenreToMovie.ID, mockingGenreToMovie.GenreID, mockingGenreToMovie.MovieID))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindAll(ctx, tx, &c)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 1)
	assert.Equal(suite.T(), mockingGenreToMovie, result[0])
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindAll_NoRows() {
	ctx := context.Background()
	c := gin.Context{}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "genre_id", "movie_id"})) // No rows
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	result, err := suite.repo.FindAll(ctx, tx, &c)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 0)
}

func (suite *GenreToMovieRepositoryTestSuite) TestFindAll_Failed() {
	ctx := context.Background()
	c := gin.Context{}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectQuery(`SELECT id, genre_id, movie_id FROM genre_to_movies`).
		WillReturnError(errors.New("FindAll Failed"))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.FindAll(ctx, tx, &c)
	assert.Error(suite.T(), err)
}

func (suite *GenreToMovieRepositoryTestSuite) TestDelete_Success() {
	ctx := context.Background()
	c := gin.Context{}
	id := mockingGenreToMovie.ID

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec(`DELETE FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	err = suite.repo.Delete(ctx, tx, id, &c)
	assert.NoError(suite.T(), err)
}

func (suite *GenreToMovieRepositoryTestSuite) TestDelete_Failed() {
	ctx := context.Background()
	c := gin.Context{}
	id := uuid.New().String()

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec(`DELETE FROM genre_to_movies WHERE id = \$1`).
		WithArgs(id).
		WillReturnError(errors.New("Delete Failed"))
	suite.mockSql.ExpectRollback()

	tx, err := suite.mockDb.BeginTx(ctx, nil)
	assert.NoError(suite.T(), err)

	err = suite.repo.Delete(ctx, tx, id, &c)
	assert.Error(suite.T(), err)
}
