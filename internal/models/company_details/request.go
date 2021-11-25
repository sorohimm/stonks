package details_models

type DetailsRequest struct {
	Symbol   string `json:"symbol" validate:"required"`
	Function string `json:"function" validate:"required,oneof=OVERVIEW EARNINGS INCOME_STATEMENT BALANCE_SHEET CASH_FLOW EARNINGS_CALENDAR"`
	From     string `json:"from" validate:"omitempty"`
	To       string `json:"to" validate:"omitempty"`
	Min      string `json:"min" validate:"omitempty"`
	Max      string `json:"max" validate:"omitempty"`
	Interval string `json:"interval" validate:"omitempty"`
}
