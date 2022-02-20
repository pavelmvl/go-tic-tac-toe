package httpServer

import (
	"fmt"
	"net/http"
)

type IPlayer interface {
	GetMark() rune
}

type IField interface {
	ToHtml() string
}

type tictactoeHandler struct {
	col     int
	row     int
	iter    int
	players []IPlayer
	field   IField
}

func (t tictactoeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, t.field.ToHtml())
}

func NewHttpGame(f IField, p ...IPlayer) error {
	t := tictactoeHandler{
		field:   f,
		players: make([]IPlayer, 0, 2),
	}
	for _, v := range p {
		t.players = append(t.players, v)
	}
	err := http.ListenAndServe(":8080", t)
	return err
}
