package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	abc := [][3]int{
		{2, 6, 5},
		{4, 2, 7},
		{1, 2, 3},
	}

	for _, abc := range abc {
		fmt.Println(minFlips(abc[0], abc[1], abc[2]))
	}
}

func minFlips(a int, b int, c int) int {
	aS := strconv.FormatInt(int64(a), 2)
	bS := strconv.FormatInt(int64(b), 2)
	cS := strconv.FormatInt(int64(c), 2)

	maxLen := max(len(aS), len(bS), len(cS))

	aS = strings.Repeat("0", maxLen-len(aS)) + aS
	bS = strings.Repeat("0", maxLen-len(bS)) + bS
	cS = strings.Repeat("0", maxLen-len(cS)) + cS

	flips := 0

	for i := 0; i < maxLen; i++ {
		if cS[i] == '1' {
			if aS[i] == '0' && bS[i] == '0' {
				flips++
			}
		} else {
			if aS[i] == '1' {
				flips++
			}
			if bS[i] == '1' {
				flips++
			}
		}
	}

	return flips
}
