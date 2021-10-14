package income_statement_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/interfaces/income_statement_interfaces"
	"stonks/internal/models/income_statement"
)

type IncomeStatementControllers struct {
	Log                    *zap.SugaredLogger
	IncomeStatementService income_statement_interfaces.IIncomeStatementService
	Validator              *validator.Validate
}

func (c *IncomeStatementControllers) GetIncomeStatement(ctx *gin.Context) {
	parameters := ctx.Request.URL.Query()

	request := income_statement_model.Request{
		Symbol: parameters.Get("symbol"),
	}

	if err := c.Validator.Struct(request); err != nil {
		c.Log.Error("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	resp, err := c.IncomeStatementService.GetIncomeStatement(parameters)
	if err != nil {
		c.Log.Error("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
