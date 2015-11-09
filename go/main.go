package main

import (
	"./resttest"
	"fmt"
	"log"
)

// Calculate total balance for resttest and output it
func main() {
	if balance, err := resttest.Balance(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(balance)
	}
}
