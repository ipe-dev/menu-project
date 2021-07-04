package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
)

type SubMenuHandler interface {
	HandleBulkCreate() gin.HandlerFunc
	HandleBulkUpdate() gin.HandlerFunc
	HandleGetList() gin.HandlerFunc
	HandleGet() gin.HandlerFunc
}
type subMenuHandler struct {
	subMenuUseCase usecase.SubMenuUseCase
}

func NewSubMenuHandler(su usecase.SubMenuUseCase) SubMenuHandler {
	return subMenuHandler{subMenuUseCase: su}
}

func (h subMenuHandler) HandleGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubMenuRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err)
			return
		}
		m, e := h.subMenuUseCase.Get(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}

func (h subMenuHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateSubMenuRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		m, e := h.subMenuUseCase.BulkCreate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subMenuHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateSubMenuRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err)
		}
		m, e := h.subMenuUseCase.BulkUpdate(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subMenuHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubMenuListRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		m, e := h.subMenuUseCase.GetList(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
