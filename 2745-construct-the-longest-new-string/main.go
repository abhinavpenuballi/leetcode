package main

func longestString(x int, y int, z int) int {
	if x > y {
		return 2 * ((y + 1) + y + z)
	} else if y > x {
		return 2 * (x + (x + 1) + z)
	}

	return 2 * (x + y + z)
}
