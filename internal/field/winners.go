package field

const NoWinner rune = rune(0)

func (f Field) IsColumnWinner(col, row int) rune {
	winner := f.GetCellValue(col, row)
	if winner == rune(0) {
		return NoWinner
	}
	lineSize := 0
	for v := 0; v < f.size; v++ {
		if f.GetCellValue(col, v) == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.size {
			return winner
		}
	}
	return NoWinner
}

func (f Field) IsRowWinner(col, row int) rune {
	winner := f.GetCellValue(col, row)
	if winner == rune(0) {
		return NoWinner
	}
	lineSize := 0
	for v := 0; v < f.size; v++ {
		if f.GetCellValue(v, row) == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.size {
			return winner
		}
	}
	return NoWinner
}

func (f Field) IsDiagStreightWinner(col, row int) rune {
	winner := f.GetCellValue(col, row)
	if winner == rune(0) {
		return NoWinner
	}
	lineSize := 0
	sub := 0
	if col > row {
		sub = row
	} else {
		sub = col
	}
	curCol := col - sub
	curRow := row - sub
	for curCol < f.size && curRow < f.size {
		if f.GetCellValue(curCol, curRow) == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.size {
			return winner
		}
		curCol++
		curRow++
	}
	return NoWinner
}

func (f Field) IsDiagReverseWinner(col, row int) rune {
	winner := f.GetCellValue(col, row)
	if winner == rune(0) {
		return NoWinner
	}
	lineSize := 0
	sub := col
	add := f.size - row - 1
	curCol := col
	curRow := row
	if sub < add {
		curCol -= sub
		curRow += sub
	} else {
		curCol -= add
		curRow += add
	}
	for curCol < f.size && curRow >= 0 {
		if f.GetCellValue(curCol, curRow) == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.size {
			return winner
		}
		curCol++
		curRow--
	}
	return NoWinner
}