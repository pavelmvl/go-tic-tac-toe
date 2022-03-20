package common

type IGame interface {
	Renew(IField, ...IPlayer)
	GetWinnerString() (string, error)
	NextMove(int, int) error
	GetIter() int
	GetFieldSize() int
	GetCurrentPlayer() IPlayer
	ToString() string
	ToHtml(...string) string
}
