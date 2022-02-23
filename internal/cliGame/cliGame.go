package cliGame

import (
	"fmt"
	"go-tic-tac-toe/internal/common"
	"io"
)

type cliGame struct {
	game   common.IGame
	curin  io.Reader
	curout io.Writer
	curerr io.Writer
}

func (c cliGame) Run() {
	var col int
	var row int
	for {
		for {
			fmt.Fprintf(
				c.curout,
				"(%d)(%c) Enter <column> <row>:",
				c.game.GetIter(),
				c.game.GetCurrentPlayer().GetMark(),
			)
			fmt.Fscan(c.curin, &col, &row)
			if err := c.game.NextMove(col, row); err != nil {
				if c.curerr != nil {
					fmt.Fprintln(c.curerr, err)
				}
				fmt.Fprintln(c.curout, "Try enter againg")
				continue
			}
			break
		}
		fmt.Fprint(c.curout, c.game.ToString())
		winnerString, winnerErr := c.game.GetWinnerString()
		if winnerErr == nil {
			fmt.Fprintln(c.curout, winnerString)
			return
		}
	}
}

func PlayCliGame(g common.IGame, in io.Reader, out io.Writer, err io.Writer) {
	c := cliGame{
		game:   g,
		curin:  in,
		curout: out,
		curerr: err,
	}
	c.Run()
}
