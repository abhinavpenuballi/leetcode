package main

func minInsertions(s string) int {
	return len(s) - longestPalindromicSubsequence(s)
}

func longestPalindromicSubsequence(str string) int {
	rev := reverse(str)

	longest, dp := 0, map[int]map[int]int{}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(rev); j++ {
			if _, ok := dp[i]; !ok {
				dp[i] = make(map[int]int)
			}

			if str[i] == rev[j] {
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
