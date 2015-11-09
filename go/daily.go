package main

import (
	"./resttest"
	"fmt"
	"log"
	"sort"
)

// Calculate daily balances from resttest API and output in table
func main() {
	balances, err := resttest.DailyBalances()
	if err != nil {
		log.Fatal(err)
	}
	displayDailyBalances(balances)
}

// Display table of date and respective daily total.
// Also include grand total sum of all days.
func displayDailyBalances(balances map[string]resttest.Cents) {
	var total resttest.Cents = 0
	dates := sortedKeys(balances)
	for _, date := range dates {
		fmt.Printf("%s\t%s\n", date, balances[date])
		total += balances[date]
	}
	fmt.Println("----------\t-----")
	fmt.Printf("TOTAL     \t%s\n", total)
}

// For given map, return array of sorted keys from map
func sortedKeys(m map[string]resttest.Cents) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
