package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/usecase"
)

type MenuHandler interface {
	HandleGetMenu(c *gin.Context) gin.HandlerFunc
	HandleCreateMenu(c *gin.Context) gin.HandlerFunc
	HandleUpdateMenu(c *gin.Context) gin.HandlerFunc
	HandleDeleteMenu(c *gin.Context) gin.HandlerFunc
	HandleGetListMenu(c *gin.Context) gin.HandlerFunc
}
type menuHandler struct {
	mu usecase.MenuUseCase
}

func (mh menuHandler) HandleCreateMenu(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := mh.mu.CreateMenu(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
