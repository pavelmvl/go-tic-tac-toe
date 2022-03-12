package field

import (
	"go-tic-tac-toe/internal/common"
)

func (f Field) isCellValid(col, row int) error {
	if col < 0 || col >= f.size {
		return common.ErrCellColumn
	}
	if row < 0 || row >= f.size {
		return common.ErrCellRow
	}
	return nil
}
