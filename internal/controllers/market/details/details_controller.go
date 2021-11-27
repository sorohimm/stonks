package details_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/details_interfaces"
	"stonks/internal/models/company_details"
)

type CompanyDetailsControllers struct {
	Log                   *zap.SugaredLogger
	CompanyDetailsService details_interface.ICompanyDetailsService
	Validator             *validator.Validate
}

func (c *CompanyDetailsControllers) GetCompanyDetails(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := details_models.DetailsRequest{
		Symbol:   values.Get("symbol"),
		Function: values.Get("function"),
		From:     values.Get("from"),
		To:       values.Get("to"),
		Interval: values.Get("timing"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Info("details_controller :: GetCompanyDetails :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	resp, err := c.CompanyDetailsService.GetCompanyDetails(values)
	if err != nil {
		c.Log.Info("details_controller :: GetCompanyDetails :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
