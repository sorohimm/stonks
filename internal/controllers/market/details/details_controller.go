package details_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	details_interface "stonks/internal/interfaces/details_interfaces"
	"stonks/internal/models/company_details"
)

type CompanyDetailsControllers struct {
	Log                   *zap.SugaredLogger
	CompanyDetailsService details_interface.ICompanyDetailsService
	Validator             *validator.Validate
}

func (c *CompanyDetailsControllers) GetCompanyDetails(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := details_models.DetailsRequest{
		Symbol:   parameters.Get("symbol"),
		Function: parameters.Get("function"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.CompanyDetailsService.GetCompanyDetails(parameters)
	if err != nil {
		c.Log.Error("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}