package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type MenuHandler interface {
	HandleGetMenu(c *gin.Context) gin.HandlerFunc
	HandleBulkCreateMenu(c *gin.Context) gin.HandlerFunc
	HandleBulkUpdateMenu(c *gin.Context) gin.HandlerFunc
	HandleGetListMenu(c *gin.Context) gin.HandlerFunc
}
type menuHandler struct {
	mu usecase.MenuUseCase
}

func (mh menuHandler) HandleGetMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuRequest
		c.BindJSON(&r)
		m, e := mh.mu.GetMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}

func (mh menuHandler) HandleBulkCreateMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateMenuRequest
		c.BindJSON(&r)
		m, e := mh.mu.BulkCreateMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (mh menuHandler) HandleBulkUpdateMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateMenuRequest
		c.BindJSON(&r)
		m, e := mh.mu.BulkUpdateMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (mh menuHandler) HandleGetListMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuListRequest
		c.BindJSON(&r)
		ms, e := mh.mu.GetMenuList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, ms)
		}
	}
}
