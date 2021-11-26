package choose_models

type Request struct {
	Function string `json:"function" validate:"required,oneof=PE PRICE FORECAST"`
	Min      string `json:"min" validate:"omitempty,numeric"`
	Max      string `json:"max" validate:"omitempty,numeric"`
	Point    string `json:"point" validate:"required_if=Function PRICE"`
	Interval string `json:"interval" validate:"omitempty,oneof= daily weekly monthly"`
}
