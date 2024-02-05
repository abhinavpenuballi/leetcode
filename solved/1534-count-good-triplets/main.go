package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[k]-arr[i]) <= c {
					count++
				}
			}
		}
	}

	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
