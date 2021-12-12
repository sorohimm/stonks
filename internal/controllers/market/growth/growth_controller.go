package growth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/growth_interfaces"
	"stonks/internal/models/growth"
)

type GrowthControllers struct {
	Log           *zap.SugaredLogger
	GrowthService growth_interfaces.IGrowthService
	Validator     *validator.Validate
}

func (c *GrowthControllers) GetGrowth(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := growth_models.Request{
		Symbol:     values.Get("symbol"),
		From:       values.Get("from"),
		To:         values.Get("to"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Info("growth_controller :: GetGrowth :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	resp, err := c.GrowthService.GetGrowth(values)
	if err != nil {
		c.Log.Infof("growth_controller :: GetGrowth :: server :: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
