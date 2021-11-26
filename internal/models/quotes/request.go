package quotes_models

type Request struct {
	Symbol     string `json:"symbol" validate:"required"`
	Function   string `json:"function" validate:"required,oneof=TIME_SERIES_DAILY TIME_SERIES_WEEKLY TIME_SERIES_MONTHLY TIME_SERIES_INTRADAY"`
	Interval   string `json:"interval"  validate:"required_if=Function TIME_SERIES_INTRADAY"`
	From       string `json:"from" validate:"omitempty,datetime=2006-01-02"`
	To         string `json:"to" validate:"omitempty,datetime=2006-01-02"`
	Date       string `json:"date" validate:"omitempty,datetime=2006-01-02"`
	OutputSize string `json:"outputsize" validate:"omitempty"`
}
