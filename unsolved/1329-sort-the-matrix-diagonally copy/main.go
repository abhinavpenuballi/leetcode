package main

func diagonalSort(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])

	for iRow, iCol, iDiag, iDone := next(rows, cols, rows, -1, rows+1); !iDone; iRow, iCol, iDiag, iDone = next(rows, cols, iRow, iCol, iDiag) {
		for jRow, jCol, jDiag, jDone := next(rows, cols, iRow, iCol, iDiag); !jDone; jRow, jCol, jDiag, jDone = next(rows, cols, jRow, jCol, jDiag) {
			if jDiag != iDiag {
				break
			}

			if mat[iRow][iCol] > mat[jRow][jCol] {
				mat[iRow][iCol], mat[jRow][jCol] = mat[jRow][jCol], mat[iRow][iCol]
			}
		}
	}

	return mat
}

func next(rows, cols, row, col, diag int) (int, int, int, bool) {
	if row == 0 && col == cols-1 {
		return -1, -1, -1, true
	}

	row++
	col++

	if row >= rows || col >= cols {
		if diag > 0 {
			row = diag - 1
			col = 0
		} else {
			col = -diag + 1
			row = 0
		}

		diag = row - col
	}

	return row, col, diag, false
}
