package quotes_models

type Session struct {
	Open   string `json:"1. open,omitempty"`
	High   string `json:"2. high,omitempty"`
	Low    string `json:"3. low,omitempty"`
	Close  string `json:"4. close,omitempty"`
	Volume string `json:"5. volume,omitempty"`
}

type Intraday1TS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Time Series (1min)"`
}

type Intraday5TS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Time Series (5min)"`
}

type Intraday15TS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Time Series (15min)"`
}

type Intraday30TS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Time Series (30min)"`
}

type Intraday60TS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Time Series (60min)"`
}

type MonthlyTS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Monthly Time Series,omitempty"`
}

type DailyTS struct {
	MetaData Meta               `json:"Meta Data,omitempty" `
	Series   map[string]Session `json:"Time Series (Daily),omitempty"`
}

type WeeklyTS struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Series   map[string]Session `json:"Weekly Time Series,omitempty"`
}

type Meta struct {
	Information   string `json:"1. Information,omitempty"`
	Symbol        string `json:"2. Symbol,omitempty"`
	LastRefreshed string `json:"3. Last Refreshed,omitempty"`
	TimeZone      string `json:"4. Time Zone,omitempty"`
}


