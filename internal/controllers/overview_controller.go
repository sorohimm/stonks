package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/overview_interfaces"
	"stonks/internal/models/overview"
)

type OverviewControllers struct {
	Log             *zap.SugaredLogger
	OverviewService overview_interfaces.IOverviewService
	Validator       *validator.Validate
}

func (c *OverviewControllers) GetOverview(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := overview_model.Request{
		Symbol: parameters.Get("symbol"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.OverviewService.GetOverview(parameters)
	if err != nil {
		c.Log.Error("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
