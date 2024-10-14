package main

import (
	"data-mining/eclat"
	"fmt"
)

func getData() [][]string {
	// Dummy for now
	transactions := [][]string{
		{"a", "b", "c"},
		{"a", "b"},
		{"a", "c"},
		{"b", "c"},
		{"a", "b", "c", "d"},
	}
	return transactions
}

func main() {
	transactions := getData()
	results := eclat.Eclat(transactions, 3)
	fmt.Print("Transactions: ")
	fmt.Println(transactions)
	fmt.Print("Results: ")
	fmt.Println(results)
}
