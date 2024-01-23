package main

func maxLengthBetweenEqualCharacters(s string) int {
	firstSeenAt := [26]int{}

	l := len(s) + 1

	for i := 0; i < 26; i++ {
		firstSeenAt[i] = l
	}

	max := -1

	for i, r := range s {
		if firstSeenAt[r-'a'] == l {
			firstSeenAt[r] = i
		}

		length := i - firstSeenAt[r] - 1

		if length > max {
			max = length
		}
	}

	return max
}
