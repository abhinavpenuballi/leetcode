package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userMinutes := map[int]map[int]struct{}{}

	for _, log := range logs {
		if _, ok := userMinutes[log[0]]; !ok {
			userMinutes[log[0]] = make(map[int]struct{})
		}
		userMinutes[log[0]][log[1]] = struct{}{}
	}

	UAMUsers := map[int]int{}

	for _, minutes := range userMinutes {
		UAMUsers[len(minutes)]++
	}

	ret := []int{}

	for i := 1; i <= k; i++ {
		ret = append(ret, UAMUsers[i])
	}

	return ret
}
