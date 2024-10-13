package main

import "fmt"

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
	fmt.Println(transactions)
}
