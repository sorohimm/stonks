package details_models

import "strconv"

type CashFlowReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                                          string `json:"reportedCurrency,omitempty"`
	OperatingCashflow                                         string `json:"operatingCashflow,omitempty"`
	PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities,omitempty"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities,omitempty"`
	ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities,omitempty"`
	ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets,omitempty"`
	DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization,omitempty"`
	CapitalExpenditures                                       string `json:"capitalExpenditures,omitempty"`
	ChangeInReceivables                                       string `json:"changeInReceivables,omitempty"`
	ChangeInInventory                                         string `json:"changeInInventory,omitempty"`
	ProfitLoss                                                string `json:"profitLoss,omitempty"`
	CashflowFromInvestment                                    string `json:"cashflowFromInvestment,omitempty"`
	CashflowFromFinancing                                     string `json:"cashflowFromFinancing,omitempty"`
	ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt,omitempty"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock,omitempty"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity,omitempty"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock,omitempty"`
	DividendPayout                                            string `json:"dividendPayout,omitempty"`
	DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock,omitempty"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock,omitempty"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock,omitempty"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,omitempty"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock,omitempty"`
	ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity,omitempty"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock,omitempty"`
	ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents,omitempty"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate,omitempty"`
	NetIncome                                                 string `json:"netIncome,omitempty"`
}

type CashFlow struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []CashFlowReport `json:"annualReports"`
	QuarterlyReports []CashFlowReport `json:"quarterlyReports"`
}

type CashFlowReportMongo struct {
	FiscalDateEnding                                          string  `json:"fiscalDateEnding,omitempty" bson:"fiscalDateEnding"`
	ReportedCurrency                                          string  `json:"reportedCurrency,omitempty" bson:"reportedCurrency"`
	OperatingCashflow                                         float64 `json:"operatingCashflow,omitempty" bson:"operatingCashflow"`
	PaymentsForOperatingActivities                            float64 `json:"paymentsForOperatingActivities,omitempty" bson:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           float64 `json:"proceedsFromOperatingActivities,omitempty" bson:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              float64 `json:"changeInOperatingLiabilities,omitempty" bson:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   float64 `json:"changeInOperatingAssets,omitempty" bson:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      float64 `json:"depreciationDepletionAndAmortization,omitempty" bson:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       float64 `json:"capitalExpenditures,omitempty" bson:"capitalExpenditures"`
	ChangeInReceivables                                       float64 `json:"changeInReceivables,omitempty" bson:"changeInReceivables"`
	ChangeInInventory                                         float64 `json:"changeInInventory,omitempty" bson:"changeInInventory"`
	ProfitLoss                                                float64 `json:"profitLoss,omitempty" bson:"profitLoss"`
	CashflowFromInvestment                                    float64 `json:"cashflowFromInvestment,omitempty" bson:"cashflowFromInvestment"`
	CashflowFromFinancing                                     float64 `json:"cashflowFromFinancing,omitempty" bson:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     float64 `json:"proceedsFromRepaymentsOfShortTermDebt,omitempty" bson:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        float64 `json:"paymentsForRepurchaseOfCommonStock,omitempty" bson:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             float64 `json:"paymentsForRepurchaseOfEquity,omitempty" bson:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     float64 `json:"paymentsForRepurchaseOfPreferredStock,omitempty" bson:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            float64 `json:"dividendPayout,omitempty" bson:"dividendPayout"`
	DividendPayoutCommonStock                                 float64 `json:"dividendPayoutCommonStock,omitempty" bson:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              float64 `json:"dividendPayoutPreferredStock,omitempty" bson:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         float64 `json:"proceedsFromIssuanceOfCommonStock,omitempty" bson:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet float64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,omitempty" bson:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      float64 `json:"proceedsFromIssuanceOfPreferredStock,omitempty" bson:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            float64 `json:"proceedsFromRepurchaseOfEquity,omitempty" bson:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           float64 `json:"proceedsFromSaleOfTreasuryStock,omitempty" bson:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            float64 `json:"changeInCashAndCashEquivalents,omitempty" bson:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      float64 `json:"changeInExchangeRate,omitempty" bson:"changeInExchangeRate"`
	NetIncome                                                 float64 `json:"netIncome,omitempty" bson:"netIncome"`
}

func (c *CashFlowReportMongo) Set(v *CashFlowReport) {
	OperatingCashflow, _ := strconv.ParseFloat(v.OperatingCashflow, 64)
	PaymentsForOperatingActivities, _ := strconv.ParseFloat(v.PaymentsForOperatingActivities, 64)
	ProceedsFromOperatingActivities, _ := strconv.ParseFloat(v.ProceedsFromOperatingActivities, 64)
	ChangeInOperatingLiabilities, _ := strconv.ParseFloat(v.ChangeInOperatingLiabilities, 64)
	ChangeInOperatingAssets, _ := strconv.ParseFloat(v.ChangeInOperatingAssets, 64)
	DepreciationDepletionAndAmortization, _ := strconv.ParseFloat(v.DepreciationDepletionAndAmortization, 64)
	CapitalExpenditures, _ := strconv.ParseFloat(v.CapitalExpenditures, 64)
	ChangeInReceivables, _ := strconv.ParseFloat(v.ChangeInReceivables, 64)
	ChangeInInventory, _ := strconv.ParseFloat(v.ChangeInInventory, 64)
	ProfitLoss, _ := strconv.ParseFloat(v.ProfitLoss, 64)
	CashflowFromInvestment, _ := strconv.ParseFloat(v.CashflowFromInvestment, 64)
	CashflowFromFinancing, _ := strconv.ParseFloat(v.CashflowFromFinancing, 64)
	ProceedsFromRepaymentsOfShortTermDebt, _ := strconv.ParseFloat(v.ProceedsFromRepaymentsOfShortTermDebt, 64)
	PaymentsForRepurchaseOfCommonStock, _ := strconv.ParseFloat(v.PaymentsForRepurchaseOfCommonStock, 64)
	PaymentsForRepurchaseOfEquity, _ := strconv.ParseFloat(v.PaymentsForRepurchaseOfEquity, 64)
	PaymentsForRepurchaseOfPreferredStock, _ := strconv.ParseFloat(v.PaymentsForRepurchaseOfPreferredStock, 64)
	DividendPayout, _ := strconv.ParseFloat(v.DividendPayout, 64)
	DividendPayoutCommonStock, _ := strconv.ParseFloat(v.DividendPayoutCommonStock, 64)
	DividendPayoutPreferredStock, _ := strconv.ParseFloat(v.DividendPayoutPreferredStock, 64)
	ProceedsFromIssuanceOfCommonStock, _ := strconv.ParseFloat(v.ProceedsFromIssuanceOfCommonStock, 64)
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet, _ := strconv.ParseFloat(v.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet, 64)
	ProceedsFromIssuanceOfPreferredStock, _ := strconv.ParseFloat(v.ProceedsFromIssuanceOfPreferredStock, 64)
	ProceedsFromRepurchaseOfEquity, _ := strconv.ParseFloat(v.ProceedsFromRepurchaseOfEquity, 64)
	ProceedsFromSaleOfTreasuryStock, _ := strconv.ParseFloat(v.ProceedsFromSaleOfTreasuryStock, 64)
	ChangeInCashAndCashEquivalents, _ := strconv.ParseFloat(v.ChangeInCashAndCashEquivalents, 64)
	ChangeInExchangeRate, _ := strconv.ParseFloat(v.ChangeInExchangeRate, 64)
	NetIncome, _ := strconv.ParseFloat(v.NetIncome, 64)

	c.FiscalDateEnding = v.FiscalDateEnding
	c.ReportedCurrency = v.ReportedCurrency

	c.OperatingCashflow = OperatingCashflow
	c.PaymentsForOperatingActivities = PaymentsForOperatingActivities
	c.ProceedsFromOperatingActivities = ProceedsFromOperatingActivities
	c.ChangeInOperatingLiabilities = ChangeInOperatingLiabilities
	c.ChangeInOperatingAssets = ChangeInOperatingAssets
	c.DepreciationDepletionAndAmortization = DepreciationDepletionAndAmortization
	c.CapitalExpenditures = CapitalExpenditures
	c.ChangeInReceivables = ChangeInReceivables
	c.ChangeInInventory = ChangeInInventory
	c.ProfitLoss = ProfitLoss
	c.CashflowFromInvestment = CashflowFromInvestment
	c.CashflowFromFinancing = CashflowFromFinancing
	c.ProceedsFromRepaymentsOfShortTermDebt = ProceedsFromRepaymentsOfShortTermDebt
	c.PaymentsForRepurchaseOfCommonStock = PaymentsForRepurchaseOfCommonStock
	c.PaymentsForRepurchaseOfEquity = PaymentsForRepurchaseOfEquity
	c.PaymentsForRepurchaseOfPreferredStock = PaymentsForRepurchaseOfPreferredStock
	c.DividendPayout = DividendPayout
	c.DividendPayoutCommonStock = DividendPayoutCommonStock
	c.DividendPayoutPreferredStock = DividendPayoutPreferredStock
	c.ProceedsFromIssuanceOfCommonStock = ProceedsFromIssuanceOfCommonStock
	c.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet = ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet
	c.ProceedsFromIssuanceOfPreferredStock = ProceedsFromIssuanceOfPreferredStock
	c.ProceedsFromRepurchaseOfEquity = ProceedsFromRepurchaseOfEquity
	c.ProceedsFromSaleOfTreasuryStock = ProceedsFromSaleOfTreasuryStock
	c.ChangeInCashAndCashEquivalents = ChangeInCashAndCashEquivalents
	c.ChangeInExchangeRate = ChangeInExchangeRate
	c.NetIncome = NetIncome

}

type CashFlowMongo struct {
	Symbol           string                `json:"symbol,omitempty" bson:"symbol,omitempty"`
	AnnualReports    []CashFlowReportMongo `json:"annual,omitempty" bson:"annual,omitempty"`
	QuarterlyReports []CashFlowReportMongo `json:"quarterly,omitempty" bson:"quarterly,omitempty"`
}

func (c *CashFlowMongo) Set(v *CashFlow) {
	var annual []CashFlowReportMongo
	for _, cell := range v.AnnualReports {
		var NewCFRM CashFlowReportMongo
		NewCFRM.Set(&cell)

		annual = append(annual, NewCFRM)
	}

	var quarterly []CashFlowReportMongo
	for _, cell := range v.QuarterlyReports {
		var NewCFRM CashFlowReportMongo
		NewCFRM.Set(&cell)

		quarterly = append(quarterly, NewCFRM)
	}

	c.Symbol = v.Symbol
	c.AnnualReports = annual
	c.QuarterlyReports = quarterly
}
