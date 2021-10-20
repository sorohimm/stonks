package details_models

// Overview function
type Overview struct {
	Symbol               string `json:"Symbol,omitempty"`
	AssetType            string `json:"AssetType,omitempty"`
	Name                 string `json:"Name,omitempty"`
	Description          string `json:"Description,omitempty"`
	Exchange             string `json:"Exchange,omitempty"`
	Currency             string `json:"Currency,omitempty"`
	Country              string `json:"Country,omitempty"`
	Sector               string `json:"Sector,omitempty"`
	Industry             string `json:"Industry,omitempty"`
	FiscalYearEnd        string `json:"FiscalYearEnd,omitempty"`
	DividendPerShare     string `json:"DividendPerShare,omitempty"`
	ProfitMargin         string `json:"ProfitMargin,omitempty"`
	MarketCapitalization string `json:"MarketCapitalization,omitempty"`
	BookValue            string `json:"BookValue,omitempty"`
	WeekHigh52           string `json:"52WeekHigh,omitempty"`
	WeekLow52            string `json:"52WeeLow,omitempty"`
	DayMovingAverage50   string `json:"50DayMovingAverage,omitempty"`
	DayMovingAverage200  string `json:"200DayMovingAverage,omitempty"`
	PayoutRatio          string `json:"PayoutRatio,omitempty"`
	DividendDate         string `json:"DividendDate,omitempty"`
}

// IncomeStatement function

type IncomeStatementReports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                  string `json:"reportedCurrency,omitempty"`
	GrossProfit                       string `json:"grossProfit,omitempty"`
	TotalRevenue                      string `json:"totalRevenue,omitempty"`
	CostOfRevenue                     string `json:"costOfRevenue,omitempty"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold,omitempty"`
	OperatingIncome                   string `json:"operatingIncome,omitempty"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative,omitempty"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment,omitempty"`
	OperatingExpenses                 string `json:"operatingExpenses,omitempty"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet,omitempty"`
	NetInterestIncome                 string `json:"netInterestIncome,omitempty"`
	InterestIncome                    string `json:"interestIncome,omitempty"`
	InterestExpense                   string `json:"interestExpense,omitempty"`
	NonInterestIncome                 string `json:"nonInterestIncome,omitempty"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome,omitempty"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization,omitempty"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax,omitempty"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense,omitempty"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense,omitempty"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations,omitempty"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax,omitempty"`
	Ebit                              string `json:"ebit,omitempty"`
	Ebitda                            string `json:"ebitda,omitempty"`
	NetIncome                         string `json:"net_income,omitempty"`
}

type IncomeStatement struct {
	Symbol           string                   `json:"symbol"`
	AnnualReports    []IncomeStatementReports `json:"annualReports"`
	QuarterlyReports []IncomeStatementReports `json:"quarterlyReports"`
}

// Earnings

type AnnualEarnings struct {
	FiscalDateEnding string `json:"fiscalDateEnding,omitempty"`
	ReportedEPS      string `json:"reportedEPS,omitempty"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string `json:"fiscalDateEnding,omitempty"`
	ReportedDate       string `json:"reportedDate,omitempty"`
	ReportedEPS        string `json:"reportedEPS,omitempty"`
	EstimatedEPS       string `json:"estimatedEPS,omitempty"`
	Surprise           string `json:"surprise,omitempty"`
	SurprisePercentage string `json:"surprisePercentage,omitempty"`
}

type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

// Balance sheet

type BalanceSheetReports struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                       string `json:"reportedCurrency,omitempty"`
	TotalAssets                            string `json:"totalAssets,omitempty"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets,omitempty"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue,omitempty"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments,omitempty"`
	Inventory                              string `json:"inventory,omitempty"`
	CurrentNetReceivables                  string `json:"currentNetReceivables,omitempty"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets,omitempty"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment,omitempty"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE,omitempty"`
	IntangibleAssets                       string `json:"intangibleAssets,omitempty"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill,omitempty"`
	Goodwill                               string `json:"goodwill,omitempty"`
	Investments                            string `json:"investments,omitempty"`
	LongTermInvestments                    string `json:"longTermInvestments,omitempty"`
	ShortTermInvestments                   string `json:"shortTermInvestments,omitempty"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets,omitempty"`
	OtherNonCurrrentAssets                 string `json:"otherNonCurrrentAssets,omitempty"`
	TotalLiabilities                       string `json:"totalLiabilities,omitempty"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities,omitempty"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable,omitempty"`
	DeferredRevenue                        string `json:"deferredRevenue,omitempty"`
	CurrentDebt                            string `json:"currentDebt,omitempty"`
	ShortTermDebt                          string `json:"shortTermDebt,omitempty"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities,omitempty"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations,omitempty"`
	LongTermDebt                           string `json:"longTermDebt,omitempty"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt,omitempty"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent,omitempty"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal,omitempty"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities,omitempty"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities,omitempty"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity,omitempty"`
	TreasuryStock                          string `json:"treasuryStock,omitempty"`
	RetainedEarnings                       string `json:"retainedEarnings,omitempty"`
	CommonStock                            string `json:"commonStock,omitempty"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding,omitempty"`
}

type BalanceSheet struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []BalanceSheetReports `json:"annualReports"`
	QuarterlyReports []BalanceSheetReports `json:"quarterlyReports"`
}

// Cash Flow

type CashFlow struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []BalanceSheetReports `json:"annualReports"`
	QuarterlyReports []BalanceSheetReports `json:"quarterlyReports"`
}

// DetailsRequest

type DetailsRequest struct {
	Symbol   string `json:"symbol" validate:"required"`
	Function string `json:"function" validate:"required,oneof=OVERVIEW EARNINGS INCOME_STATEMENT BALANCE_SHEET CASH_FLOW EARNINGS_CALENDAR"`
}
