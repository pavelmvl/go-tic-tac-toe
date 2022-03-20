package httpGame

import (
	"fmt"
	"go-tic-tac-toe/internal/common"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrFuncNil = fmt.Errorf("Func should be nil")
)

type (
	FuncNewIField  func(int, int) (common.IField, error)
	FuncNewIPlayer func(string) common.IPlayer
)

type HttpGame struct {
	game       common.IGame
	newIField  FuncNewIField
	newIPlayer FuncNewIPlayer
}

func (h HttpGame) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	switch r.URL.Path {
	case "/":
		if h.game == nil {
			fmt.Fprint(w, h.game.ToHtml())
			return
		}
	case "/new":
		values := r.URL.Query()
		if len(values) > 0 {
			size, okSize := values["fieldSize"]
			winSeq, okWinSeq := values["winSeq"]
			side, okSide := values["side"]
			if okSize && okWinSeq && okSide {
				sizeInt, errSizeInt := strconv.Atoi(size[0])
				if errSizeInt != nil {
					fmt.Fprint(w, errSizeInt)
					return
				}
				winSeqInt, errWinSeqInt := strconv.Atoi(winSeq[0])
				if errSizeInt != nil {
					fmt.Fprint(w, errWinSeqInt)
					return
				}
				f, errF := h.newIField(sizeInt, winSeqInt)
				if errF != nil {
					fmt.Fprint(w, errF)
					return
				}
				firstPlayer := h.newIPlayer(side[0])
				var secondSide string
				switch firstPlayer.GetMark() {
				case 'X':
					secondSide = "O"
				default:
					secondSide = "X"
				}
				secondPlayer := h.newIPlayer(secondSide)
				h.game.Renew(f, firstPlayer, secondPlayer)
				fmt.Fprint(w, h.game.ToHtml())
			}
		} else {
			fmt.Fprint(w, newHtml)
		}
		return
	default:
		winnerString, winnerErr := h.game.GetWinnerString()
		if winnerErr == nil {
			fmt.Fprint(w, h.game.ToHtml(winnerString))
			return
		}
		coords := strings.SplitN(r.URL.Path, "/", 3)
		if len(coords) < 3 {
			http.Error(w, "400 bad request", 400)
			return
		}
		col, errCol := strconv.Atoi(coords[1])
		if errCol != nil {
			http.Error(w, "400 bad request", 400)
			return
		}
		row, errRow := strconv.Atoi(coords[2])
		if errRow != nil {
			http.Error(w, "400 bad request", 400)
			return
		}
		moveErr := h.game.NextMove(col, row)
		if moveErr != nil {
			fmt.Fprint(w, h.game.ToHtml(moveErr.Error()))
			return
		}
		winnerString, winnerErr = h.game.GetWinnerString()
		if winnerErr == nil {
			fmt.Fprint(w, h.game.ToHtml(winnerString))
			return
		}
	}
	fmt.Fprint(w, h.game.ToHtml())
}

func NewHttpGame(g common.IGame, newIField FuncNewIField, newIPlayer FuncNewIPlayer) error {
	if newIField == nil {
		return ErrFuncNil
	}
	if newIPlayer == nil {
		return ErrFuncNil
	}
	handle := HttpGame{
		game:       g,
		newIField:  newIField,
		newIPlayer: newIPlayer,
	}
	err := http.ListenAndServe(":8080", handle)
	return err
}
