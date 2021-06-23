package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.BindJSON(&r)
		m, e := h.subMenuUseCase.GetSubMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}

func (h subMenuHandler) HandleBulkCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkCreateSubMenuRequest
		c.BindJSON(&r)
		m, e := h.subMenuUseCase.BulkCreateSubMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subMenuHandler) HandleBulkUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.BulkUpdateSubMenuRequest
		c.BindJSON(&r)
		m, e := h.subMenuUseCase.BulkUpdateSubMenu(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
func (h subMenuHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetSubMenuListRequest
		c.BindJSON(&r)
		m, e := h.subMenuUseCase.GetSubMenuList(r)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, m)
		}
	}
}
