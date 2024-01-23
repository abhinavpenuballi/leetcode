package main

func maxLength(arr []string) int {
	return next(arr, "", [26]bool{})
}

func next(arr []string, currStr string, chars [26]bool) int {
	maxLen := len(currStr)

	if len(arr) == 0 {
		return maxLen
	}

	for i, str := range arr {
		tempChars, valid := chars, true

		for _, char := range str {
			index := char - 'a'

			if tempChars[index] {
				valid = false
				break
			} else {
				tempChars[index] = true
			}
		}

		if !valid {
			continue
		}

		val := next(arr[i+1:], currStr+str, tempChars)

		if val > maxLen {
			maxLen = val
		}
	}

	return maxLen
}
