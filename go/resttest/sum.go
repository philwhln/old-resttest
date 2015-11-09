package resttest

import (
	"strconv"
)

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
