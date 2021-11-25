package details_models

import (
	"strconv"
)

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

type IncomeStatementMongo struct {
	Symbol           string                        `json:"symbol"`
	AnnualReports    []IncomeStatementReportsMongo `json:"annualReports"`
	QuarterlyReports []IncomeStatementReportsMongo `json:"quarterlyReports"`
}

type IncomeStatementReportsMongo struct {
	FiscalDateEnding                  string  `json:"fiscalDateEnding,omitempty"                  bson:"fiscalDateEnding,omitempty"`
	ReportedCurrency                  string  `json:"reportedCurrency,omitempty"                  bson:"reportedCurrency,omitempty"`
	GrossProfit                       float64 `json:"grossProfit,omitempty"                       bson:"grossProfit,omitempty"`
	TotalRevenue                      float64 `json:"totalRevenue,omitempty"                      bson:"totalRevenue,omitempty"`
	CostOfRevenue                     float64 `json:"costOfRevenue,omitempty"                     bson:"costOfRevenue,omitempty"`
	CostofGoodsAndServicesSold        float64 `json:"costofGoodsAndServicesSold,omitempty"        bson:"costofGoodsAndServicesSold,omitempty"`
	OperatingIncome                   float64 `json:"operatingIncome,omitempty"                   bson:"operatingIncome,omitempty"`
	SellingGeneralAndAdministrative   float64 `json:"sellingGeneralAndAdministrative,omitempty"   bson:"sellingGeneralAndAdministrative,omitempty"`
	ResearchAndDevelopment            float64 `json:"researchAndDevelopment,omitempty"            bson:"researchAndDevelopment,omitempty"`
	OperatingExpenses                 float64 `json:"operatingExpenses,omitempty"                 bson:"operatingExpenses,omitempty"`
	InvestmentIncomeNet               float64 `json:"investmentIncomeNet,omitempty"               bson:"investmentIncomeNet,omitempty"`
	NetInterestIncome                 float64 `json:"netInterestIncome,omitempty"                 bson:"netInterestIncome,omitempty"`
	InterestIncome                    float64 `json:"interestIncome,omitempty"                    bson:"interestIncome,omitempty"`
	InterestExpense                   float64 `json:"interestExpense,omitempty"                   bson:"interestExpense,omitempty"`
	NonInterestIncome                 float64 `json:"nonInterestIncome,omitempty"                 bson:"nonInterestIncome,omitempty"`
	OtherNonOperatingIncome           float64 `json:"otherNonOperatingIncome,omitempty"           bson:"otherNonOperatingIncome,omitempty"`
	DepreciationAndAmortization       float64 `json:"depreciationAndAmortization,omitempty"       bson:"depreciationAndAmortization,omitempty"`
	IncomeBeforeTax                   float64 `json:"incomeBeforeTax,omitempty"                   bson:"incomeBeforeTax,omitempty"`
	IncomeTaxExpense                  float64 `json:"incomeTaxExpense,omitempty"                  bson:"incomeTaxExpense,omitempty"`
	InterestAndDebtExpense            float64 `json:"interestAndDebtExpense,omitempty"            bson:"interestAndDebtExpense,omitempty"`
	NetIncomeFromContinuingOperations float64 `json:"netIncomeFromContinuingOperations,omitempty" bson:"netIncomeFromContinuingOperations,omitempty"`
	ComprehensiveIncomeNetOfTax       float64 `json:"comprehensiveIncomeNetOfTax,omitempty"       bson:"comprehensiveIncomeNetOfTax,omitempty"`
	Ebit                              float64 `json:"ebit,omitempty"                              bson:"ebit,omitempty"`
	Ebitda                            float64 `json:"ebitda,omitempty"                            bson:"ebitda,omitempty"`
	NetIncome                         float64 `json:"net_income,omitempty"                        bson:"net_income,omitempty"`
}

func (s *IncomeStatementReportsMongo) Set(v IncomeStatementReports) {
	GrossProfit, _ := strconv.ParseFloat(v.GrossProfit, 64)
	TotalRevenue, _ := strconv.ParseFloat(v.TotalRevenue, 64)
	CostOfRevenue, _ := strconv.ParseFloat(v.CostOfRevenue, 64)
	CostofGoodsAndServicesSold, _ := strconv.ParseFloat(v.CostofGoodsAndServicesSold, 64)
	OperatingIncome, _ := strconv.ParseFloat(v.OperatingIncome, 64)
	SellingGeneralAndAdministrative, _ := strconv.ParseFloat(v.SellingGeneralAndAdministrative, 64)
	ResearchAndDevelopment, _ := strconv.ParseFloat(v.ResearchAndDevelopment, 64)
	OperatingExpenses, _ := strconv.ParseFloat(v.OperatingExpenses, 64)
	InvestmentIncomeNet, _ := strconv.ParseFloat(v.InvestmentIncomeNet, 64)
	NetInterestIncome, _ := strconv.ParseFloat(v.NetInterestIncome, 64)
	InterestIncome, _ := strconv.ParseFloat(v.InterestIncome, 64)
	InterestExpense, _ := strconv.ParseFloat(v.InterestExpense, 64)
	NonInterestIncome, _ := strconv.ParseFloat(v.NonInterestIncome, 64)
	OtherNonOperatingIncome, _ := strconv.ParseFloat(v.OtherNonOperatingIncome, 64)
	DepreciationAndAmortization, _ := strconv.ParseFloat(v.DepreciationAndAmortization, 64)
	IncomeBeforeTax, _ := strconv.ParseFloat(v.IncomeBeforeTax, 64)
	IncomeTaxExpense, _ := strconv.ParseFloat(v.IncomeTaxExpense, 64)
	InterestAndDebtExpense, _ := strconv.ParseFloat(v.InterestAndDebtExpense, 64)
	NetIncomeFromContinuingOperations, _ := strconv.ParseFloat(v.NetIncomeFromContinuingOperations, 64)
	ComprehensiveIncomeNetOfTax, _ := strconv.ParseFloat(v.ComprehensiveIncomeNetOfTax, 64)
	Ebit, _ := strconv.ParseFloat(v.Ebit, 64)
	Ebitda, _ := strconv.ParseFloat(v.Ebitda, 64)
	NetIncome, _ := strconv.ParseFloat(v.NetIncome, 64)

	s.FiscalDateEnding = v.FiscalDateEnding
	s.ReportedCurrency = v.ReportedCurrency
	s.GrossProfit = GrossProfit
	s.TotalRevenue = TotalRevenue
	s.CostOfRevenue = CostOfRevenue
	s.CostofGoodsAndServicesSold = CostofGoodsAndServicesSold
	s.OperatingIncome = OperatingIncome
	s.SellingGeneralAndAdministrative = SellingGeneralAndAdministrative
	s.ResearchAndDevelopment = ResearchAndDevelopment
	s.OperatingExpenses = OperatingExpenses
	s.InvestmentIncomeNet = InvestmentIncomeNet
	s.NetInterestIncome = NetInterestIncome
	s.InterestIncome = InterestIncome
	s.InterestExpense = InterestExpense
	s.NonInterestIncome = NonInterestIncome
	s.OtherNonOperatingIncome = OtherNonOperatingIncome
	s.DepreciationAndAmortization = DepreciationAndAmortization
	s.IncomeBeforeTax = IncomeBeforeTax
	s.IncomeTaxExpense = IncomeTaxExpense
	s.InterestAndDebtExpense = InterestAndDebtExpense
	s.NetIncomeFromContinuingOperations = NetIncomeFromContinuingOperations
	s.ComprehensiveIncomeNetOfTax = ComprehensiveIncomeNetOfTax
	s.Ebit = Ebit
	s.Ebitda = Ebitda
	s.NetIncome = NetIncome
}

func (s *IncomeStatementMongo) Set(v IncomeStatement) {
	var annual []IncomeStatementReportsMongo
	for _, cell := range v.AnnualReports {
		var NewISM IncomeStatementReportsMongo
		NewISM.Set(cell)

		annual = append(annual, NewISM)
	}

	var quarterly []IncomeStatementReportsMongo
	for _, cell := range v.QuarterlyReports {
		var NewISM IncomeStatementReportsMongo
		NewISM.Set(cell)

		quarterly = append(quarterly, NewISM)
	}

	s.Symbol = v.Symbol
	s.AnnualReports = annual
	s.QuarterlyReports = quarterly
}
