package choose_models

type Request struct {
	Function       string `json:"function" validate:"required,oneof=PE PRICE FORECAST"`
	Min      string `json:"min" validate:"omitempty"`
	Max      string `json:"max" validate:"omitempty"`
	Point    string `json:"point" validate:"required_if=By price"`
	Interval string `json:"interval" validate:"omitempty,oneof= daily weekly monthly"`
}
