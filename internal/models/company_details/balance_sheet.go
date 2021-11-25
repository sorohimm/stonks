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
	TotalAssets                            float64 `json:"totalAssets,omitempty"                            bson:"totalAssets"`
	TotalCurrentAssets                     float64 `json:"totalCurrentAssets,omitempty"                     bson:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  float64 `json:"cashAndCashEquivalentsAtCarryingValue,omitempty"  bson:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            float64 `json:"cashAndShortTermInvestments,omitempty"            bson:"cashAndShortTermInvestments"`
	Inventory                              float64 `json:"inventory,omitempty"                              bson:"inventory"`
	CurrentNetReceivables                  float64 `json:"currentNetReceivables,omitempty"                  bson:"currentNetReceivables"`
	TotalNonCurrentAssets                  float64 `json:"totalNonCurrentAssets,omitempty"                  bson:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 float64 `json:"propertyPlantEquipment,omitempty"                 bson:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE float64 `json:"accumulatedDepreciationAmortizationPPE,omitempty" bson:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       float64 `json:"intangibleAssets,omitempty"                       bson:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      float64 `json:"intangibleAssetsExcludingGoodwill,omitempty"      bson:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               float64 `json:"goodwill,omitempty"                               bson:"goodwill"`
	Investments                            float64 `json:"investments,omitempty"                            bson:"investments"`
	LongTermInvestments                    float64 `json:"longTermInvestments,omitempty"                    bson:"longTermInvestments"`
	ShortTermInvestments                   float64 `json:"shortTermInvestments,omitempty"                   bson:"shortTermInvestments"`
	OtherCurrentAssets                     float64 `json:"otherCurrentAssets,omitempty"                     bson:"otherCurrentAssets"`
	OtherNonCurrrentAssets                 float64 `json:"otherNonCurrrentAssets,omitempty"                 bson:"otherNonCurrrentAssets"`
	TotalLiabilities                       float64 `json:"totalLiabilities,omitempty"                       bson:"totalLiabilities"`
	TotalCurrentLiabilities                float64 `json:"totalCurrentLiabilities,omitempty"                bson:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 float64 `json:"currentAccountsPayable,omitempty"                 bson:"currentAccountsPayable"`
	DeferredRevenue                        float64 `json:"deferredRevenue,omitempty"                        bson:"deferredRevenue"`
	CurrentDebt                            float64 `json:"currentDebt,omitempty"                            bson:"currentDebt"`
	ShortTermDebt                          float64 `json:"shortTermDebt,omitempty"                          bson:"shortTermDebt"`
	TotalNonCurrentLiabilities             float64 `json:"totalNonCurrentLiabilities,omitempty"             bson:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                float64 `json:"capitalLeaseObligations,omitempty"                bson:"capitalLeaseObligations"`
	LongTermDebt                           float64 `json:"longTermDebt,omitempty"                           bson:"longTermDebt"`
	CurrentLongTermDebt                    float64 `json:"currentLongTermDebt,omitempty"                    bson:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 float64 `json:"longTermDebtNoncurrent,omitempty"                 bson:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 float64 `json:"shortLongTermDebtTotal,omitempty"                 bson:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                float64 `json:"otherCurrentLiabilities,omitempty"                bson:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             float64 `json:"otherNonCurrentLiabilities,omitempty"             bson:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 float64 `json:"totalShareholderEquity,omitempty"                 bson:"totalShareholderEquity"`
	TreasuryStock                          float64 `json:"treasuryStock,omitempty"                          bson:"treasuryStock"`
	RetainedEarnings                       float64 `json:"retainedEarnings,omitempty"                       bson:"retainedEarnings"`
	CommonStock                            float64 `json:"commonStock,omitempty"                            bson:"commonStock"`
	CommonStockSharesOutstanding           float64 `json:"commonStockSharesOutstanding,omitempty"           bson:"commonStockSharesOutstanding"`
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
	AnnualReports    []BalanceSheetReportsMongo `json:"annual,omitempty"    bson:"annual,omitempty"`
	QuarterlyReports []BalanceSheetReportsMongo `json:"quarterly,omitempty" bson:"quarterly,omitempty"`
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
