package main

func longestValidParentheses(s string) int {
	longest := 0

	for i := 0; i < len(s); i++ {
		for j, stack := i, 0; j < len(s); j++ {
			switch s[j] {
			case '(':
				stack++
			case ')':
				stack--
			}

			if stack < 0 {
				break
			} else if stack == 0 {
				longest = max(longest, j-i+1)
			}
		}
	}

	return longest
}
