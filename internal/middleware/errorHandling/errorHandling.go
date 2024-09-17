package errorHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrDB                  = fmt.Errorf("database error")
	ErrConnectDb           = fmt.Errorf("can't connect to database")
)

// обрабатываем кастомные ошибки
func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch err.Err {
		case ErrInternalServerError:
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError.Error()})
			return
		case ErrDB:
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrDB.Error()})
			return
		default:
			c.JSON(418, gin.H{"error": "I'm a teapot"})
			return
		}
	}
}
