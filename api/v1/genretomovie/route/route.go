package genreroute

import (
	"bioskuy/api/v1/genretomovie/controller"
	"bioskuy/api/v1/genretomovie/repository"
	"bioskuy/api/v1/genretomovie/service"
	"bioskuy/helper"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GenreToMovieRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config *helper.Config) {
	genretomovieRepo := repository.NewGenreToMovieRepository()
	genretomovieService := service.NewGenreToMovieService(genretomovieRepo, validate, db)
	genretomovieController := controller.NewGenreToMovieController(genretomovieService)
	v1 := router.Group("/api/v1")
	{
		genretomovie := v1.Group("/genretomovie")
		{
			genretomovie.POST("/", genretomovieController.Create)
			genretomovie.GET("/", genretomovieController.FindAll)
			genretomovie.GET("/:genretomovieId", genretomovieController.FindById)
			genretomovie.DELETE("/:genretomovieId", genretomovieController.Delete)
		}
	}
}