package aggregate_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/aggregate_interfaces"
	"stonks/internal/models/aggregate"
)

type AggregateControllers struct {
	Log              *zap.SugaredLogger
	AggregateService aggregate_interfaces.IAggregateService
	Validator        *validator.Validate
}

func (c *AggregateControllers) GetAggregate(ctx *gin.Context) {
	var request aggregate_models.Request

	err := ctx.BindJSON(&request)
	if err != nil {
		c.Log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad json :/"})
		return
	}

	resp, err := c.AggregateService.GetAggregate(request)
	if err != nil {
		c.Log.Info("aggregate_controller :: GetAggregate :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
