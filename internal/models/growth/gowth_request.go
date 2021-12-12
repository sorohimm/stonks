package growth_models

type Request struct {
	Symbol string `json:"symbol" validate:"required"`
	From   string `json:"from" validate:"required,datetime=2006-01-02"`
	To     string `json:"to" validate:"omitempty,datetime=2006-01-02"`
}
