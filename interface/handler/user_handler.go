package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
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
		id, _ := c.Get("id")
		r.ID = id.(int)
		u, e := h.UserUseCase.Get(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, *u)
		}
	}
}
func (h userHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.CreateUserRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		e = h.UserUseCase.Create(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
func (h userHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.UpdateUserRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		id, _ := c.Get("id")
		r.ID = id.(int)
		e = h.UserUseCase.Update(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
