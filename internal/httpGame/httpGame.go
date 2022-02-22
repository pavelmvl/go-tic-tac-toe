package httpGame

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type IGame interface {
	GetWinnerString() (string, error)
	NextMove(int, int) error
	ToHtml(...string) string
}

type HttpGame struct {
	game IGame
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
		winnerString, winnerErr := h.game.GetWinnerString()
		if winnerErr == nil {
			fmt.Fprint(w, h.game.ToHtml(moveErr.Error(), winnerString))
			return
		}
		fmt.Fprint(w, h.game.ToHtml(moveErr.Error()))
		return
	}
	fmt.Fprint(w, h.game.ToHtml())
}

func NewHttpGame(g IGame) error {
	handle := HttpGame{
		game: g,
	}
	err := http.ListenAndServe(":8080", handle)
	return err
}
