package main

import (
	"eclat"
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
	results := eclat.eclat(transactions, 2)
	fmt.Println(transactions)
	fmt.Println(results)
}
