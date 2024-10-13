package eclat

import "strconv"

func getTransactionIDs(transactions [][]string) map[string][]int {
	var transactionIDs = make(map[string][]int)
	for i, transaction := range transactions {
		for id := range transaction {
			id := strconv.Itoa(id)
			if _, exists := transactionIDs[id]; exists {
				transactionIDs[id] = append(transactionIDs[id], i)
			} else {
				transactionIDs[id] = []int{i}
			}
		}
	}
	return transactionIDs
}

func eclat(prefix string, transactions [][]string, min_support int, freq_itemsets int) {
	transactionIDs := getTransactionIDs(transactions)
	for id, transaction := range transactionIDs {
		
	}
}
