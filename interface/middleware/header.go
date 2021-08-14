package middleware

import (
	"github.com/gin-gonic/gin"
)

type HeaderMiddleware interface {
	SetHeader() gin.HandlerFunc
}
type headerMiddleware struct {
}

func NewHeaderMiddleware() HeaderMiddleware {
	return headerMiddleware{}
}

func (h headerMiddleware) SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
	}
}
