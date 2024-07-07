package helper

import (
	"bioskuy/exception"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
	DriverName     string
	SecretKey string
	DurationJWT string
}

func NewConfig( c *gin.Context) *Config {
	errEnv := godotenv.Load()
	if errEnv != nil {
		c.Error(exception.InternalServerError{Message: errEnv.Error()}).SetType(gin.ErrorTypePublic)
	}
	
	return &Config{
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost: os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		DriverName: os.Getenv("DRIVER_NAME"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DurationJWT: os.Getenv("DURATION_JWT"),
	}
}