package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/error"
	"github.com/ipe-dev/menu_project/usecase"
)

type FoodStuffHandler interface {
	HandleBulkCreate() gin.HandlerFunc
	HandleBulkUpdate() gin.HandlerFunc
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
		var r usecase.GetFoodStuffRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
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
		var r usecase.BulkCreateFoodStuffRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		f, e := h.foodStuffUseCase.BulkCreate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateFoodStuffRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		f, e := h.foodStuffUseCase.BulkUpdate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetFoodStuffListRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
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
		var r usecase.ChangeFoodStuffStatusRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
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
