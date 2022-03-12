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
