package game

import (
	"fmt"
	"go-tic-tac-toe/internal/common"
)

type Game struct {
	iter      int
	playerIdx int
	winner    *common.IPlayer
	draw      bool
	players   []common.IPlayer
	field     common.IField
}

func NewGame(f common.IField, p ...common.IPlayer) Game {
	g := Game{
		iter:      0,
		playerIdx: 0,
		winner:    nil,
		draw:      false,
		players:   p,
		field:     f,
	}
	return g
}

var (
	ErrWinnerExist = fmt.Errorf("WinnerExists")
	ErrNewWinner   = fmt.Errorf("new winner")
	ErrNoWinner    = fmt.Errorf("no winner")
)

func (g *Game) NextMove(col, row int) error {
	if g.winner != nil {
		return ErrWinnerExist
	}
	currentPlayer := &g.players[g.playerIdx]
	if err := g.field.AssignCell(col, row, (*currentPlayer).GetMark()); err != nil {
		return err
	}
	if winner := g.field.IsCellWinner(col, row); winner != 0 {
		g.winner = currentPlayer
		return nil
	}
	if g.field.IsFieldFull() {
		g.draw = true
		return nil
	}
	g.playerIdx++
	if g.playerIdx >= len(g.players) {
		g.playerIdx = 0
		g.iter++
	}
	return nil
}

func (g Game) GetWinnerString() (string, error) {
	if g.draw {
		return "friendship is winner", nil
	}
	if g.winner != nil {
		return fmt.Sprintf("%c is winner", (*g.winner).GetMark()), nil
	}
	return "", ErrNoWinner
}

func (g Game) ToHtml(s ...string) string {
	return g.field.ToHtml(s...)
}

func (g Game) ToString() string {
	return g.field.ToString()
}

func (g Game) GetIter() int {
	return g.iter
}

func (g Game) GetCurrentPlayer() common.IPlayer {
	return g.players[g.playerIdx]
}
