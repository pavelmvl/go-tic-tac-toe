package player

import (
	"fmt"
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
