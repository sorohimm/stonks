package choose_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/choose_interfaces"
	choose_models "stonks/internal/models/choose"
)

type ChooseControllers struct {
	Log           *zap.SugaredLogger
	ChooseService choose_interfaces.IChooseService
	Validator     *validator.Validate
}

func (c *ChooseControllers) GetChoose(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := choose_models.Request{
		Function: values.Get("function"),
		Min:      values.Get("min"),
		Max:      values.Get("max"),
		Point:    values.Get("point"),
		Interval: values.Get("interval"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Infof("choose_controller :: GetChoose :: validation :: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	if request.Function == "PRICE" {
		if request.Point != "open" || request.Point != "high" || request.Point != "low" || request.Point != "close" {
			c.Log.Info("choose_controller :: GetChoose :: validation :: invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
			return
		}
	}

	if request.Min == "" && request.Max == "" {
		c.Log.Info("choose_controller :: GetChoose :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	if request.Interval == "" {
		values.Set("interval", "daily")
	}

	resp, err := c.ChooseService.GetChoose(values)
	if err != nil {
		c.Log.Info("choose_controller :: GetChoose :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
