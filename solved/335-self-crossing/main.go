package main

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func isSelfCrossing(distance []int) bool {
	dirs := [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

	p, lines := point{}, [5]line{}

	for move, dist := range distance {
		dir := dirs[move%4]

		start := p
		p.x += dist * dir[0]
		p.y += dist * dir[1]
		end := p

		line := line{start, end}

		if move >= 5 && crosses(lines[0], line, true) {
			return true
		} else if move >= 4 && crosses(lines[1], line, false) {
			return true
		} else if move >= 3 && crosses(lines[2], line, true) {
			return true
		}

		lines = append(lines, line)
	}

	return false
}

func append(lines [5]line, line line) [5]line {
	lines[0], lines[1], lines[2], lines[3], lines[4] = lines[1], lines[2], lines[3], lines[4], line
	return lines
}

func crosses(line1, line2 line, perpendicular bool) bool {
	if perpendicular {
		if line1.start.x != line1.end.x {
			line1, line2 = line2, line1
		}

		if line1.start.y > line1.end.y {
			line1.start, line1.end = line1.end, line1.start
		}
		if line2.start.x > line2.end.x {
			line2.start, line2.end = line2.end, line2.start
		}

		if line1.start.y <= line2.start.y && line2.start.y <= line1.end.y &&
			line2.start.x <= line1.start.x && line1.start.x <= line2.end.x {
			return true
		}
	} else {
		if line1.start.x == line2.start.x {
			if line1.start.x > line1.end.x {
				line1.start, line1.end = line1.end, line1.start
			}
			if line2.start.x > line2.end.x {
				line2.start, line2.end = line2.end, line2.start
			}

			if line1.start.y > line2.start.y {
				line1, line2 = line2, line1
			}

			if line1.end.y >= line2.start.y {
				return true
			}
		} else if line1.start.y == line2.start.y {
			if line1.start.y > line1.end.y {
				line1.start, line1.end = line1.end, line1.start
			}
			if line2.start.y > line2.end.y {
				line2.start, line2.end = line2.end, line2.start
			}

			if line1.start.x > line2.start.x {
				line1, line2 = line2, line1
			}

			if line1.end.x >= line2.start.x {
				return true
			}
		}
	}

	return false
}
