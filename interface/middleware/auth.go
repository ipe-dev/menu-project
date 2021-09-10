package middleware

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
	"github.com/ipe-dev/menu_project/usecase/requests"
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

type User struct {
	ID   int
	Name string
}

func (m authMiddleware) NewGinJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(os.Getenv("SECRET")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"id":   v.ID,
					"name": v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return int(claims["id"].(float64))
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var r requests.LoginRequest
			e := c.BindJSON(&r)
			if e != nil {
				return nil, errors.NewValidateError(e, c.Request)
			}
			user, err := m.UserUseCase.LoginAuthenticate(r)
			if err != nil {
				return nil, err
			}
			return &User{
				ID:   user.ID,
				Name: user.Name,
			}, nil
		},
	})
	return authMiddleWare, err
}
