package main

func partitionLabels(s string) []int {
	ends := map[rune]int{}

	for i, r := range s {
		ends[r] = i
	}

	partitons := []int{}

	var start, end int

	for i, r := range s {
		end = max(end, ends[r])

		if end == i {
			partitons = append(partitons, end-start+1)
			start, end = i+1, 0
		}
	}

	return partitons
}
