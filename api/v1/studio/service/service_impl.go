package service

import (
	"bioskuy/api/v1/studio/dto"
	"bioskuy/api/v1/studio/entity"
	"bioskuy/api/v1/studio/repository"
	"bioskuy/exception"
	"bioskuy/helper"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	Repo     repository.StudioRepository
	Validate *validator.Validate
	DB *sql.DB
}


func NewStudioService(repo repository.StudioRepository, validate *validator.Validate, DB *sql.DB) StudioService {
	return &userService{Repo: repo, Validate: validate, DB: DB}
}

func (s *userService) Create(ctx context.Context, request dto.CreateStudioRequest, c *gin.Context) (dto.StudioResponse, error) {
	var StudioResponse = dto.StudioResponse{}

	err := s.Validate.Struct(request)
	if err != nil {
		c.Error(exception.ValidationError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return StudioResponse, err
	}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return StudioResponse, err
	}
	defer helper.CommitAndRollback(tx, c)

	studio := entity.Studio{
		Name:  request.Name,
		Capacity: request.Capacity,
	}

	result, err := s.Repo.Save(ctx, tx, studio, c)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return StudioResponse, err
	}

	StudioResponse.ID = result.ID
	StudioResponse.Name = result.Name
	StudioResponse.Capacity = result.Capacity

	return StudioResponse, nil
}

func (s *userService) FindByID(ctx context.Context, id string, c *gin.Context) (dto.StudioResponse, error){

	StudioResponse := dto.StudioResponse{}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}
	defer helper.CommitAndRollback(tx, c)

	result, err := s.Repo.FindByID(ctx, tx, id, c)
	if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}

	StudioResponse.ID = result.ID
	StudioResponse.Name = result.Name
	StudioResponse.Capacity = result.Capacity

	return StudioResponse, nil
}

func (s *userService) FindAll(ctx context.Context, c *gin.Context) ([]dto.StudioResponse, error){
	StudioResponses := []dto.StudioResponse{}

	tx, err := s.DB.Begin()
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponses, err
	}
	defer helper.CommitAndRollback(tx, c)

	result, err := s.Repo.FindAll(ctx, tx, c)
	if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponses, err
	}

	for _, studio := range result {
		StudioResponse := dto.StudioResponse{}

		StudioResponse.ID = studio.ID
		StudioResponse.Name = studio.Name
		StudioResponse.Capacity = studio.Capacity

		StudioResponses = append(StudioResponses, StudioResponse)
		
	}

	return StudioResponses, nil
}

func (s *userService) Update(ctx context.Context, request dto.UpdateStudioRequest, c *gin.Context) (dto.StudioResponse, error){
	StudioResponse := dto.StudioResponse{}
    var studio entity.Studio

    err := s.Validate.Struct(request)
    if err != nil {
		c.Error(exception.ValidationError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}

    resultStudio, err := s.FindByID(ctx, request.ID, c)
    if err != nil {
		c.Error(exception.NotFoundError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}

    tx, err := s.DB.Begin()
    if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}
    defer helper.CommitAndRollback(tx, c)

    studio.ID = resultStudio.ID
	studio.Name = resultStudio.Name
	studio.Capacity = resultStudio.Capacity

	if request.Name != "" {
		studio.Name = request.Name
	}

    result, err := s.Repo.Update(ctx, tx, studio, c)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return  StudioResponse, err
	}

    StudioResponse.ID = result.ID
	StudioResponse.Name = resultStudio.Name
	StudioResponse.Capacity = resultStudio.Capacity

    return StudioResponse, nil
}

func (s *userService) Delete(ctx context.Context, id string, c *gin.Context) error{
	studio := entity.Studio{}

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

	studio.ID = resultUser.ID

	err = s.Repo.Delete(ctx, tx, id, c)
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return err
	}

	return nil
}
