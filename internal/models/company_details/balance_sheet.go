package details_models

import "strconv"

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

type BalanceSheetReportsMongo struct {
	FiscalDateEnding                       string  `json:"fiscalDateEnding,omitempty"                       bson:"fiscalDateEnding,omitempty"`
	ReportedCurrency                       string  `json:"reportedCurrency,omitempty"                       bson:"reportedCurrency,omitempty"`
	TotalAssets                            float64 `json:"totalAssets,omitempty"                            bson:"totalAssets,omitempty"`
	TotalCurrentAssets                     float64 `json:"totalCurrentAssets,omitempty"                     bson:"totalCurrentAssets,omitempty"`
	CashAndCashEquivalentsAtCarryingValue  float64 `json:"cashAndCashEquivalentsAtCarryingValue,omitempty"  bson:"cashAndCashEquivalentsAtCarryingValue,omitempty"`
	CashAndShortTermInvestments            float64 `json:"cashAndShortTermInvestments,omitempty"            bson:"cashAndShortTermInvestments,omitempty"`
	Inventory                              float64 `json:"inventory,omitempty"                              bson:"inventory,omitempty"`
	CurrentNetReceivables                  float64 `json:"currentNetReceivables,omitempty"                  bson:"currentNetReceivables,omitempty"`
	TotalNonCurrentAssets                  float64 `json:"totalNonCurrentAssets,omitempty"                  bson:"totalNonCurrentAssets,omitempty"`
	PropertyPlantEquipment                 float64 `json:"propertyPlantEquipment,omitempty"                 bson:"propertyPlantEquipment,omitempty"`
	AccumulatedDepreciationAmortizationPPE float64 `json:"accumulatedDepreciationAmortizationPPE,omitempty" bson:"accumulatedDepreciationAmortizationPPE,omitempty"`
	IntangibleAssets                       float64 `json:"intangibleAssets,omitempty"                       bson:"intangibleAssets,omitempty"`
	IntangibleAssetsExcludingGoodwill      float64 `json:"intangibleAssetsExcludingGoodwill,omitempty"      bson:"intangibleAssetsExcludingGoodwill,omitempty"`
	Goodwill                               float64 `json:"goodwill,omitempty"                               bson:"goodwill,omitempty"`
	Investments                            float64 `json:"investments,omitempty"                            bson:"investments,omitempty"`
	LongTermInvestments                    float64 `json:"longTermInvestments,omitempty"                    bson:"longTermInvestments,omitempty"`
	ShortTermInvestments                   float64 `json:"shortTermInvestments,omitempty"                   bson:"shortTermInvestments,omitempty"`
	OtherCurrentAssets                     float64 `json:"otherCurrentAssets,omitempty"                     bson:"otherCurrentAssets,omitempty"`
	OtherNonCurrrentAssets                 float64 `json:"otherNonCurrrentAssets,omitempty"                 bson:"otherNonCurrrentAssets,omitempty"`
	TotalLiabilities                       float64 `json:"totalLiabilities,omitempty"                       bson:"totalLiabilities,omitempty"`
	TotalCurrentLiabilities                float64 `json:"totalCurrentLiabilities,omitempty"                bson:"totalCurrentLiabilities,omitempty"`
	CurrentAccountsPayable                 float64 `json:"currentAccountsPayable,omitempty"                 bson:"currentAccountsPayable,omitempty"`
	DeferredRevenue                        float64 `json:"deferredRevenue,omitempty"                        bson:"deferredRevenue,omitempty"`
	CurrentDebt                            float64 `json:"currentDebt,omitempty"                            bson:"currentDebt,omitempty"`
	ShortTermDebt                          float64 `json:"shortTermDebt,omitempty"                          bson:"shortTermDebt,omitempty"`
	TotalNonCurrentLiabilities             float64 `json:"totalNonCurrentLiabilities,omitempty"             bson:"totalNonCurrentLiabilities,omitempty"`
	CapitalLeaseObligations                float64 `json:"capitalLeaseObligations,omitempty"                bson:"capitalLeaseObligations,omitempty"`
	LongTermDebt                           float64 `json:"longTermDebt,omitempty"                           bson:"longTermDebt,omitempty"`
	CurrentLongTermDebt                    float64 `json:"currentLongTermDebt,omitempty"                    bson:"currentLongTermDebt,omitempty"`
	LongTermDebtNoncurrent                 float64 `json:"longTermDebtNoncurrent,omitempty"                 bson:"longTermDebtNoncurrent,omitempty"`
	ShortLongTermDebtTotal                 float64 `json:"shortLongTermDebtTotal,omitempty"                 bson:"shortLongTermDebtTotal,omitempty"`
	OtherCurrentLiabilities                float64 `json:"otherCurrentLiabilities,omitempty"                bson:"otherCurrentLiabilities,omitempty"`
	OtherNonCurrentLiabilities             float64 `json:"otherNonCurrentLiabilities,omitempty"             bson:"otherNonCurrentLiabilities,omitempty"`
	TotalShareholderEquity                 float64 `json:"totalShareholderEquity,omitempty"                 bson:"totalShareholderEquity,omitempty"`
	TreasuryStock                          float64 `json:"treasuryStock,omitempty"                          bson:"treasuryStock,omitempty"`
	RetainedEarnings                       float64 `json:"retainedEarnings,omitempty"                       bson:"retainedEarnings,omitempty"`
	CommonStock                            float64 `json:"commonStock,omitempty"                            bson:"commonStock,omitempty"`
	CommonStockSharesOutstanding           float64 `json:"commonStockSharesOutstanding,omitempty"           bson:"commonStockSharesOutstanding,omitempty"`
}

func (r *BalanceSheetReportsMongo) Set(v BalanceSheetReports) {
	TotalAssets, _ := strconv.ParseFloat(v.TotalAssets, 64)
	TotalCurrentAssets, _ := strconv.ParseFloat(v.TotalCurrentAssets, 64)
	CashAndCashEquivalentsAtCarryingValue, _ := strconv.ParseFloat(v.CashAndCashEquivalentsAtCarryingValue, 64)
	CashAndShortTermInvestments, _ := strconv.ParseFloat(v.CashAndShortTermInvestments, 64)
	Inventory, _ := strconv.ParseFloat(v.Inventory, 64)
	CurrentNetReceivables, _ := strconv.ParseFloat(v.CurrentNetReceivables, 64)
	TotalNonCurrentAssets, _ := strconv.ParseFloat(v.TotalNonCurrentAssets, 64)
	PropertyPlantEquipment, _ := strconv.ParseFloat(v.PropertyPlantEquipment, 64)
	AccumulatedDepreciationAmortizationPPE, _ := strconv.ParseFloat(v.AccumulatedDepreciationAmortizationPPE, 64)
	IntangibleAssets, _ := strconv.ParseFloat(v.IntangibleAssets, 64)
	IntangibleAssetsExcludingGoodwill, _ := strconv.ParseFloat(v.IntangibleAssetsExcludingGoodwill, 64)
	Goodwill, _ := strconv.ParseFloat(v.Goodwill, 64)
	Investments, _ := strconv.ParseFloat(v.Investments, 64)
	LongTermInvestments, _ := strconv.ParseFloat(v.LongTermInvestments, 64)
	ShortTermInvestments, _ := strconv.ParseFloat(v.ShortTermInvestments, 64)
	OtherCurrentAssets, _ := strconv.ParseFloat(v.OtherCurrentAssets, 64)
	OtherNonCurrrentAssets, _ := strconv.ParseFloat(v.OtherNonCurrrentAssets, 64)
	TotalLiabilities, _ := strconv.ParseFloat(v.TotalLiabilities, 64)
	TotalCurrentLiabilities, _ := strconv.ParseFloat(v.TotalCurrentLiabilities, 64)
	CurrentAccountsPayable, _ := strconv.ParseFloat(v.CurrentAccountsPayable, 64)
	DeferredRevenue, _ := strconv.ParseFloat(v.DeferredRevenue, 64)
	CurrentDebt, _ := strconv.ParseFloat(v.CurrentDebt, 64)
	ShortTermDebt, _ := strconv.ParseFloat(v.ShortTermDebt, 64)
	TotalNonCurrentLiabilities, _ := strconv.ParseFloat(v.TotalNonCurrentLiabilities, 64)
	CapitalLeaseObligations, _ := strconv.ParseFloat(v.CapitalLeaseObligations, 64)
	LongTermDebt, _ := strconv.ParseFloat(v.LongTermDebt, 64)
	CurrentLongTermDebt, _ := strconv.ParseFloat(v.CurrentLongTermDebt, 64)
	LongTermDebtNoncurrent, _ := strconv.ParseFloat(v.LongTermDebtNoncurrent, 64)
	ShortLongTermDebtTotal, _ := strconv.ParseFloat(v.ShortLongTermDebtTotal, 64)
	OtherCurrentLiabilities, _ := strconv.ParseFloat(v.OtherCurrentLiabilities, 64)
	OtherNonCurrentLiabilities, _ := strconv.ParseFloat(v.OtherNonCurrentLiabilities, 64)
	TotalShareholderEquity, _ := strconv.ParseFloat(v.TotalShareholderEquity, 64)
	TreasuryStock, _ := strconv.ParseFloat(v.TreasuryStock, 64)
	RetainedEarnings, _ := strconv.ParseFloat(v.RetainedEarnings, 64)
	CommonStock, _ := strconv.ParseFloat(v.CommonStock, 64)
	CommonStockSharesOutstanding, _ := strconv.ParseFloat(v.CommonStockSharesOutstanding, 64)

	r.FiscalDateEnding = v.FiscalDateEnding
	r.ReportedCurrency = v.ReportedCurrency
	r.TotalAssets = TotalAssets
	r.TotalCurrentAssets = TotalCurrentAssets
	r.CashAndCashEquivalentsAtCarryingValue = CashAndCashEquivalentsAtCarryingValue
	r.CashAndShortTermInvestments = CashAndShortTermInvestments
	r.Inventory = Inventory
	r.CurrentNetReceivables = CurrentNetReceivables
	r.TotalNonCurrentAssets = TotalNonCurrentAssets
	r.PropertyPlantEquipment = PropertyPlantEquipment
	r.AccumulatedDepreciationAmortizationPPE = AccumulatedDepreciationAmortizationPPE
	r.IntangibleAssets = IntangibleAssets
	r.IntangibleAssetsExcludingGoodwill = IntangibleAssetsExcludingGoodwill
	r.Goodwill = Goodwill
	r.Investments = Investments
	r.LongTermInvestments = LongTermInvestments
	r.ShortTermInvestments = ShortTermInvestments
	r.OtherCurrentAssets = OtherCurrentAssets
	r.OtherNonCurrrentAssets = OtherNonCurrrentAssets
	r.TotalLiabilities = TotalLiabilities
	r.TotalCurrentLiabilities = TotalCurrentLiabilities
	r.CurrentAccountsPayable = CurrentAccountsPayable
	r.DeferredRevenue = DeferredRevenue
	r.CurrentDebt = CurrentDebt
	r.ShortTermDebt = ShortTermDebt
	r.TotalNonCurrentLiabilities = TotalNonCurrentLiabilities
	r.CapitalLeaseObligations = CapitalLeaseObligations
	r.LongTermDebt = LongTermDebt
	r.CurrentLongTermDebt = CurrentLongTermDebt
	r.LongTermDebtNoncurrent = LongTermDebtNoncurrent
	r.ShortLongTermDebtTotal = ShortLongTermDebtTotal
	r.OtherCurrentLiabilities = OtherCurrentLiabilities
	r.OtherNonCurrentLiabilities = OtherNonCurrentLiabilities
	r.TotalShareholderEquity = TotalShareholderEquity
	r.TreasuryStock = TreasuryStock
	r.RetainedEarnings = RetainedEarnings
	r.CommonStock = CommonStock
	r.CommonStockSharesOutstanding = CommonStockSharesOutstanding
}

type BalanceSheetMongo struct {
	Symbol           string                     `json:"symbol"    bson:"symbol"`
	AnnualReports    []BalanceSheetReportsMongo `json:"annual"    bson:"annual"`
	QuarterlyReports []BalanceSheetReportsMongo `json:"quarterly" bson:"quarterly"`
}

func (e *BalanceSheetMongo) Set(v BalanceSheet) {
	var annual []BalanceSheetReportsMongo
	for _, cell := range v.AnnualReports {
		var NewEM BalanceSheetReportsMongo
		NewEM.Set(cell)

		annual = append(annual, NewEM)
	}

	var quarterly []BalanceSheetReportsMongo
	for _, cell := range v.QuarterlyReports {
		var NewQM BalanceSheetReportsMongo
		NewQM.Set(cell)

		quarterly = append(quarterly, NewQM)
	}

	e.Symbol = v.Symbol
	e.AnnualReports = annual
	e.QuarterlyReports = quarterly
}
