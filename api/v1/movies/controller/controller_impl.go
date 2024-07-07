package controller

import (
	"bioskuy/api/v1/movies/dto"
	"bioskuy/api/v1/movies/entity"
	"bioskuy/api/v1/movies/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type movieControllerImpl struct {
	Service service.MovieService
}

func NewMovieController(service service.MovieService) MovieController {
	return &movieControllerImpl{Service: service}
}

func (ctrl *movieControllerImpl) CreateMovie(c *gin.Context) {
	var createDTO dto.CreateMovieDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := entity.Movie{
		Title:       createDTO.Title,
		Description: createDTO.Description,
		Price:       createDTO.Price,
		Duration:    createDTO.Duration,
		Status:      createDTO.Status,
	}
	createdMovie, err := ctrl.Service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.MovieResponseDTO{ID: createdMovie.ID, Title: createdMovie.Title, Description: createdMovie.Description, Price: createdMovie.Price, Duration: createdMovie.Duration, Status: createdMovie.Status})
}

func (ctrl *movieControllerImpl) GetMovie(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	movie, err := ctrl.Service.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MovieResponseDTO{ID: movie.ID, Title: movie.Title, Description: movie.Description, Price: movie.Price, Duration: movie.Duration, Status: movie.Status})
}

func (ctrl *movieControllerImpl) UpdateMovie(c *gin.Context) {
	var updateDTO dto.UpdateMovieDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := entity.Movie{
		ID:          updateDTO.ID,
		Title:       updateDTO.Title,
		Description: updateDTO.Description,
		Price:       updateDTO.Price,
		Duration:    updateDTO.Duration,
		Status:      updateDTO.Status,
	}
	updatedMovie, err := ctrl.Service.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MovieResponseDTO{ID: updatedMovie.ID, Title: updatedMovie.Title, Description: updatedMovie.Description, Price: updatedMovie.Price, Duration: updatedMovie.Duration, Status: updatedMovie.Status})
}

func (ctrl *movieControllerImpl) DeleteMovie(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deletedMovie, err := ctrl.Service.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MovieResponseDTO{ID: deletedMovie.ID, Title: deletedMovie.Title, Description: deletedMovie.Description, Price: deletedMovie.Price, Duration: deletedMovie.Duration, Status: deletedMovie.Status})
}
