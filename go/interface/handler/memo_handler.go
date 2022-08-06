package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type MemoHandler interface {
	HandleGet() gin.HandlerFunc
	HandleCreate() gin.HandlerFunc
	HandleUpdate() gin.HandlerFunc
	HandleGetList() gin.HandlerFunc
}

func NewMemoHandler(u usecase.MemoUseCase) MemoHandler {
	return memoHandler{u}
}

type memoHandler struct {
	memoUsecase usecase.MemoUseCase
}

func (h memoHandler) HandleGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.GetMemoRequest
		e := c.BindJSON(&r)
		userID, _ := c.Get("id")
		r.UserID = userID.(int)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		memo, e := h.memoUsecase.Get(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, memo)
		}
	}
}

func (h memoHandler) HandleCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.CreateMemoRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		e = h.memoUsecase.Create(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
func (h memoHandler) HandleUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.UpdateMemoRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		e = h.memoUsecase.Update(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
func (h memoHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r requests.GetMemoListRequest
		id, _ := c.Get("id")
		r.UserID = id.(int)
		memos, e := h.memoUsecase.GetList(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, memos)
		}
	}
}
