package resttest

import (
	"errors"
)

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
