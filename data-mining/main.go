package main

import (
	"bufio"
	"data-mining/apriori"
	"data-mining/eclat"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter input: ")
	if !scanner.Scan() {

	}
	inputText := scanner.Text()
	minSupport, err := strconv.Atoi(inputText)
	if err != nil {
		fmt.Println("This is not a correct number, try again")
		main()
	}
	transactions := getData()
	fmt.Printf("For minimum Support = %d\n", minSupport)
	fmt.Print("Transactions: ")
	fmt.Println(transactions)
	eclatResults := eclat.Eclat(transactions, minSupport)
	fmt.Print("Eclat Results: ")
	fmt.Println(eclatResults)
	fmt.Println("-----")
	aprioriResults := apriori.Apriori(transactions, minSupport)
	fmt.Print("Apriori Results: ")
	fmt.Println(aprioriResults)
}
