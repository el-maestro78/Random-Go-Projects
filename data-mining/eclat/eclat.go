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
func intersect(transaction1, transaction2 []int) []int {
	elements := make(map[int]bool)
	for _, v := range transaction1 {
		elements[v] = true
	}
	var result []int
	for _, v := range transaction2 {
		if elements[v] {
			result = append(result, v)
			delete(elements, v)
		}
	}
	return result
}

func eclat(transactions [][]string, minSupport int) []string {
	var freqItemsets []string
	transactionItems := getTransactionIDs(transactions)
	for id, items := range transactionItems {
		if len(items) >= minSupport {
			eclatRecursive([]string{id}, items, transactionItems, minSupport, freqItemsets)
		}
	}
	return freqItemsets
}

func eclatRecursive(prefix []string, transactions []int, transactionItems map[string][]int, minSupport int, freqItemsets []string) {
	freqItemsets = append(freqItemsets, prefix[])
	for item, transaction := range transactionItems {
		if item <= prefix[len(prefix)-1] {
			new_transactions := intersect(transactions, transactionItems[item])
			if len(new_transactions) >= minSupport {
				new_prefix := append(prefix, item)
				eclatRecursive(new_prefix, new_transactions, transactionItems, minSupport, freqItemsets)
			}
		}
	}
}
