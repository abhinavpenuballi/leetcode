package main

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := map[rune]struct{}{}

	for _, jewel := range jewels {
		jewelsMap[jewel] = struct{}{}
	}

	num := 0

	for _, stone := range stones {
		if _, ok := jewelsMap[stone]; ok {
			num++
		}
	}

	return num
}
