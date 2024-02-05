package main

func findThePrefixCommonArray(A []int, B []int) []int {
	seenInA, notYetSeenInA := map[int]bool{}, map[int]struct{}{}

	common := 0

	prefixCommonArray := []int{}

	for i := 0; i < len(A); i++ {
		seenInA[A[i]] = true

		if seenInA[B[i]] {
			common++
		} else {
			notYetSeenInA[B[i]] = struct{}{}
		}

		for val := range notYetSeenInA {
			if seenInA[val] {
				common++
				delete(notYetSeenInA, val)
			}
		}

		prefixCommonArray = append(prefixCommonArray, common)
	}

	return prefixCommonArray
}
