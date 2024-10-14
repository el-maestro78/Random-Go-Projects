package apriori

import (
	"fmt"
)

func getItems(transactions [][]string) []string {
	itemsMap := make(map[string]struct{})
	var items []string
	for _, row := range transactions {
		for _, item := range row {
			itemsMap[item] = struct{}{}
		}
	}
	for item := range itemsMap {
		items = append(items, item)
	}
	return items
}

func getCombinations(items []string, k int) [][]string {
	var combinations [][]string
	var combination []string
	var backtrack func(start int)

	backtrack = func(start int) {
		if len(combination) == k {
			combinationCopy := make([]string, k)
			copy(combinationCopy, combination)
			combinations = append(combinations, combinationCopy)
			return
		}
		for i := start; i < len(items); i++ {
			combination = append(combination, items[i])
			backtrack(i + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0)
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
				//key := strings.Join(comb, "")
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

func addToFrequent(map1, frequent map[string]int) map[string]int {
	for key, value := range map1 {
		frequent[key] = value
	}
	return frequent
}

func Apriori(transactions [][]string, minSupport int) map[string]int {
	var freqItemsets map[string]int
	itemKeys := getItems(transactions)
	C1 := generateCandidate1Itemsets(transactions)
	L1 := filter(C1, minSupport)
	//freqItemsets = append(freqItemsets, L1)
	freqItemsets = addToFrequent(freqItemsets, L1)
	k := 2
	for {
		fmt.Printf("Apriori recursions %d\n", k-1)
		Ck := generateCandidateKItemsets(transactions, k, itemKeys)
		Lk := filter(Ck, minSupport)
		if Lk == nil {
			break
		}
		//freqItemsets = append(freqItemsets, Lk)
		freqItemsets = addToFrequent(freqItemsets, Lk)
		k++
	}
	return freqItemsets
}
