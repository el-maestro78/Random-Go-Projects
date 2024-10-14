package apriori

func generateOneItemsets(transactions [][]string) map[string]int {
	C1 := make(map[string]int)
	for _, row := range transactions {
		for _, item := range row {
			C1[item]++
		}
	}
	return C1
}

func filter(transactions map[string]int, minSupport int) map[string]int {
	var filtered map[string]int
	for key, value := range transactions {
		if value >= minSupport {
			filtered[key] = value
		}
	}
	return filtered
}

func Apriori(transactions [][]string, minSupport int) []map[string]int {
	var freqItemsets []map[string]int
	oneItemSets := generateOneItemsets(transactions)
	L1 := filter(oneItemSets, minSupport)
	freqItemsets = append(freqItemsets, L1)
	k := 2
	for {
		Ck := 0
		Lk := 0
		if Lk < 1 {
			break
		}
		freqItemsets = append(freqItemsets, Lk)
		k++
	}

	return freqItemsets
}
