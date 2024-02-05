package main

func cellsInRange(s string) (cells []string) {
	for col := s[0]; col <= s[3]; col++ {
		for row := s[1]; row <= s[4]; row++ {
			cells = append(cells, string(col)+string(row))
		}
	}

	return cells
}
