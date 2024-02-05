package main

func countRoutes(locations []int, start int, finish int, fuel int) int {
	return next(locations, start, finish, fuel, &map[int]map[int]int{}) % 1000000007
}

func next(locations []int, curr, finish, fuel int, dp *map[int]map[int]int) int {
	if val, ok := (*dp)[curr][fuel]; ok {
		return val
	}

	if fuel < 0 {
		return 0
	}

	paths := 0

	if curr == finish {
		paths++
	}

	for i := 0; i < curr; i++ {
		paths += next(locations, i, finish, fuel-abs(locations[curr]-locations[i]), dp) % 1000000007
	}

	for i := curr + 1; i < len(locations); i++ {
		paths += next(locations, i, finish, fuel-abs(locations[curr]-locations[i]), dp) % 1000000007
	}

	if _, ok := (*dp)[curr]; !ok {
		(*dp)[curr] = make(map[int]int)
	}

	(*dp)[curr][fuel] = paths
	return paths
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
