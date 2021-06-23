package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type SubFoodStuffHandler interface {
	HandleBulkCreate() gin.HandlerFunc
	HandleBulkUpdate() gin.HandlerFunc
	HandleGetList() gin.HandlerFunc
	HandleGet() gin.HandlerFunc
	HandleChangeStatus() gin.HandlerFunc
}
type subFoodStuffHandler struct {
	subFoodStuffUseCase usecase.SubFoodStuffUseCase
}

func NewSubFoodStuffHandler(u usecase.SubFoodStuffUseCase) SubFoodStuffHandler {
	return subFoodStuffHandler{subFoodStuffUseCase: u}
}

func (h subFoodStuffHandler) HandleGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.Get(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}

func (h subFoodStuffHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.BulkCreate(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateSubFoodStuffRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.BulkUpdate(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubFoodStuffListRequest
		c.BindJSON(&r)
		m, e := h.subFoodStuffUseCase.GetList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subFoodStuffHandler) HandleChangeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.ChangeSubBuyStatusRequest
		c.BindJSON(&r)
		e := h.subFoodStuffUseCase.ChangeStatus(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
