package middleware

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
)

type AuthMiddleware interface {
	NewGinJWTMiddleware() (*jwt.GinJWTMiddleware, error)
}

func NewAuthMiddleware(u usecase.UserUseCase) AuthMiddleware {
	return authMiddleware{u}
}

type authMiddleware struct {
	usecase.UserUseCase
}

func (m authMiddleware) NewGinJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(os.Getenv("SECRET")),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var r usecase.LoginRequest
			e := c.BindJSON(&r)
			if e != nil {
				return nil, errors.NewValidateError(e, c.Request)
			}
			user, err := m.UserUseCase.LoginAuthenticate(r)
			if err != nil {
				return nil, err
			}
			return *user, nil
		},
	})
	return authMiddleWare, err
}
