package main

func findErrorNums(nums []int) []int {
	hash := make([]int, len(nums)+1)

	for _, num := range nums {
		hash[num]++
	}

	var gotDuplicate, gotMissing bool

	ret := []int{0, 0}

	for i := 1; i < len(hash); i++ {
		switch hash[i] {
		case 2:
			ret[0] = i
			gotDuplicate = true
		case 0:
			ret[1] = i
			gotMissing = true
		}

		if gotDuplicate && gotMissing {
			break
		}
	}

	return ret
}
