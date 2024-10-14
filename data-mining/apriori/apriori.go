package apriori

import "fmt"

func getItems(transactions [][]string) []string {
	var items []string
	for _, row := range transactions {
		for _, item := range row {
			items = append(items, item)
		}
	}
	return items
}

func getCombinations(items []string, k int) [][]string {
	var combinations [][]string
	for index, item := range items {
		if index <= (len(items) - 1) {
			for i := 0; i < k; i++ {
				combinations = append(combinations, []string{item, items[index+i]})
			}
		}
		items = append(items[:index], items[index+1:]...)
	}
	return combinations
}

func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func generateCandidateKItemsets(transactions [][]string, k int, items []string) map[string]int {
	C := make(map[string]int)
	combinations := getCombinations(items, k)
	for _, row := range transactions {
		for _, comb := range combinations {
			present := true
			for _, item := range comb {
				if !contains(row, item) {
					present = false
					break
				}
			}
			if present {
				key := fmt.Sprintf("%v", comb)
				C[key]++
			}
		}
	}
	return C
}

func generateCandidate1Itemsets(transactions [][]string) map[string]int {
	C1 := make(map[string]int)
	for _, row := range transactions {
		for _, item := range row {
			C1[item]++
		}
	}
	return C1
}

func filter(transactions map[string]int, minSupport int) map[string]int {
	filtered := make(map[string]int)
	for key, value := range transactions {
		if value >= minSupport {
			filtered[key] = value
		}
	}
	if len(filtered) == 0 {
		return nil
	}
	return filtered
}

func Apriori(transactions [][]string, minSupport int) []map[string]int {
	var freqItemsets []map[string]int
	itemKeys := getItems(transactions)
	C1 := generateCandidate1Itemsets(transactions)
	L1 := filter(C1, minSupport)
	freqItemsets = append(freqItemsets, L1)
	k := 2
	for {
		Ck := generateCandidateKItemsets(transactions, k, itemKeys)
		Lk := filter(Ck, minSupport)
		if Lk == nil {
			break
		}
		freqItemsets = append(freqItemsets, Lk)
		k++
	}
	return freqItemsets
}
