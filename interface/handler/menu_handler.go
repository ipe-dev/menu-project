package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type MenuHandler interface {
	HandleGet() gin.HandlerFunc
	HandleBulkCreate() gin.HandlerFunc
	HandleBulkUpdate() gin.HandlerFunc
	HandleGetList() gin.HandlerFunc
}

func NewMenuHandler(u usecase.MenuUseCase) MenuHandler {
	return menuHandler{u}
}

type menuHandler struct {
	menuUsecase usecase.MenuUseCase
}

func (h menuHandler) HandleGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuRequest
		c.BindJSON(&r)
		menu, e := h.menuUsecase.Get(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menu)
		}
	}
}

func (h menuHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateMenuRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.BulkCreate(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateMenuRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.BulkUpdate(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuListRequest
		c.BindJSON(&r)
		menus, e := h.menuUsecase.GetList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
