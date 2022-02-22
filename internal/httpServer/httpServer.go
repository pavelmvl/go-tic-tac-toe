package httpServer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type IPlayer interface {
	GetMark() rune
}

type IField interface {
	AssignCell(col, row int, mark rune) error
	ToHtml() string
}

type tictactoeHandler struct {
	iter      *int
	playerIdx *int
	players   []IPlayer
	field     IField
}

func (t tictactoeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" {
		fmt.Fprint(w, t.field.ToHtml())
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

	if err := t.field.AssignCell(col, row, t.players[*t.playerIdx].GetMark()); err != nil {
		http.Error(w, "500 internal server error", 500)
		return
	}
	*t.playerIdx++
	if *t.playerIdx >= len(t.players) {
		*t.playerIdx = 0
		*t.iter++
	}
	fmt.Println(*t.iter, *t.playerIdx, t.players, t.field)

	fmt.Fprint(w, t.field.ToHtml())
}

func NewHttpGame(f IField, p ...IPlayer) error {
	t := tictactoeHandler{
		iter:      new(int),
		playerIdx: new(int),
		field:     f,
		players:   make([]IPlayer, 0, 2),
	}
	for _, v := range p {
		t.players = append(t.players, v)
	}
	err := http.ListenAndServe(":8080", t)
	return err
}
