package resttest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//const TRANSACTIONS_PAGE_URL = "http://resttest.bench.co/transactions/%d.json"
const TRANSACTIONS_PAGE_URL = "http://0.0.0.0:8000/transactions/%d.json"

// Limit how many pages we support
const MAX_PAGE = 999

type Transaction struct {
	Amount string `json:"Amount"`
}

type TransactionsPage struct {
	TotalCount   int `json:"totalCount"`
	Transactions []Transaction
}

func Balance() (balance float64, err error) {
	var transactions int
	var sum float64
	advertisedTransactions := -1
	for page := 1; page == 1 || transactions < advertisedTransactions; page++ {
		var tp TransactionsPage
		tp, err = transactionsPage(page)
		if err != nil {
			return 0.0, err
		}
		if page == 1 {
			advertisedTransactions = tp.TotalCount
		}
		for _, t := range tp.Transactions {
			// ParseFloat return float64
			var amount float64
			amount, err = strconv.ParseFloat(t.Amount, 64)
			if err != nil {
				return
			}
			sum += amount
			transactions += 1
		}
		if page == MAX_PAGE {
			return 0.0, errors.New("Exceeded MAX_PAGE")
		}
	}
	return sum, nil
}

func transactionsPage(page int) (tp TransactionsPage, err error) {
	var j []byte
	if j, err = transactionPageJson(page); err != nil {
		return
	}
	err = json.Unmarshal(j, &tp)
	return
}

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
