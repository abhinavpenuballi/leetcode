package main

import "fmt"

func main() {
	str1, str2 := "ABCABC", "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	var shorterString string

	l1, l2 := len(str1), len(str2)

	if l1 < l2 {
		shorterString = str1
	} else {
		shorterString = str2
	}

	l := len(shorterString)

	for pick := l; pick > 0; pick-- {
		if !(l1%pick == 0 && l2%pick == 0) {
			continue
		}

		picked := shorterString[0:pick]

		if isGCD(str1, str2, picked) {
			return picked
		}
	}

	return ""
}

func isGCD(str1, str2, picked string) bool {
	l1, l2, l := len(str1), len(str2), len(picked)

	m1, m2 := l1/l, l2/l

	w1, w2 := "", ""

	for i := 0; i < m1; i++ {
		w1 += picked
	}

	for i := 0; i < m2; i++ {
		w2 += picked
	}

	return str1 == w1 && str2 == w2
}
