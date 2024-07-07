package helper

import (
	"bioskuy/exception"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func CommitAndRollback(tx *sql.Tx, c *gin.Context) {

	err := recover()
	if err != nil{
		errRollback := tx.Rollback()
		if errRollback != nil {
			c.Error(exception.InternalServerError{Message: errRollback.Error()}).SetType(gin.ErrorTypePublic)
			return
		}
	}else{
		errCommit := tx.Commit()
		if errCommit != nil {
			c.Error(exception.InternalServerError{Message: errCommit.Error()}).SetType(gin.ErrorTypePublic)
			return
		}
	}
}