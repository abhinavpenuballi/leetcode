package main

func decodeMessage(key string, message string) string {
	table, char := map[rune]rune{' ': ' '}, 'a'

	for _, k := range key {
		if _, ok := table[k]; !ok {
			table[k] = char
			char++
		}
	}

	decrypted := ""

	for _, m := range message {
		decrypted += string(table[m])
	}

	return decrypted
}
