package choose_models

type Request struct {
	By     string `json:"by" validate:"required,oneof=pe price forecast"`
	Min   string `json:"min" validate:"omitempty"`
	Max     string `json:"max" validate:"omitempty"`
}
