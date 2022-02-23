package common

type IGame interface {
	GetWinnerString() (string, error)
	NextMove(int, int) error
	GetIter() int
	GetCurrentPlayer() IPlayer
	ToString() string
	ToHtml(...string) string
}
