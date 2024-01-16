package main

import "fmt"

func main() {
	ns := []int{1, 2, 3, 10, 100}

	for _, n := range ns {
		fmt.Println(divisorGame(n))
	}
}

func divisorGame(n int) bool {
	return alice(n, &map[int]bool{}, &map[int]bool{})
}

func alice(n int, adp, bdp *map[int]bool) bool {
	if val, ok := (*adp)[n]; ok {
		return val
	}

	for i := 1; i < n; i++ {
		if n%i == 0 {
			(*bdp)[n-i] = bob(n-i, adp, bdp)

			if (*bdp)[n-i] {
				return true
			}
		}
	}

	(*adp)[n] = false

	return false
}

func bob(n int, adp, bdp *map[int]bool) bool {
	if val, ok := (*bdp)[n]; ok {
		return val
	}

	for i := 1; i < n; i++ {
		if n%i == 0 {
			(*adp)[n-i] = alice(n-i, adp, bdp)

			if !(*adp)[n-i] {
				return false
			}
		}
	}

	(*bdp)[n] = false

	return true
}
