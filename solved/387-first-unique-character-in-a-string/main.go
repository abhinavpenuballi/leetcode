package main

func firstUniqChar(s string) int {
	occurrences, firsts := [26]int{}, []int{}

	for i, l := range s {
		occurrences[l-'a']++

		if occurrences[l-'a'] == 1 {
			firsts = append(firsts, i)
		}
	}

	for _, first := range firsts {
		if occurrences[s[first]-'a'] == 1 {
			return first
		}
	}

	return -1
}
