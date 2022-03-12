package field

import (
	"go-tic-tac-toe/internal/common"
)

func (f Field) isCellFree(col, row int) error {
	if f.cells[col][row] != rune(0) {
		return common.ErrCellBusy
	}
	return nil
}
