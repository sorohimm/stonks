package details_models

type DetailsRequest struct {
	Symbol   string `json:"symbol" validate:"required"`
	Function string `json:"function" validate:"required,oneof=OVERVIEW EARNINGS INCOME_STATEMENT BALANCE_SHEET CASH_FLOW EARNINGS_CALENDAR"`
	From     string `json:"from" validate:"omitempty"`
	To       string `json:"to" validate:"omitempty"`
	Timing   string `json:"timing" validate:"omitempty"`
}
