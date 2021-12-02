package details_controller

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/interfaces/details_interfaces"
)

type CompanyDetailsControllers struct {
	Log                   *zap.SugaredLogger
	CompanyDetailsService details_interface.ICompanyDetailsService
	Validator             *validator.Validate
}
