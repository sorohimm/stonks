package growth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/growth_interfaces"
	"stonks/internal/models/growth"
	"stonks/internal/validate"
)

type GrowthControllers struct {
	Log           *zap.SugaredLogger
	GrowthService growth_interfaces.IGrowthService
	Validator     *validator.Validate
}

func (c *GrowthControllers) GetQuotes(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := growth_models.Request{
		Symbol:     values.Get("symbol"),
		From:       values.Get("from"),
		To:         values.Get("to"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Info("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	ok := validate.Date(request.From, request.To)
	if !ok {
		c.Log.Info("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	resp, err := c.GrowthService.GetGrowth(values)
	if err != nil {
		c.Log.Info("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
