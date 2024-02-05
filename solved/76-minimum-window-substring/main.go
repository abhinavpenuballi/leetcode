package main

func minWindow(s string, t string) string {
	occurrences := map[byte]int{}

	for i := 0; i < len(t); i++ {
		occurrences[t[i]]++
	}

	var l, r int
	var minWindow string

	for r < len(s) {
		if _, ok := occurrences[s[r]]; ok {
			occurrences[s[r]]--
		}

		if isValidWindow(occurrences) {
			for l <= r {
				if minWindow == "" || len(minWindow) > len(s[l:r+1]) {
					minWindow = s[l : r+1]
				}

				if _, ok := occurrences[s[l]]; ok {
					occurrences[s[l]]++

					if occurrences[s[l]] > 0 {
						l++
						break
					}
				}

				l++
			}
		}
		r++
	}

	return minWindow
}

func isValidWindow(occurrences map[byte]int) bool {
	for _, v := range occurrences {
		if v > 0 {
			return false
		}
	}

	return true
}
