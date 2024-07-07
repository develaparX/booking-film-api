package service

import (
	"bioskuy/api/v1/genretomovie/dto"
	"bioskuy/api/v1/genretomovie/entity"
	"bioskuy/api/v1/genretomovie/repository"
	"bioskuy/exception"
	"bioskuy/helper"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type genretomovieServiceImpl struct {
	Repo repository.GenreToMovieRepository
	Validate *validator.Validate
	DB *sql.DB
}

func NewGenreToMovieService(repo repository.GenreToMovieRepository, validate *validator.Validate, DB *sql.DB) GenreToMovieService {
	return &genretomovieServiceImpl{
		Repo: repo,
		Validate: validate,
		DB: DB,
	}
}

func (s *genretomovieServiceImpl) Create(ctx context.Context, request dto.CreateGenreToMovieRequest, c *gin.Context) (dto.GenreToMovieResponse, error) {
	var GenretomovieResponse = dto.GenreToMovieResponse{}

	err := s.Validate.Struct(request)
	if err != nil {
		c.Error(exception.ValidationError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return GenretomovieResponse, err
	}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return GenretomovieResponse, err
	}
	defer helper.CommitAndRollback(tx, c)

	genretomovie := entity.GenreToMovie{
		GenreID: request.GenreID ,
		MovieID: request.MovieID,
		
	}

	result, err := s.Repo.Save(ctx, tx, genretomovie, c)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return GenretomovieResponse, err
	}

	GenretomovieResponse.ID = result.ID
	GenretomovieResponse.GenreID = result.GenreID
	GenretomovieResponse.MovieID = result.MovieID

	return GenretomovieResponse, nil
}

func (s *genretomovieServiceImpl) FindByID(ctx context.Context, id string, c *gin.Context) (dto.GenreToMovieResponse, error){

	GenreToMovieResponse := dto.GenreToMovieResponse{}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  GenreToMovieResponse, err
	}
	defer helper.CommitAndRollback(tx, c)

	result, err := s.Repo.FindByID(ctx, tx, id, c)
	if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  GenreToMovieResponse, err
	}

	GenreToMovieResponse.ID = result.ID
	GenreToMovieResponse.GenreID = result.GenreID
	GenreToMovieResponse.MovieID = result.MovieID

	return GenreToMovieResponse, nil
}

func (s *genretomovieServiceImpl) FindAll(ctx context.Context, c *gin.Context) ([]dto.GenreToMovieResponse, error){
	GenreToMovieResponses := []dto.GenreToMovieResponse{}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  GenreToMovieResponses, err
	}
	defer helper.CommitAndRollback(tx, c)

	result, err := s.Repo.FindAll(ctx, tx, c)
	if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  GenreToMovieResponses, err
	}

	for _, genretomovie := range result {
		GenretomovieResponse := dto.GenreToMovieResponse{}

		GenretomovieResponse.ID = genretomovie.ID
		GenretomovieResponse.GenreID = genretomovie.GenreID
		GenretomovieResponse.MovieID = genretomovie.MovieID

		GenreToMovieResponses = append(GenreToMovieResponses, GenretomovieResponse)
		
	}

	return GenreToMovieResponses, nil
}

func (s *genretomovieServiceImpl) Delete(ctx context.Context, id string, c *gin.Context) error{
	genretomovie := entity.GenreToMovie{}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return err
	}
	defer helper.CommitAndRollback(tx, c)

	resultUser, err := s.Repo.FindByID(ctx, tx, id, c)
	if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return err
	}

	genretomovie.ID = resultUser.ID

	err = s.Repo.Delete(ctx, tx, id, c)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return err
	}

	return nil
}
