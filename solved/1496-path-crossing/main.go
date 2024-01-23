package main

import "fmt"

func main() {
	paths := []string{
		"NES",
		"NESWW",
	}

	for _, path := range paths {
		fmt.Println(isPathCrossing(path))
	}
}

type point [2]int

func isPathCrossing(path string) bool {
	visited := map[point]bool{}
	location := point{0, 0}
	visited[location] = true

	for _, direction := range path {
		location = move(location, direction)

		if visited[location] {
			return true
		}

		visited[location] = true
	}

	return false
}

func move(location point, direction rune) point {
	switch direction {
	case 'N':
		location[1]++
	case 'E':
		location[0]++
	case 'S':
		location[1]--
	case 'W':
		location[0]--
	}

	return location
}
