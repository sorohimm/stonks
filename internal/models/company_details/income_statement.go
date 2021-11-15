package details_models

import (
	"errors"
	"net/url"
	"stonks/internal/models"
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

type DQuarterlyIncomeStatement struct {
	Symbol            string              `json:"symbol"`
	QuarterlyEarnings []IncomeStatementReports `json:"quarterlyReports"`
}

type DAnnualIncomeStatement struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings []IncomeStatementReports `json:"annualReports"`
}

type IncomeStatement struct {
	Symbol           string                   `json:"symbol"`
	AnnualReports    []IncomeStatementReports `json:"annualReports"`
	QuarterlyReports []IncomeStatementReports `json:"quarterlyReports"`
}

func (i *IncomeStatement) Annual(t models.Timing) ([]IncomeStatementReports, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	var res []IncomeStatementReports
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range i.AnnualReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-5]

			if fde >= t.From && fde <= t.To {
				res = append(res, el)
			}
		}

		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}
		return res, nil
	}

	if t.HasFrom() && !t.HasTo() && !t.HasDate() {
		for _, el := range i.AnnualReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-5]
			if fde >= t.From {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if !t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range i.AnnualReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-5]
			if fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasDate() && !t.HasFrom() && !t.HasTo() {
		for _, el := range i.AnnualReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-5]
			if fde == t.Date {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	return res, errors.New("bad timing")
}

func (i *IncomeStatement) Quarterly(t models.Timing) ([]IncomeStatementReports, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	var res []IncomeStatementReports
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range i.QuarterlyReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde >= t.From && fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasFrom() && !t.HasTo() && !t.HasDate() {
		for _, el := range i.QuarterlyReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde >= t.From {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if !t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range i.QuarterlyReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasDate() && !t.HasFrom() && !t.HasTo() {
		for _, el := range i.QuarterlyReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde == t.Date {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	return res, errors.New("bad timing")
}

func (i *IncomeStatement) ByTiming(values url.Values) (interface{}, error) {
	var timings models.Timing
	timings.Set(values)
	if values.Get("timing") == "annual" {
		res, err := i.Annual(timings)
		if err != nil {
			return DAnnualIncomeStatement{Symbol: values.Get("symbol"),
				AnnualEarnings: i.AnnualReports}, nil
		}
		return DAnnualIncomeStatement{Symbol: values.Get("symbol"), AnnualEarnings: res}, nil
	} else if values.Get("timing") == "quarterly" {
		res, err := i.Quarterly(timings)
		if err != nil {
			return DQuarterlyIncomeStatement{Symbol: values.Get("symbol"),
				QuarterlyEarnings: i.QuarterlyReports}, nil
		}
		return DQuarterlyIncomeStatement{Symbol: values.Get("symbol"), QuarterlyEarnings: res}, nil
	}
	return nil, errors.New("unresolved error")
}

