package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/interface/middleware"
	"github.com/ipe-dev/menu_project/usecase"
)

type UserHandler interface {
	Get(*gin.Context) gin.HandlerFunc
	Create(*gin.Context) gin.HandlerFunc
	Update(*gin.Context) gin.HandlerFunc
	Login(*gin.Context) gin.HandlerFunc
}
type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return userHandler{u}
}
func (h userHandler) Get(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
		} else {
			u, e := h.UserUseCase.Get(r)
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
			} else {
				c.JSON(http.StatusOK, *u)
			}
		}
	}
}
func (h userHandler) Create(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.CreateUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
		} else {
			e := h.UserUseCase.Create(r)
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
			} else {
				c.JSON(http.StatusOK, gin.H)
			}
		}
	}
}
func (h userHandler) Update(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.UpdateUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
		} else {
			e := h.UserUseCase.Update(r)
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
			} else {
				c.JSON(http.StatusOK, gin.H)
			}
		}
	}
}

func (h userHandler) Login(c *gin.Context) gin.HandlerFunc {
	m := middleware.NewAuthMddleware(h.UserUseCase)

	return m.LoginHandler(c)
}
