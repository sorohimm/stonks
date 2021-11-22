package choose_models

type Request struct {
	By       string `json:"by" validate:"required,oneof=pe price forecast"`
	Min      string `json:"min" validate:"omitempty"`
	Max      string `json:"max" validate:"omitempty"`
	Point    string `json:"point" validate:"required_if=By price,oneof=open high low close"`
	Interval string `json:"interval" validate:"omitempty"`
}
