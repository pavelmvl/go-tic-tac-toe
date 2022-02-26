package common

import (
	"errors"
)

type IField interface {
	GetSize() int
	GetWinSeq() int
	AssignCell(col, row int, mark rune) error
	GetCellValue(col, row int) (rune, error)
	IsCellWinner(col, row int) rune
	IsFieldFull() bool
	ToString() string
}

const (
	NoWinner = rune(0)
)

var (
	ErrFieldSize  = errors.New("Incorrect 'size' value. 'size' should be greater than '2'")
	ErrWinSeqSize = errors.New("Incorrect 'winSeq' value. 'winSeq' should be greater than '2' and less or equal 'size'")
	ErrCellColumn = errors.New("Incorrect cell: column check failed")
	ErrCellRow    = errors.New("Incorrect cell: row check failed")
	ErrCellBusy   = errors.New("Cell is not free")
)
