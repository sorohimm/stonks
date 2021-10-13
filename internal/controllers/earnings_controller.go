package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/earnings_interfaces"
	"stonks/internal/models/earnings"
)

type EarningsControllers struct {
	Log             *zap.SugaredLogger
	EarningsService earnings_interfaces.IEarningsService
	Validator       *validator.Validate
}

func (c *EarningsControllers) GetEarnings(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := earnings_models.Request{
		Symbol: parameters.Get("symbol"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.EarningsService.GetEarnings(parameters)
	if err != nil {
		c.Log.Warn("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
