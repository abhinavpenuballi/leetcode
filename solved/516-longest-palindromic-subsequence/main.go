package main

func longestPalindromeSubseq(s string) int {
	r := reverse(s)

	longest, dp := 0, map[int]map[int]int{}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(r); j++ {
			if _, ok := dp[i]; !ok {
				dp[i] = make(map[int]int)
			}

			if s[i] == r[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

			longest = max(longest, dp[i][j])
		}
	}

	return longest
}

func reverse(s string) string {
	r := ""

	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return r
}
