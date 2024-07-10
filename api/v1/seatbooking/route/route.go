package movieroute

import (
	seatRepo "bioskuy/api/v1/seat/repository"
	"bioskuy/api/v1/seatbooking/controller"
	seatbookingRepo "bioskuy/api/v1/seatbooking/repository"
	"bioskuy/api/v1/seatbooking/service"
	showtimeRepo "bioskuy/api/v1/showtime/repository"
	"bioskuy/helper"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SeatBookingRoute(router *gin.Engine, validate *validator.Validate, db *sql.DB, config *helper.Config) {

	seatbookingRepo := seatbookingRepo.NewSeatBookingRepository()
	seatRepo := seatRepo.NewSeatRepository()
	showtimeRepo := showtimeRepo.NewShowtimeRepository()

	seatBookingService := service.NewSeatBookingService(seatbookingRepo, showtimeRepo, seatRepo, validate, db)
	seatBookinngController := controller.NewSeatbookingController(seatBookingService)
	v1 := router.Group("/api/v1")
	{
		showtimeRoutes := v1.Group("/bookings")
		{
			showtimeRoutes.POST("/", seatBookinngController.Create)
			showtimeRoutes.GET("/", seatBookinngController.FindAll)
			showtimeRoutes.GET("/:seatbookingId", seatBookinngController.FindById)
			showtimeRoutes.DELETE("/:seatbookingId", seatBookinngController.Delete)
		}
	}
}
