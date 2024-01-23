package main

import (
	"sort"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	type pair struct{ plantTime, growTime int }

	pairs := []pair{}

	for i := 0; i < len(plantTime); i++ {
		pairs = append(pairs, pair{plantTime[i], growTime[i]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].growTime > pairs[j].growTime
	})

	var plantTimeSum, lastBloomTime int

	for _, val := range pairs {
		plantTimeSum += val.plantTime
		lastBloomTime = max(lastBloomTime, plantTimeSum+val.growTime)
	}

	return lastBloomTime
}
