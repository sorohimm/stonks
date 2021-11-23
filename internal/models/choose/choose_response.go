package choose_models

type ByPrice struct {
	Tickers []Price `json:"tickers"`
}

type ByPE struct {
	Tickers []PE `json:"tickers"`
}

type Price struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type PE struct {
	Symbol string  `json:"symbol"`
	PE     float64 `json:"pe"`
}
