package main

import "slices"

func main()

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	contentChildren := 0

	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			contentChildren++
			i++
		}
	}

	return contentChildren
}
