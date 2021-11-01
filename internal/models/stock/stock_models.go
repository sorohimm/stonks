package stock_models

type Session struct {
	Open   string `json:"1. open,omitempty"`
	High   string `json:"2. high,omitempty"`
	Low    string `json:"3. low,omitempty"`
	Close  string `json:"4. close,omitempty"`
	Volume string `json:"5. volume,omitempty"`
}

type Intraday1TimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Time Series (1min)"`
}

type Intraday5TimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Time Series (5min)"`
}

type Intraday15TimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Time Series (15min)"`
}

type Intraday30TimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Time Series (30min)"`
}

type Intraday60TimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Time Series (60min)"`
}

type DailyTimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty" `
	Cell     map[string]Session `json:"Time Series (Daily),omitempty"`
}

type WeeklyTimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Weekly Time Series,omitempty"`
}

type MonthlyTimeSeries struct {
	MetaData Meta               `json:"Meta Data,omitempty"`
	Cell     map[string]Session `json:"Monthly Time Series,omitempty"`
}

type Request struct {
	Symbol     string `json:"symbol" validate:"required"`
	Function   string `json:"function" validate:"required,oneof=TIME_SERIES_DAILY TIME_SERIES_WEEKLY TIME_SERIES_MONTHLY TIME_SERIES_INTRADAY"`
	Interval   string `json:"interval"  validate:"required_if=Function TIME_SERIES_INTRADAY"`
	OutputSize string `json:"outputsize" validate:"omitempty"`
}

type Meta struct {
	Information   string `json:"1. Information,omitempty"`
	Symbol        string `json:"2. Symbol,omitempty"`
	LastRefreshed string `json:"3. Last Refreshed,omitempty"`
	OutputSize    string `json:"4. Output Size,omitempty"`
	TimeZone      string `json:"5. Time Zone,omitempty"`
}
