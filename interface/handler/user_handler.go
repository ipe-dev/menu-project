package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type UserHandler interface {
	Get() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
}
type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return userHandler{u}
}
func (h userHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
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
func (h userHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.CreateUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		} else {
			e := h.UserUseCase.Create(r)
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
			} else {
				c.JSON(http.StatusOK, gin.H{})
			}
		}
	}
}
func (h userHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.UpdateUserRequest
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		} else {
			e := h.UserUseCase.Update(r)
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error_message": e})
			} else {
				c.JSON(http.StatusOK, gin.H{})
			}
		}
	}
}
