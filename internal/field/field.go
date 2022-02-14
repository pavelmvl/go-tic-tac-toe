package field

import (
	"fmt"
)

type Field struct {
	size  int
	cells [][]rune
}

func New(size int) (Field, error) {
	f := Field{}
	if size < 3 {
		return f, fmt.Errorf("Incorrect 'size' value (%d). 'size' should be greater than '2'", size)
	}
	f.size = size
	f.cells = make([][]rune, size)
	for i := range f.cells {
		f.cells[i] = make([]rune, size)
	}
	return f, nil
}

func (f Field) IsCellValid(col, row int) error {
	if col < 0 || col >= f.size {
		return fmt.Errorf("Incorrect cell: column check failed")
	}
	if row < 0 || row >= f.size {
		return fmt.Errorf("Incorrect cell: row check failed")
	}
	return nil
}

func (f Field) IsCellFree(col, row int) error {
	if f.cells[col][row] != rune(0) {
		return fmt.Errorf("Cell is not free (%c)", f.cells[col][row])
	}
	return nil
}

func (f *Field) AssignCell(col, row int, mark rune) error {
	f.cells[col][row] = mark
	return nil
}

func (f Field) GetCellValue(col, row int) rune {
	return f.cells[col][row]
}
