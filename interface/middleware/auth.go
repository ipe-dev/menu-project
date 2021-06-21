package middleware

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type AuthMiddleware interface {
	NewGinJWTMiddleware(c *gin.Context) (*jwt.GinJWTMiddleware, error)
}

func NewAuthMddleware(u usecase.UserUseCase) AuthMiddleware {
	return authMiddleware{u}
}

type authMiddleware struct {
	usecase.UserUseCase
}

func (m authMiddleware) NewGinJWTMiddleware(c *gin.Context) (*jwt.GinJWTMiddleware, error) {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(os.Getenv("SECRET")),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var r usecase.LoginRequest
			c.BindJSON(&r)
			user, ok := m.UserUseCase.LoginAuthenticate(r)
			if ok {
				return *user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
	})
	return authMiddleWare, err
}
