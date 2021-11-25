package config

import "os"

type Config struct {
	ApplicationPort string
	NewsAuthData
	MarketAuthData
	DbAuthenticationData
	DetailsCollections
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

func New() *Config {
	return &Config{
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
	}
}
