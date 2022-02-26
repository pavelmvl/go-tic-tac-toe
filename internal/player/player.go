package player

import (
	"fmt"
	"go-tic-tac-toe/internal/common"
)

type Player struct {
	Priority int
	mark     rune
}

func (p Player) String() string {
	return fmt.Sprintf("'Player %d marking '%c''", p.Priority, p.mark)
}

func (p Player) GetMark() rune {
	return p.mark
}

func NewIPlayer(mark string) common.IPlayer {
	return New(mark)
}

func New(mark string) Player {
	var player Player
	switch mark {
	case "x", "X":
		player.Priority = 0
		player.mark = 'X'
	default:
		player.Priority = 1
		player.mark = 'O'
	}
	return player
}
