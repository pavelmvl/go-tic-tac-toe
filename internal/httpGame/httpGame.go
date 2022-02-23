package httpGame

import (
	"fmt"
	"go-tic-tac-toe/internal/common"
	"net/http"
	"strconv"
	"strings"
)

type HttpGame struct {
	game common.IGame
}

func (h HttpGame) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, h.game.ToHtml())
		return
	case "/new":
		fmt.Fprint(w, "TODO new()")
		return
	default:
		winnerString, winnerErr := h.game.GetWinnerString()
		if winnerErr == nil {
			fmt.Fprint(w, h.game.ToHtml(winnerString))
			return
		}
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
	winnerString, winnerErr := h.game.GetWinnerString()
	if winnerErr == nil {
		fmt.Fprint(w, h.game.ToHtml(winnerString))
		return
	}
	fmt.Fprint(w, h.game.ToHtml())
}

func NewHttpGame(g common.IGame) error {
	handle := HttpGame{
		game: g,
	}
	err := http.ListenAndServe(":8080", handle)
	return err
}
