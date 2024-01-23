package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"zza",
		"bac",
		"bdda",
		"abc",
		"mmuqezwmomeplrtskz",
	}

	for _, test := range tests {
		fmt.Println(robotWithString(test))
	}
}

func robotWithString(s string) string {
	minChar := getMinChar(s, "")

	dp := map[string]string{",": strings.Repeat("z", 100001)}

	return next(s, "", "", minChar, &dp)
}

func next(s, t, paper string, minChar byte, dp *map[string]string) string {
	state := s + "," + t

	/* if val, ok := (*dp)[state]; ok {
		(*dp)[state] = min(val, paper)
		return (*dp)[state]
	} */

	if s == "" && t == "" {
		(*dp)[state] = paper
		return (*dp)[state]
	}

	if !strings.Contains(s, string(minChar)) && !strings.Contains(t, string(minChar)) {
		minChar = getMinChar(s, t)
	}

	if s == "" {
		(*dp)[state] = moveFromTToPaper(s, t, paper, minChar, dp)
		return (*dp)[state]
	}

	if t == "" {
		(*dp)[state] = moveFromSToT(s, t, paper, minChar, dp)
		return (*dp)[state]
	}

	if t[len(t)-1] == minChar {
		(*dp)[state] = moveFromTToPaper(s, t, paper, minChar, dp)
		return (*dp)[state]
	}

	if s[0] == minChar {
		(*dp)[state] = moveFromSToT(s, t, paper, minChar, dp)
		return (*dp)[state]
	}

	(*dp)[state] = min(
		moveFromSToT(s, t, paper, minChar, dp),
		moveFromTToPaper(s, t, paper, minChar, dp),
	)
	return (*dp)[state]
}

func getMinChar(s, t string) byte {
	minChar := byte('z')

	for i := 0; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] < minChar {
			minChar = t[i]
		}
	}

	return minChar
}

func moveFromSToT(s, t, paper string, minChar byte, dp *map[string]string) string {
	return next(s[1:], t+string(s[0]), paper, minChar, dp)
}

func moveFromTToPaper(s, t, paper string, minChar byte, dp *map[string]string) string {
	return next(s, t[:len(t)-1], paper+string(t[len(t)-1]), minChar, dp)
}
