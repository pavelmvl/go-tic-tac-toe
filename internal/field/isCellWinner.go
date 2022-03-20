package field

import (
	"go-tic-tac-toe/internal/common"
)

func (f Field) IsCellWinner(col, row int) rune {
	return f.isCellWinner(col, row)
}

func (f Field) isCellWinner(col, row int) rune {
	var win rune
	win = f.isColumnWinner(col, row)
	if win != common.NoWinner {
		return win
	}
	win = f.isRowWinner(col, row)
	if win != common.NoWinner {
		return win
	}
	win = f.isDiagStreightWinner(col, row)
	if win != common.NoWinner {
		return win
	}
	win = f.isDiagReverseWinner(col, row)
	if win != common.NoWinner {
		return win
	}
	return common.NoWinner
}

func (f Field) isColumnWinner(col, row int) rune {
	winner, _ := f.GetCellValue(col, row)
	if winner == rune(0) {
		return common.NoWinner
	}
	lineSize := 0
	for v := 0; v < f.size; v++ {
		if w, _ := f.GetCellValue(col, v); w == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.winSeq {
			return winner
		}
	}
	return common.NoWinner
}

func (f Field) isRowWinner(col, row int) rune {
	winner, _ := f.GetCellValue(col, row)
	if winner == rune(0) {
		return common.NoWinner
	}
	lineSize := 0
	for v := 0; v < f.size; v++ {
		if w, _ := f.GetCellValue(v, row); w == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.winSeq {
			return winner
		}
	}
	return common.NoWinner
}

func (f Field) isDiagStreightWinner(col, row int) rune {
	winner, _ := f.GetCellValue(col, row)
	if winner == rune(0) {
		return common.NoWinner
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
		if w, _ := f.GetCellValue(curCol, curRow); w == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.winSeq {
			return winner
		}
		curCol++
		curRow++
	}
	return common.NoWinner
}

func (f Field) isDiagReverseWinner(col, row int) rune {
	winner, _ := f.GetCellValue(col, row)
	if winner == rune(0) {
		return common.NoWinner
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
		if w, _ := f.GetCellValue(curCol, curRow); w == winner {
			lineSize++
		} else {
			lineSize = 0
		}
		if lineSize >= f.winSeq {
			return winner
		}
		curCol++
		curRow--
	}
	return common.NoWinner
}
