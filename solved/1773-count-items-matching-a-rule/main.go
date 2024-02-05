package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	matches, ruleKeys := 0, map[string]int{"type": 0, "color": 1, "name": 2}

	for _, item := range items {
		if item[ruleKeys[ruleKey]] == ruleValue {
			matches++
		}
	}

	return matches
}
