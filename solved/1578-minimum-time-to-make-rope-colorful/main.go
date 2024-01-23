package main

import "fmt"

func main() {
	var colors string
	var neededTime []int

	colors = "abaac"
	neededTime = []int{1, 2, 3, 4, 5}

	minCost := minCost(colors, neededTime)

	fmt.Println(minCost)
}

func minCost(colors string, neededTime []int) int {
	prev := colors[0]
	from, to := 0, 0

	minCost := 0

	for i := 1; i < len(colors); i++ {
		if prev == colors[i] {
			to = i
			continue
		}

		if to > from {
			minCost += getLowestCosts(neededTime, from, to)
		}

		prev = colors[i]
		from, to = i, i
	}

	if to > from {
		minCost += getLowestCosts(neededTime, from, to)
	}

	return minCost
}

func getLowestCosts(neededTime []int, from, to int) int {
	var max, cost int

	for i := from; i <= to; i++ {
		cost += neededTime[i]

		if max < neededTime[i] {
			max = neededTime[i]
		}
	}

	cost -= max

	return cost
}
