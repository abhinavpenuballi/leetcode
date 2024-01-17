package main

func uniqueOccurrences(arr []int) bool {
	occurrences := map[int]int{}

	for _, i := range arr {
		occurrences[i]++
	}

	occurrencesOccurred := map[int]bool{}

	for _, i := range occurrences {
		if occurrencesOccurred[i] {
			return false
		}

		occurrencesOccurred[i] = true
	}

	return true
}
