package config

import (
	"os"
	"time"
)

type Config struct {
	CurrentTime     time.Time
	ApplicationPort string
	NewsAuthData
	MarketAuthData
	DbAuthenticationData
	DetailsCollections
	QuotesCollection
}

type NewsAuthData struct {
	NewsKey string
}

type MarketAuthData struct {
	MarketKey string
}

type DbAuthenticationData struct {
	DbUsername string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	DbURI      string
}

type DetailsCollections struct {
	Overview        string
	Earnings        string
	BalanceSheet    string
	IncomeStatement string
	CashFlow        string
}

type QuotesCollection struct {
	Daily      string
	Weekly     string
	Monthly    string
	Intraday1  string
	Intraday5  string
	Intraday15 string
	Intraday30 string
	Intraday60 string
}

func New() *Config {
	return &Config{
		CurrentTime:     time.Now(),
		ApplicationPort: os.Getenv("APPLICATION_PORT"),
		NewsAuthData: NewsAuthData{
			NewsKey: os.Getenv("NEWS_KEY"),
		},
		MarketAuthData: MarketAuthData{
			MarketKey: os.Getenv("MARKET_KEY"),
		},
		DbAuthenticationData: DbAuthenticationData{
			DbUsername: os.Getenv("MONGODB_USER"),
			DbPassword: os.Getenv("MONGODB_PASSWORD"),
			DbHost:     os.Getenv("MONGODB_HOST"),
			DbPort:     os.Getenv("MONGODB_PORT"),
			DbName:     os.Getenv("MONGODB_DATABASE"),
			DbURI:      os.Getenv("MONGODB_URI"),
		},
		DetailsCollections: DetailsCollections{
			Overview:        os.Getenv("OVERVIEW_COLLECTION"),
			Earnings:        os.Getenv("EARNINGS_COLLECTION"),
			BalanceSheet:    os.Getenv("CASH_FLOW_COLLECTION"),
			IncomeStatement: os.Getenv("BALANCE_SHEET_COLLECTION"),
			CashFlow:        os.Getenv("INCOME_STATEMENT_COLLECTION"),
		},
		QuotesCollection: QuotesCollection{
			Daily: os.Getenv("TIME_SERIES_DAILY"),
			Weekly: os.Getenv("TIME_SERIES_WEEKLY"),
			Monthly: os.Getenv("TIME_SERIES_MONTHLY"),

			Intraday1: os.Getenv("ONE_MIN"),
			Intraday5: os.Getenv("FIVE_MIN"),
			Intraday15: os.Getenv("FIFTEEN_MIN"),
			Intraday30: os.Getenv("THIRTY_MIN"),
			Intraday60: os.Getenv("SIXTY_MIN"),
		},
	}
}
