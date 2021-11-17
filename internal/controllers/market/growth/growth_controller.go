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

func (g *GrowthControllers) GetQuotes(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := growth_models.Request{
		Symbol:     values.Get("symbol"),
		From:       values.Get("from"),
		To:         values.Get("to"),
	}

	if err := g.Validator.Struct(request); err != nil {
		g.Log.Info("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	resp, err := g.GrowthService.GetGrowth(values)
	if err != nil {
		g.Log.Info("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
