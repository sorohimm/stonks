package stock_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/stock_interfaces"
	stock_models "stonks/internal/models/stock"
)

type StockControllers struct {
	Log          *zap.SugaredLogger
	StockService stock_interfaces.IStockService
	Validator    *validator.Validate
}

func (c *StockControllers) GetStock(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := stock_models.Request{
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

	resp, err := c.StockService.GetStock(parameters)
	if err != nil {
		c.Log.Error("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
