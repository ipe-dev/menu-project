package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMiddleware interface {
	ErrorHandle() gin.HandlerFunc
}
type errorMiddleware struct {
}

func NewErrorMiddleware() ErrorMiddleware {
	return errorMiddleware{}
}

func (e errorMiddleware) ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Err,
			})
		}
	}
}
