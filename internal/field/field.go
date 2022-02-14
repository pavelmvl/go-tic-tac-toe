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
