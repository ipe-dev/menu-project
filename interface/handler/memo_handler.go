package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase"
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
		var r usecase.GetMemoRequest
		e := c.BindJSON(&r)
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
		var r usecase.CreateMemoRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		memos, e := h.memoUsecase.Create(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, memos)
		}
	}
}
func (h memoHandler) HandleUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.UpdateMemoRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		memos, e := h.memoUsecase.Update(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, memos)
		}
	}
}
func (h memoHandler) HandleGetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r usecase.GetMemoListRequest
		e := c.BindJSON(&r)
		if e != nil {
			err := errors.NewValidateError(e, c.Request)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		memos, e := h.memoUsecase.GetList(r)
		if e != nil {
			c.Error(e).SetType(gin.ErrorTypePublic)
		} else {
			c.JSON(http.StatusOK, memos)
		}
	}
}
