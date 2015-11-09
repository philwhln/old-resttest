package main

import (
	"./resttest"
	"fmt"
	"log"
)

func main() {
	if balance, err := resttest.Balance(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f\n", balance)
	}
}
