package main

import "sort"

func sortVowels(s string) string {
	rs := []rune(s)

	isVowel := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	vowels := []rune{}

	for _, val := range rs {
		if isVowel[val] {
			vowels = append(vowels, val)
		}
	}

	sort.Slice(vowels, func(i, j int) bool { return vowels[i] < vowels[j] })

	index := 0

	for i, val := range rs {
		if isVowel[val] {
			rs[i] = vowels[index]
			index++
		}
	}

	return string(rs)
}
