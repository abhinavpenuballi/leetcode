package main

import (
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k > 1 {
		sByteArr := []byte(s)
		sort.Slice(sByteArr, func(i, j int) bool { return sByteArr[i] < sByteArr[j] })
		return string(sByteArr)
	} else {
		minChar := 'z'

		for _, char := range s {
			if char < minChar {
				minChar = char
			}
		}

		ss := []string{}

		for i, char := range s {
			if char == minChar {
				ss = append(ss, s[i:]+s[0:i])
			}
		}

		minStr := ss[0]

		for i := 1; i < len(ss); i++ {
			if ss[i] < minStr {
				minStr = ss[i]
			}
		}

		return minStr
	}
}
