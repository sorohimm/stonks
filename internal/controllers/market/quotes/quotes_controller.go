package quotes_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/quotes_interfaces"
	"stonks/internal/models/quotes"
)

type QuotesControllers struct {
	Log           *zap.SugaredLogger
	QuotesService quotes_interfaces.IStockService
	Validator     *validator.Validate
}

func (c *QuotesControllers) GetQuotes(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := quotes_models.Request{
		Symbol:     parameters.Get("symbol"),
		Function:   parameters.Get("function"),
		Interval:   parameters.Get("interval"),
		OutputSize: parameters.Get("outputsize"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.QuotesService.GetQuotes(parameters)
	if err != nil {
		c.Log.Error("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
