package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ss := []string{"3[a]2[bc]", "3[a2[c]]", "2[abc]3[cd]ef", "abc3[cd]xyz", "3[z]2[2[y]pq4[2[jk]e1[f]]]ef", "sd2[f2[e]g]i"}

	for _, s := range ss {
		fmt.Println(decodeString(s))
	}
}

type stack struct {
	numStr string
	str    string
}

func decodeString(s string) string {
	stacks := []stack{}
	stackIndex := -1

	str := ""

	numStr := ""

	encounteredStr := ""

	var prev byte

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			if stackIndex >= 0 {
				stacks[stackIndex].str += encounteredStr
				encounteredStr = ""
			}

			numStr += string(s[i])
		} else if s[i] == '[' {
			stackIndex++

			if len(stacks) <= stackIndex {
				stacks = append(stacks, stack{})
			}

			stacks[stackIndex].numStr = numStr

			numStr = ""
		} else if 'a' <= s[i] && s[i] <= 'z' {
			if stackIndex == -1 {
				str += string(s[i])
			} else {
				encounteredStr += string(s[i])
			}
		} else if s[i] == ']' {
			if stackIndex >= 0 && prev != ']' {
				stacks[stackIndex].str += encounteredStr
				encounteredStr = ""
			}

			num, _ := strconv.Atoi(stacks[stackIndex].numStr)

			currStr := strings.Repeat(stacks[stackIndex].str, num)

			if stackIndex == 0 {
				str += currStr
			} else {
				stacks[stackIndex-1].str += currStr
			}

			stacks = stacks[:stackIndex]

			stackIndex--
		}

		prev = s[i]
	}

	return str
}
