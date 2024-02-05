package main

func restoreString(s string, indices []int) string {
	t := make([]rune, len(s))

	for i, c := range s {
		t[indices[i]] = c
	}

	return string(t)
}
