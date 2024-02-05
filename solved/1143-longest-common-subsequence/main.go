package main

func longestCommonSubsequence(text1 string, text2 string) int {
	maxMatched, dp := 0, map[int]map[int]int{}

	for i1 := 0; i1 < len(text1); i1++ {
		for i2 := 0; i2 < len(text2); i2++ {
			if _, ok := dp[i1]; !ok {
				dp[i1] = make(map[int]int)
			}

			if text1[i1] == text2[i2] {
				dp[i1][i2] = dp[i1-1][i2-1] + 1
			} else {
				dp[i1][i2] = max(dp[i1-1][i2], dp[i1][i2-1])
			}

			maxMatched = max(maxMatched, dp[i1][i2])
		}
	}

	return maxMatched
}
