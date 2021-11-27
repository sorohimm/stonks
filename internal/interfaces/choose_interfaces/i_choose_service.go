package choose_interfaces

import "net/url"

type IChooseService interface {
	GetChoose(url.Values) (interface{}, error)
}
