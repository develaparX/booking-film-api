package genreroute

import (
	"bioskuy/api/v1/genre/controller"
	"bioskuy/api/v1/genre/repository"
	"bioskuy/api/v1/genre/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GenreRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config interface{}) {
	// Genre setup
	genreRepo := repository.NewGenreRepository(db)
	genreService := service.NewGenreService(genreRepo)
	genreController := controller.NewGenreController(genreService)
	v1 := router.Group("/api/v1")
	{
		genre := v1.Group("/genre")
		{
			genre.GET("/", genreController.GetAll)
			genre.POST("/", genreController.CreateGenre)
			genre.GET("/:id", genreController.GetGenre)
			genre.PUT("/:id", genreController.UpdateGenre)
			genre.DELETE("/:id", genreController.DeleteGenre)
		}
	}
}
