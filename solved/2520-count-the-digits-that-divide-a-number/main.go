package main

func countDigits(num int) (count int) {
	for numCopy := num; numCopy > 0; numCopy /= 10 {
		if num%(numCopy%10) == 0 {
			count++
		}
	}

	return count
}
