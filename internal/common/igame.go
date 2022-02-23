package common

type IGame interface {
	GetWinnerString() (string, error)
	NextMove(int, int) error
	ToHtml(...string) string
}
