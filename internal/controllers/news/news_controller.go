package news_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"

	"stonks/internal/interfaces/news_interfaces"
	"stonks/internal/models/news"
)

type NewsControllers struct {
	Log         *zap.SugaredLogger
	NewsService news_interface.INewsService
	Validator   *validator.Validate
}

func (c *NewsControllers) GetNews(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := news_models.Request{
		Company:  parameters.Get("q"),
		SortBy:   parameters.Get("sort_by"),
		Page:     parameters.Get("page"),
		PageSize: parameters.Get("page_size"),
		From:     parameters.Get("from"),
		To:       parameters.Get("to"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("news_controller :: GetNews :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	resp, err := c.NewsService.GetNews(parameters)
	if err != nil {
		c.Log.Error("news_controller :: GetNews :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
