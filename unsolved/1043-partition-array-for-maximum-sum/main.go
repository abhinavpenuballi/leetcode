package main

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := map[int]int{}

	for i := 1; i <= len(arr); i++ {
		for j, maxSeen := 1, 0; j <= k && i-j >= 0; j++ {
			maxSeen = max(maxSeen, arr[i-j])

			dp[i] = max(dp[i], dp[i-j]+j*maxSeen)
		}
	}

	return dp[len(arr)]
}
