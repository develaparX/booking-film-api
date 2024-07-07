package route

import (
	"bioskuy/api/v1/user/controller"
	"bioskuy/api/v1/user/repository"
	"bioskuy/api/v1/user/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config interface{}) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, validate)
	userController := controller.NewUserController(userService)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", userController.CreateUser)
			users.GET("/", userController.GetAllUsers)
			users.GET("/:id", userController.GetUserByID)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}
}
