package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type SubMenuHandler interface {
	HandleBulkCreateSubMenu(c *gin.Context) gin.HandlerFunc
	HandleBulkUpdateSubMenu(c *gin.Context) gin.HandlerFunc
	HandleGetListSubMenu(c *gin.Context) gin.HandlerFunc
}
type subMenuHandler struct {
	subMenuUseCase usecase.SubMenuUseCase
}

func NewSubMenuHandler(su usecase.SubMenuUseCase) SubMenuHandler {
	return subMenuHandler{subMenuUseCase: su}
}

func (h subMenuHandler) HandleGetSubMenu(c *gin.Context) gin.HandlerFunc {
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

func (h subMenuHandler) HandleBulkCreateSubMenu(c *gin.Context) gin.HandlerFunc {
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
func (h subMenuHandler) HandleBulkUpdateSubMenu(c *gin.Context) gin.HandlerFunc {
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
func (h subMenuHandler) HandleGetListSubMenu(c *gin.Context) gin.HandlerFunc {
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
