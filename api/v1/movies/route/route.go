package route

import (
	"bioskuy/api/v1/movies/controller"
	"bioskuy/api/v1/movies/repository"
	"bioskuy/api/v1/movies/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config interface{}) {
	// Genre setup
	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepo)
	movieController := controller.NewMovieController(movieService)
	v1 := router.Group("/api/v1")
	{
		movieRoutes := v1.Group("/movies")
		{
			movieRoutes.POST("/", movieController.CreateMovie)
			movieRoutes.GET("/:id", movieController.GetMovie)
			movieRoutes.PUT("/", movieController.UpdateMovie)
			movieRoutes.DELETE("/:id", movieController.DeleteMovie)
		}
	}
}
