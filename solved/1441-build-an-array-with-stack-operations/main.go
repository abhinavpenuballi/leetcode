package main

func buildArray(target []int, n int) []string {
	required, max := map[int]bool{}, 0

	for _, val := range target {
		required[val] = true

		if val > max {
			max = val
		}
	}

	result := []string{}

	for i := 1; i <= n; i++ {
		if i > max {
			break
		}

		if required[i] {
			result = append(result, "Push")
		} else {
			result = append(result, "Push", "Pop")
		}
	}

	return result
}
