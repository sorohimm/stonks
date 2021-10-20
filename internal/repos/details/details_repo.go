package details_repo

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/company_details"
)

type CompanyDetailsRepo struct {
	client *http.Client
}

var structs map[string]interface{}

func init() {
	structs = make(map[string]interface{})
	structs["OVERVIEW"] = details_models.Overview{}
	structs["EARNINGS"] = details_models.Earnings{}
	structs["INCOME_STATEMENT"] = details_models.IncomeStatement{}
	structs["BALANCE_SHEET"] = details_models.BalanceSheet{}
	structs["CASH_FLOW"] = details_models.CashFlow{}
}

func (r *CompanyDetailsRepo) GetCompanyDetails(request *http.Request) (interface{}, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return nil, err
	}

	detailsBody := structs[request.URL.Query().Get("function")]

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&detailsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return detailsBody, nil
}
