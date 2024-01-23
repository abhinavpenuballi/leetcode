package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	lettersArr := [26]int{}

	for _, val := range letters {
		lettersArr[val-'a']++
	}

	return next(words, lettersArr, score, "")
}

func next(words []string, letters [26]int, score []int, str string) int {
	if len(words) == 0 {
		return getScore(str, score)
	}

	max := getScore(str, score)

	for i := 0; i < len(words); i++ {
		valid, lettersCopy := true, letters

		for _, val := range words[i] {
			lettersCopy[val-'a']--

			if lettersCopy[val-'a'] < 0 {
				valid = false
				break
			}
		}

		if valid {
			val := next(words[i+1:], lettersCopy, score, str+words[i])

			if val > max {
				max = val
			}
		}
	}

	return max
}

func getScore(str string, score []int) int {
	currScore := 0

	for _, val := range str {
		currScore += score[val-'a']
	}

	return currScore
}
