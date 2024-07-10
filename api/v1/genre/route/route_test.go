package genreroute

import (
	"bioskuy/helper"
	"database/sql"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TestGenreRoute(t *testing.T) {
	type args struct {
		router   *gin.Engine
		validate *validator.Validate
		db       *sql.DB
		config   *helper.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenreRoute(tt.args.router, tt.args.validate, tt.args.db, tt.args.config)
		})
	}
}
