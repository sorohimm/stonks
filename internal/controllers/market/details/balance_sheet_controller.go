package details_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stonks/internal/models/company_details"
)

func (c *CompanyDetailsControllers) GetBalanceSheet(ctx *gin.Context) {
	values := ctx.Request.URL.Query()

	request := details_models.DetailsRequest{
		Symbol:   values.Get("symbol"),
		From:     values.Get("from"),
		To:       values.Get("to"),
		Interval: values.Get("timing"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Info("details_controller :: GetCompanyDetails :: validation :: invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request :/"})
		return
	}

	values.Set("function", "BALANCE_SHEET")
	resp, err := c.CompanyDetailsService.GetCompanyDetails(values)
	if err != nil {
		c.Log.Info("details_controller :: GetCompanyDetails :: unknown server error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
