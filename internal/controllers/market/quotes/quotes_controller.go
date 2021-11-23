package quotes_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/quotes_interfaces"
	"stonks/internal/models/quotes"
	"stonks/internal/validate"
)

type QuotesControllers struct {
	Log           *zap.SugaredLogger
	QuotesService quotes_interfaces.IStockService
	Validator     *validator.Validate
}

func (c *QuotesControllers) GetQuotes(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := quotes_models.Request{
		Symbol:     values.Get("symbol"),
		Function:   values.Get("function"),
		Interval:   values.Get("interval"),
		From:       values.Get("from"),
		To:         values.Get("to"),
		Date:       values.Get("date"),
		OutputSize: values.Get("outputsize"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Info("quotes_controller :: GetQuotes :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	ok := validate.Date(request.From, request.To, request.Date)
	if !ok {
		c.Log.Info("quotes_controller :: GetQuotes :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	resp, err := c.QuotesService.GetQuotes(values)
	if err != nil {
		c.Log.Info("quotes_controller :: GetQuotes :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
