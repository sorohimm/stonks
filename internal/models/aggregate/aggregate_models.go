package aggregate_models

type Ticker struct {
	Symbol      string  `json:"symbol,omitempty"`
	Coefficient float64 `json:"coefficient,omitempty" json:"coefficient,omitempty"`
}

type Request struct {
	Tickers []Ticker `json:"tickers,omitempty"`
}

type Response struct {
	Tickers   []Ticker `json:"tickers,omitempty"`
	Aggregate float64   `json:"aggregate,omitempty"`
}