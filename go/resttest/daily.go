package resttest

import (
	"errors"
)

// Return map of date to list of daily transactions (in no particular order)
func DailyTransactions() (daily map[string][]Transaction, err error) {
	var count int
	totalCount := -1
	daily = make(map[string][]Transaction)
	for page := 1; page == 1 || count < totalCount; page++ {
		var tp TransactionsPage
		if page > MAX_PAGE {
			return daily, errors.New("Exceeded MAX_PAGE")
		}
		tp, err = transactionsPage(page)
		if err != nil {
			return daily, err
		}
		if page == 1 {
			totalCount = tp.TotalCount
		}
		for _, t := range tp.Transactions {
			date := t.Date
			daily[date] = append(daily[date], t)
		}
		count += len(tp.Transactions)
	}
	return daily, nil
}

// Return map of date to daily total balance
func DailyBalances() (balances map[string]Cents, err error) {
	var daily map[string][]Transaction
	daily, err = DailyTransactions()
	if err != nil {
		return balances, err
	}
	balances = make(map[string]Cents)
	for date, transactions := range daily {
		var sum Cents
		sum, err = sumTransactions(transactions)
		balances[date] = sum
		if err != nil {
			return balances, err
		}
	}
	return balances, nil
}
