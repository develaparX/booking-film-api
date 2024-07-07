package main

import (
	genreroute "bioskuy/api/v1/genre/genreRoute"
	movieroute "bioskuy/api/v1/movies/movieRoute"
	"bioskuy/api/v1/user/route"
	"bioskuy/app"
	"bioskuy/exception"
	"bioskuy/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	var c *gin.Context
	router := gin.Default()
	validate := validator.New()
	config := helper.NewConfig(c)
	db := app.GetConnection(config)
	defer db.Close()

	router.Use(exception.ErrorHandler)

	route.UserRoute(router, validate, db, config)
	genreroute.GenreRoute(router, validate, db, config)
	movieroute.MovieRoute(router, validate, db, config)

	err := router.Run(":3000")
	if err != nil {
		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
		return
	}
}
