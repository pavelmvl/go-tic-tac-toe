package field

import (
	"github.com/stretchr/testify/assert"
	"go-tic-tac-toe/internal/common"
	"testing"
)

func TestIsCellFree(t *testing.T) {
	size := 3
	winSeq := 3
	f, err := New(size, winSeq)
	assert.Equal(t, nil, err, "Create field fail")
	f.AssignCell(0, 0, 'X')
	assert.Equal(t, common.ErrCellBusy, f.isCellFree(0, 0), "Cell 0,0 should be non free")
	assert.Equal(t, nil, f.isCellFree(0, 1), "Cell 0,1 should be free")
}
