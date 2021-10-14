package income_statement_repo

import (
	"encoding/json"
	"log"
	"net/http"
	ism "stonks/internal/models/income_statement"
)

type IncomeStatementRepo struct {
	client http.Client
}

func (r *IncomeStatementRepo) GetIncomeStatement(request *http.Request) (ism.IncomeStatement, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return ism.IncomeStatement{}, err
	}

	incomeStatementBody := ism.IncomeStatement{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&incomeStatementBody)
	if err != nil {
		log.Print(err.Error())
	}
	return incomeStatementBody, nil
}
