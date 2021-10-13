package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces"
	"stonks/internal/models"
)

type NewsControllers struct {
	Log         *zap.SugaredLogger
	NewsService interfaces.INewsService
	Validator   *validator.Validate
}

func (c *NewsControllers) GetNews(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := models.Request{
		Company:  parameters.Get("q"),
		SortBy:   parameters.Get("sort_by"),
		Page:     parameters.Get("page"),
		PageSize: parameters.Get("page_size"),
		From:     parameters.Get("from"),
		To:       parameters.Get("to"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.NewsService.GetNews(parameters)
	if err != nil {
		c.Log.Warn("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
