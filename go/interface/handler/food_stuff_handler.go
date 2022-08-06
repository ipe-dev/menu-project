package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type FoodStuffHandler interface {
	HandleBulkCreate() gin.HandlerFunc
	HandleGetList() gin.HandlerFunc
	HandleGet() gin.HandlerFunc
	HandleChangeBuyStatus() gin.HandlerFunc
}
type foodStuffHandler struct {
	foodStuffUseCase usecase.FoodStuffUseCase
}

func NewFoodStuffHandler(u usecase.FoodStuffUseCase) FoodStuffHandler {
	return foodStuffHandler{foodStuffUseCase: u}
}

func (h foodStuffHandler) HandleGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.GetFoodStuffRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		f, e := h.foodStuffUseCase.Get(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, f)
		}

	}
}

func (h foodStuffHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.BulkCreateFoodStuffRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		e = h.foodStuffUseCase.BulkCreate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

func (h foodStuffHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.GetFoodStuffListRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		f, e := h.foodStuffUseCase.GetList(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) HandleChangeBuyStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.ChangeFoodStuffStatusRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		e = h.foodStuffUseCase.ChangeBuyStatus(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
