package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type SubFoodStuffHandler interface {
	HandleBulkCreateSubFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleBulkUpdateSubFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleGetListSubFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleGetSubFoodStuff(c *gin.Context) gin.HandlerFunc
	HandleChangeStatus(c *gin.Context) gin.HandlerFunc
}
type subFoodStuffHandler struct {
	subFoodStuffUseCase usecase.SubFoodStuffUseCase
}

func NewSubFoodStuffHandler(u usecase.SubFoodStuffUseCase) SubFoodStuffHandler {
	return subFoodStuffHandler{subFoodStuffUseCase: u}
}

func (h subFoodStuffHandler) HandleGetSubFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.GetSubFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}

func (h subFoodStuffHandler) HandleBulkCreateSubFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.BulkCreateSubFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleBulkUpdateSubFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.BulkUpdateSubFoodStuff(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleGetListSubFoodStuff(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubFoodStuffListRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.GetSubFoodStuffList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleChangeStatus(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.ChangeSubBuyStatusRequest
		c.BindJSON(&r)
		e := h.subFoodStuffUseCase.ChangeBuyStatus(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
