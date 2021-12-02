package details_interface

import (
	"net/url"
)

type ICompanyDetailsService interface {
	GetCompanyDetails(url.Values) (interface{}, error)
}
