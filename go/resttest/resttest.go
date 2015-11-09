package resttest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TRANSACTIONS_PAGE_URL = "http://resttest.bench.co/transactions/%d.json"

type Transaction struct {
	Amount string `json:"Amount"`
}

type TransactionsPage struct {
	TotalCount   int `json:"totalCount"`
	Transactions []Transaction
}

func Balance() (balance float32, err error) {
	// TODO : Get more than one page
	page := 1
	body, err := transactionPageJson(page)
	if err != nil {
		return 0.0, err
	}
	// TODO : Remove this
	fmt.Printf("%s\n", body)
	// TODO : Replace fake value
	return 1.0, nil
}

func transactionPage(page int) (tp TransactionsPage, err error) {
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
