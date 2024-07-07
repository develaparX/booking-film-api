package route

import (
	"bioskuy/api/v1/studio/controller"
	"bioskuy/api/v1/studio/repository"
	"bioskuy/api/v1/studio/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func StudioRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB) {
	
	studioRepo := repository.NewStudioRepository()
	studioService := service.NewStudioService(studioRepo, validate, db)
	studioController := controller.NewStudioController(studioService)

	v1 := router.Group("/api/v1")
	{
		studios := v1.Group("/studios")
		{
			studios.POST("/", studioController.Create)
			studios.GET("/:studioId", studioController.FindById)
			studios.GET("/", studioController.FindAll)
			studios.PUT("/:studioId", studioController.Update)
			studios.DELETE("/:studioId", studioController.Delete)
		}
	}
}
