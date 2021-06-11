package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type FoodStuffHandler interface {
	HandleBulkCreateFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleBulkUpdateFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleGetListFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleGetFoodStuff(c *gin.Context) gin.HandlerFunc
}
type foodStuffHandler struct {
	foodStuffUseCase usecase.FoodStuffUseCase
}

func NewFoodStuffHandler(u usecase.FoodStuffUseCase) FoodStuffHandler {
	return foodStuffHandler{foodStuffUseCase: u}
}

func (h foodStuffHandler) HandleGetFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetFoodStuffRequest
		c.BindJSON(&r)
		f, e := h.foodStuffUseCase.GetFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}

func (h foodStuffHandler) HandleBulkCreateFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateFoodStuffRequest
		c.BindJSON(&r)
		f, e := h.foodStuffUseCase.BulkCreateFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) HandleBulkUpdateFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateFoodStuffRequest
		c.BindJSON(&r)
		f, e := h.foodStuffUseCase.BulkUpdateFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) HandleGetListFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetFoodStuffListRequest
		c.BindJSON(&r)
		f, e := h.foodStuffUseCase.GetFoodStuffList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}
func (h foodStuffHandler) ChangeBuyStatus(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.ChangeFoodStuffStatusRequest
		c.BindJSON(&r)
		e := h.foodStuffUseCase.ChangeBuyStatus(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
