package resttest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TRANSACTIONS_PAGE_URL = "http://resttest.bench.co/transactions/%d.json"

// Limit how many pages we support
const MAX_PAGE = 999

type Transaction struct {
	Amount string `json:"Amount"`
	Date   string `json:"Date"`
}

type TransactionsPage struct {
	TotalCount   int `json:"totalCount"`
	Transactions []Transaction
}

// Fetch the data for specific page of resttest API in a structured
// formated
func transactionsPage(page int) (tp TransactionsPage, err error) {
	var j []byte
	if j, err = transactionPageJson(page); err != nil {
		return
	}
	err = json.Unmarshal(j, &tp)
	return
}

// Fetch the json blob for a specific page of the resttest API
func transactionPageJson(page int) (json []byte, err error) {
	url := fmt.Sprintf(TRANSACTIONS_PAGE_URL, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
