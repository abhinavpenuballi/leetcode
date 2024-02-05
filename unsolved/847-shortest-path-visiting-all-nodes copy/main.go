package main

func shortestPathLength(graph [][]int) int {
	maxHops := 1000

	for i := range graph {
		maxHops = min(maxHops, next(graph, i, map[int]struct{}{}, 0))
	}

	return maxHops
}

func next(graph [][]int, currNode int, visited map[int]struct{}, hops int) int {
	if _, ok := visited[currNode]; ok {
		return 1000
	}

	visitedCopy := copyMap(visited)

	visitedCopy[currNode] = struct{}{}

	if len(visitedCopy) == len(graph) {
		return hops
	}

	maxHops := 1000

	for _, node := range graph[currNode] {
		maxHops = min(maxHops, next(graph, node, visited, hops+1))
	}

	return maxHops
}

func copyMap(in map[int]struct{}) map[int]struct{} {
	out := map[int]struct{}{}

	for k, v := range in {
		out[k] = v
	}

	return out
}
