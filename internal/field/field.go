package field

import (
	"fmt"
)

type Field struct {
	size int
}

func New(size int) (Field, error) {
	if size > 2 {
		return Field{
			size: size,
		}, nil
	}
	return Field{}, fmt.Errorf("Incorrect 'size' value (%d). 'size' should be greater than '2'", size)
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
