package details_interface

import (
	"net/http"
)

type ICompanyDetailsRepo interface {
	GetCompanyDetails(*http.Request) (interface{}, error)
}
