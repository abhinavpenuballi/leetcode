package main

func longestPalindrome(s string) int {
	occurrences := map[rune]int{}

	for _, r := range s {
		occurrences[r]++
	}

	sum, hasOdd := 0, false

	for _, occurrence := range occurrences {
		if occurrence%2 == 0 {
			sum += occurrence
		} else {
			sum += occurrence - 1
			hasOdd = true
		}
	}

	if hasOdd {
		sum++
	}

	return sum
}
