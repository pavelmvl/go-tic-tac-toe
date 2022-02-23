package common

type IField interface {
	AssignCell(col, row int, mark rune) error
	IsCellWinner(col, row int) rune
	IsFieldFull() bool
	ToString() string
	ToHtml(...string) string
}
