package choose_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/choose_interfaces"
	choose_models "stonks/internal/models/choose"
	"stonks/internal/validate"
)

type ChooseControllers struct {
	Log           *zap.SugaredLogger
	ChooseService choose_interfaces.IChooseService
	Validator     *validator.Validate
}

func (c *ChooseControllers) GetChoose(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := choose_models.Request{
		By:       values.Get("by"),
		Min:      values.Get("min"),
		Max:      values.Get("max"),
		Point:    values.Get("point"),
		Interval: values.Get("interval"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Infof("invalid request: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if request.By == "price" {
		if request.Point != "open" || request.Point !=  "high" || request.Point != "low" || request.Point !=  "close" {
			c.Log.Info("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
	}

	if request.Min == "" && request.Max == "" {
		c.Log.Info("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}


	err := validate.Price(request.Min, request.Max)
	if err != nil {
		c.Log.Info("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if request.Interval == "" {
		values.Set("interval", "daily")
	}

	resp, err := c.ChooseService.GetChoose(values)
	if err != nil {
		c.Log.Info("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
