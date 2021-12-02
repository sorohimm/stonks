package details_models

type DetailsRequest struct {
	Symbol   string `json:"symbol" validate:"required"`
	//Function string `json:"function" validate:"required,oneof=OVERVIEW EARNINGS INCOME_STATEMENT BALANCE_SHEET CASH_FLOW"`
	From     string `json:"from" validate:"omitempty,datetime=2006-01-02"`
	To       string `json:"to" validate:"omitempty,datetime=2006-01-02"`
	Interval string `json:"interval" validate:"omitempty"`
}
