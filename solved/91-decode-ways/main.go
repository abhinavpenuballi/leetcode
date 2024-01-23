package main

import (
	"fmt"
	"strconv"
)

func main() {
	ss := []string{"21203110231012", "123", "12", "226", "12123", "12012123", "111111111111111111111111111111111111111111111"}

	for _, s := range ss {
		fmt.Println(numDecodings(s))
	}
}

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	line := 1
	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			n, _ := strconv.Atoi(s[i : j+1])
			if 1 <= n && n <= 26 {
				dp[i] += dp[j+1]
			} else {
				break
			}
			line++
		}
	}
	return dp[0]
}
