package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	incomings := map[int]struct{}{}

	for _, edge := range edges {
		incomings[edge[1]] = struct{}{}
	}

	set := []int{}

	for vertex := 0; vertex < n; vertex++ {
		if _, ok := incomings[vertex]; !ok {
			set = append(set, vertex)
		}
	}

	return set
}
