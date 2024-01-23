package main

import (
	"fmt"
	"strings"
)

func garbageCollectionBkp(garbage []string, travel []int) int {
	timeTaken := 0

	for _, garbageType := range []string{"M", "P", "G"} {
		gotGarbage := false
		timeTakenForType := 0
		for i, garbageAtEachHouse := range garbage {
			count := strings.Count(garbageAtEachHouse, garbageType)
			if count > 0 {
				gotGarbage = true
				timeTakenForType += count

				if i < len(travel) {
					timeTakenForType += travel[i]
				}
			} else {

			}

		}

		if gotGarbage {
			timeTaken += timeTakenForType
		}
		fmt.Println(timeTaken)
	}

	return timeTaken
}

func garbageCollection(garbage []string, travel []int) int {
	garbageTill := map[string]int{"M": -1, "P": -1, "G": -1}
	garbageTypes := []string{"M", "P", "G"}

	for i, garbageAtEachHouse := range garbage {
		for _, garbageType := range garbageTypes {
			if strings.Contains(garbageAtEachHouse, garbageType) {
				garbageTill[garbageType] = i
			}
		}
	}

	timeTaken := 0

	for garbageType, till := range garbageTill {
		for i := 0; i <= till; i++ {
			timeTaken += strings.Count(garbage[i], garbageType)

			if i < len(travel) {
				timeTaken += travel[i]
			}
		}
	}

	return timeTaken
}
