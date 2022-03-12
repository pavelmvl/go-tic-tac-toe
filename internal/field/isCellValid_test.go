package field

import (
	"github.com/stretchr/testify/assert"
	"go-tic-tac-toe/internal/common"
	"testing"
)

func TestisCellValid(t *testing.T) {
	size := 3
	winSeq := 3
	f, err := New(size, winSeq)
	assert.Equal(t, nil, err, "Create field error")
	assert.Equal(t, nil, f.isCellValid(0, 0), "Cell 0,0 should be valid")
	assert.Equal(t, common.ErrCellColumn, f.isCellValid(5, 0), "Cell 5,0 should be invalid")
	assert.Equal(t, common.ErrCellRow, f.isCellValid(0, 5), "Cell 0,5 should be invalid")
	assert.Equal(t, common.ErrCellColumn, f.isCellValid(5, 5), "Cell 5,5 should be invalid")
}
