package resttest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const TRANSACTIONS_PAGE_URL = "http://resttest.bench.co/transactions/%d.json"

// Limit how many pages we support
const MAX_PAGE = 999

type Transaction struct {
	Amount string `json:"Amount"`
}

type TransactionsPage struct {
	TotalCount   int `json:"totalCount"`
	Transactions []Transaction
}

// Page through the resttest API, calculating the overview total balance
// from the transactions amounts listed there.
func Balance() (balance Cents, err error) {
	var count int
	totalCount := -1
	for page := 1; page == 1 || count < totalCount; page++ {
		var tp TransactionsPage
		var sum Cents
		if page > MAX_PAGE {
			return balance, errors.New("Exceeded MAX_PAGE")
		}
		tp, err = transactionsPage(page)
		if err != nil {
			return balance, err
		}
		if page == 1 {
			totalCount = tp.TotalCount
		}
		sum, err = sumTransactions(tp.Transactions)
		if err != nil {
			return balance, err
		}
		balance += sum
		count += len(tp.Transactions)
	}
	return balance, nil
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

// Given an array of transactions, calculate the sum
func sumTransactions(transactions []Transaction) (sum Cents, err error) {
	for _, t := range transactions {
		var dollars float64
		var cents Cents
		dollars, err = strconv.ParseFloat(t.Amount, 64)
		if err != nil {
			return sum, err
		}
		cents = Cents(dollars * 100.0)
		// XXX Check for int64 overflow
		sum += cents
	}
	return sum, nil
}
