package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/error"
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
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		menu, e := h.menuUsecase.Get(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, menu)
		}
	}
}

func (h menuHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateMenuRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		menus, e := h.menuUsecase.BulkCreate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateMenuRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		menus, e := h.menuUsecase.BulkUpdate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
func (h menuHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMenuListRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := error.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		menus, e := h.menuUsecase.GetList(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, menus)
		}
	}
}
