package controller

import (
	"bioskuy/api/v1/genre/dto"
	"bioskuy/api/v1/genre/entity"
	"bioskuy/api/v1/genre/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type genreControllerImpl struct {
	Service service.GenreService
}

func NewGenreController(service service.GenreService) GenreController {
	return &genreControllerImpl{Service: service}
}

func (ctrl *genreControllerImpl) CreateGenre(c *gin.Context) {
	var createDTO dto.CreateGenreDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre := entity.Genre{
		Name: createDTO.Name,
	}
	createdGenre, err := ctrl.Service.CreateGenre(genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.GenreResponseDTO{ID: createdGenre.ID, Name: createdGenre.Name})
}

func (ctrl *genreControllerImpl) GetGenre(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	genre, err := ctrl.Service.GetGenreByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.GenreResponseDTO{ID: genre.ID, Name: genre.Name})
}

func (ctrl *genreControllerImpl) UpdateGenre(c *gin.Context) {
	var updateDTO dto.UpdateGenreDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre := entity.Genre{
		ID:   updateDTO.ID,
		Name: updateDTO.Name,
	}
	updatedGenre, err := ctrl.Service.UpdateGenre(genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.GenreResponseDTO{ID: updatedGenre.ID, Name: updatedGenre.Name})
}

func (ctrl *genreControllerImpl) DeleteGenre(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deletedGenre, err := ctrl.Service.DeleteGenre(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.GenreResponseDTO{ID: deletedGenre.ID, Name: deletedGenre.Name})
}
