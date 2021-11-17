package growth_models

type Meta struct {
	Symbol string `json:"symbol" bson:"symbol"`
}

type GrowthSession struct {
	Date  string `json:"date,omitempty"`
	Open  string `json:"open,omitempty"`
	High  string `json:"high,omitempty"`
	Low   string `json:"low,omitempty"`
	Close string `json:"close,omitempty"`
}

type Growth struct {
	OpenGrowth  string `json:"open_growth,omitempty"`
	HighGrowth  string `json:"high_growth,omitempty"`
	LowGrowth   string `json:"low_growth,omitempty"`
	CloseGrowth string `json:"close_growth,omitempty"`
}

type Response struct {
	Meta   Meta          `json:"_meta" bson:"_meta,omitempty"`
	Begin  GrowthSession `json:"begin"`
	End    GrowthSession `json:"end"`
	Growth Growth        `json:"growth"`
}
