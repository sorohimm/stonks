package growth_models

type Request struct {
	Symbol string `json:"symbol" validate:"required"`
	From   string `json:"from" validate:"required"`
	To     string `json:"to" validate:"omitempty"`
}
