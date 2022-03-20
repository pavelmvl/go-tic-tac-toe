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

func (g *Game) Renew(f common.IField, p ...common.IPlayer) {
	g.iter = 0
	g.playerIdx = 0
	g.winner = nil
	g.draw = false
	g.players = p
	g.field = f
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

var htmlTemplate string = `<html>
<head>
	<title>Tic-tac-toe</title>
	<style>
.field { display:table; outline:2px solid black; border-collapse:collapse; }
.row { display:table-row; }
.cell { display:table-cell; outline:1px solid black; border-collapse:collapse; margin:0; padding:0; width:%dpx; height:%dpx; font-size:50px; text-align: center; vertical-align:middle; }
	</style>
</head>
<body>
	<div class="config">
		<p>
			Fisrt player: %c<br/>
			Second player: %c<br/>
			Win sequence length: %d<br/>
		</p>
		<p><button onclick="location.href='/new'">New game</button></p>
	</div>
	<div class="field">%s</div>
	<div class="extra">%s</div>
</body>
</html>`

func (g Game) ToHtml(extra ...string) string {
	extraByte := make([]byte, 0, 512)
	for _, v := range extra {
		extraByte = append(extraByte, []byte(v)...)
		extraByte = append(extraByte, []byte("<br/>")...)
	}
	table := make([]byte, 0, 512)
	for row := 0; row < g.GetFieldSize(); row++ {
		table = append(table, []byte("<div class=\"row\">")...)
		for col := 0; col < g.GetFieldSize(); col++ {
			mark, _ := g.field.GetCellValue(col, row)
			if mark == common.NoWinner {
				mark = ' '
			}
			table = append(table, []byte(fmt.Sprintf("<div class=\"cell\" onclick=\"location.href='/%d/%d';\">%c</div>", col, row, mark))...)
		}
		table = append(table, []byte("</div>")...)
	}
	// TODO width = 100, height = 100. Make it parametrized
	return fmt.Sprintf(
		htmlTemplate,
		100,
		100,
		rune(g.players[0].GetMark()),
		rune(g.players[1].GetMark()),
		g.field.GetWinSeq(),
		string(table),
		string(extraByte),
	)
}

func (g Game) ToString() string {
	return g.field.ToString()
}

func (g Game) GetIter() int {
	return g.iter
}

func (g Game) GetFieldSize() int {
	return g.field.GetSize()
}

func (g Game) GetCurrentPlayer() common.IPlayer {
	return g.players[g.playerIdx]
}
