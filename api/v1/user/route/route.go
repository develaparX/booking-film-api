package route

import (
	"bioskuy/api/v1/user/controller"
	"bioskuy/api/v1/user/repository"
	"bioskuy/api/v1/user/service"
	"bioskuy/auth"
	"bioskuy/helper"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config *helper.Config) {

	authService := auth.NewService(config)
	
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, validate, db, authService)
	userController := controller.NewUserController(userService)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/google/login", userController.LoginWithGoogle)
			users.GET("/google/callback", userController.CallbackFromGoogle)
			users.GET("/:userId", userController.GetUserByID)
			users.GET("/", userController.GetAllUsers)
			users.PUT("/:userId", userController.UpdateUser)
			users.DELETE("/:userId", userController.DeleteUser)
		}
	}
}