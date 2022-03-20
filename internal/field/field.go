package field

import (
	"go-tic-tac-toe/internal/common"
)

type Field struct {
	size   int
	winSeq int
	cells  [][]rune
}

func New(size int, winSeq int) (Field, error) {
	f := Field{}
	if size < 3 {
		return f, common.ErrFieldSize
	}
	if winSeq < 3 || winSeq > size {
		return f, common.ErrWinSeqSize
	}
	f.size = size
	f.winSeq = winSeq
	f.cells = make([][]rune, size)
	for i := range f.cells {
		f.cells[i] = make([]rune, size)
	}
	return f, nil
}

func NewIField(size, winSeq int) (common.IField, error) {
	return New(size, winSeq)
}

func (f Field) GetSize() int {
	return f.size
}

func (f Field) GetWinSeq() int {
	return f.winSeq
}

func (f Field) isCellValid(col, row int) error {
	if col < 0 || col >= f.size {
		return common.ErrCellColumn
	}
	if row < 0 || row >= f.size {
		return common.ErrCellRow
	}
	return nil
}

func (f Field) isCellFree(col, row int) error {
	if f.cells[col][row] != rune(0) {
		return common.ErrCellBusy
	}
	return nil
}

func (f Field) AssignCell(col, row int, mark rune) error {
	var err error
	err = f.isCellValid(col, row)
	if err != nil {
		return err
	}
	err = f.isCellFree(col, row)
	if err != nil {
		return err
	}
	f.cells[col][row] = mark
	return nil
}

func (f Field) GetCellValue(col, row int) (rune, error) {
	var err error
	err = f.isCellValid(col, row)
	if err != nil {
		return 0, err
	}
	return f.cells[col][row], nil
}

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

func (f Field) IsFieldFull() bool {
	for col := 0; col < f.size; col++ {
		for row := 0; row < f.size; row++ {
			if winner, _ := f.GetCellValue(col, row); winner == common.NoWinner {
				return false
			}
		}
	}
	return true
}

func (f Field) ToString() string {
	buf := make([]byte, 0, f.size*(2*f.size+1)+1)
	for row := 0; row < f.size; row++ {
		buf = append(buf, []byte(" ")...)
		for col := 0; col < f.size; col++ {
			mark, _ := f.GetCellValue(col, row)
			if mark == common.NoWinner {
				mark = ' '
			}
			buf = append(buf, []byte(string(mark))...)
			buf = append(buf, []byte(" ")...)
		}
		buf = append(buf, []byte("\n")...)
	}
	return string(buf)
}

func (f Field) Print() {
	print(f.ToString())
}
