package growth_models

type GrowthSession struct {
	Date  string  `json:"date,omitempty"`
	Open  float64 `json:"open,omitempty"`
	High  float64 `json:"high,omitempty"`
	Low   float64 `json:"low,omitempty"`
	Close float64 `json:"close,omitempty"`
}

type Growth struct {
	OpenGrowth  float64 `json:"open_growth,omitempty"`
	HighGrowth  float64 `json:"high_growth,omitempty"`
	LowGrowth   float64 `json:"low_growth,omitempty"`
	CloseGrowth float64 `json:"close_growth,omitempty"`
}

type Response struct {
	Symbol string        `json:"symbol" bson:"symbol,omitempty"`
	Begin  GrowthSession `json:"begin" bson:"begin,omitempty"`
	End    GrowthSession `json:"end" bson:"end,omitempty"`
	Growth Growth        `json:"growth,omitempty"`
}

func (r *Response) Empty() bool {
	return r.Begin.Open == 0 || r.Begin.High == 0 || r.Begin.Low == 0 || r.Begin.Close == 0
}
