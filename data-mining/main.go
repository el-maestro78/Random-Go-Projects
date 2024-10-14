package main

import (
	"data-mining/apriori"
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
	fmt.Print("Transactions: ")
	fmt.Println(transactions)
	eclatResults := eclat.Eclat(transactions, 3)
	fmt.Print("Eclat Results: ")
	fmt.Println(eclatResults)
	fmt.Println("-----")
	aprioriResults := apriori.Apriori(transactions, 3)
	fmt.Print("Apriori Results: ")
	fmt.Println(aprioriResults)
}
