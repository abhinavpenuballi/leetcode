package main

import "fmt"

func main() {
	findKthLargest([]int{1, 2, 3, 4, 5}, 2)
}

func findKthLargest(nums []int, k int) int {
	heap := nums[:k]

	for _, num := range nums {
		heap = insertIntoHeap(heap, num)
	}

	return 0
}

func insertIntoHeap(heap []int, num int) []int {
	fmt.Println(len(heap))
	return heap
}
