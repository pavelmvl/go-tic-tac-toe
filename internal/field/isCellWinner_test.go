package field

import (
	"github.com/stretchr/testify/assert"
	//"go-tic-tac-toe/internal/common"
	"testing"
)

func TestIsCellWinnerDiag33Streight(t *testing.T) {
	f, err := New(3, 3)
	assert.Nil(t, err)
	f.AssignCell(0, 0, 'X')
	f.AssignCell(1, 1, 'X')
	f.AssignCell(2, 2, 'X')
	assert.Equal(t, 'X', f.isCellWinner(0, 0))
	assert.Equal(t, 'X', f.isCellWinner(1, 1))
	assert.Equal(t, 'X', f.isCellWinner(2, 2))
	assert.NotEqual(t, 'X', f.isCellWinner(0, 2))
}

func TestIsCellWinnerDiag33Reverse(t *testing.T) {
	f, err := New(3, 3)
	assert.Nil(t, err)
	f.AssignCell(0, 2, 'X')
	f.AssignCell(1, 1, 'X')
	f.AssignCell(2, 0, 'X')
	assert.Equal(t, 'X', f.isCellWinner(0, 2))
	assert.Equal(t, 'X', f.isCellWinner(1, 1))
	assert.Equal(t, 'X', f.isCellWinner(0, 2))
	assert.NotEqual(t, 'X', f.isCellWinner(2, 2))
}

func TestIsCellWinnerDiag33Column(t *testing.T) {
	f, err := New(3, 3)
	assert.Nil(t, err)
	f.AssignCell(0, 0, 'X')
	f.AssignCell(0, 1, 'X')
	f.AssignCell(0, 2, 'X')
	assert.Equal(t, 'X', f.isCellWinner(0, 0))
	assert.Equal(t, 'X', f.isCellWinner(0, 1))
	assert.Equal(t, 'X', f.isCellWinner(0, 2))
	assert.NotEqual(t, 'X', f.isCellWinner(2, 2))
}

func TestIsCellWinnerDiag33Row(t *testing.T) {
	f, err := New(3, 3)
	assert.Nil(t, err)
	f.AssignCell(0, 2, 'X')
	f.AssignCell(1, 2, 'X')
	f.AssignCell(2, 2, 'X')
	assert.Equal(t, 'X', f.isCellWinner(0, 2))
	assert.Equal(t, 'X', f.isCellWinner(1, 2))
	assert.Equal(t, 'X', f.isCellWinner(2, 2))
	assert.NotEqual(t, 'X', f.isCellWinner(0, 0))
}
