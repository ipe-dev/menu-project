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
	menuUsecase usecase.MenuUseCase
}

func (h menuHandler) HandleGetMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuRequest
		c.BindJSON(&r)
		menu, e := h.menuUsecase.GetMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menu)
		}
	}
}

func (h menuHandler) HandleBulkCreateMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateMenuRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.BulkCreateMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleBulkUpdateMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateMenuRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.BulkUpdateMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleGetListMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuListRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.GetMenuList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
