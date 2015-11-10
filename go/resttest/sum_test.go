package resttest

import (
	"log"
	"testing"
)

func TestSumTransactions(t *testing.T) {
	t1234 := Transaction{
		Amount: "12.34",
		Date:   "2014-12-03",
	}
	t234n := Transaction{
		Amount: "-2.34",
		Date:   "2014-12-03",
	}
	t2234n := Transaction{
		Amount: "-22.34",
		Date:   "2014-12-03",
	}
	var transactions []Transaction = make([]Transaction, 0, 10)
	testSum(transactions, 0)
	transactions = append(transactions, t1234)
	testSum(transactions, 1234)
	transactions = append(transactions, t1234)
	testSum(transactions, 2468)
	transactions = append(transactions, t234n)
	testSum(transactions, 2234)
	transactions = append(transactions, t2234n)
	testSum(transactions, 0)
	transactions = append(transactions, t2234n)
	testSum(transactions, -2234)
}

func testSum(transactions []Transaction, expectedSum Cents) {
	sum, err := sumTransactions(transactions)
	if err != nil {
		log.Fatal(err)
	}
	if sum != expectedSum {
		log.Fatalf("transactions:%v expectedSum:%v got:%v", transactions, expectedSum, sum)
	}
}
