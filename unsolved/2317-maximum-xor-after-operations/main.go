package main

func maximumXOR(nums []int) int {
	ans := 0

	next(getPossibles(nums), 0, &ans)

	return ans
}

func next(possibles [][]int, xor int, ans *int) {
	if len(possibles) == 0 {
		if xor > *ans {
			*ans = xor
		}
		return
	}

	for i := 0; i < len(possibles[0]); i++ {
		next(possibles[1:], xor^possibles[0][i], ans)
	}
}

func getPossibles(nums []int) [][]int {
	possibles := [][]int{}

	for _, num := range nums {
		possibleMap := map[int]struct{}{}

		for i := 0; i <= num; i++ {
			possibleMap[num&(num^i)] = struct{}{}
		}

		possibleArr := []int{}

		for i := range possibleMap {
			possibleArr = append(possibleArr, i)
		}

		possibles = append(possibles, possibleArr)
	}

	return possibles
}
