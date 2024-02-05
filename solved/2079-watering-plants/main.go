package main

func wateringPlants(plants []int, capacity int) int {
	steps := 0

	for currPos, currCapacity := -1, capacity; currPos < len(plants)-1; currPos++ {
		if plants[currPos+1] <= currCapacity {
			currCapacity -= plants[currPos+1]
			steps++
		} else {
			steps += 2*currPos + 3
			currCapacity = capacity - plants[currPos+1]
		}
	}

	return steps
}
