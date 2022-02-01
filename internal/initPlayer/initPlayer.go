package initPlayer

import (
	"fmt"
)

type Player struct {
	Priority int
	Mark     rune
}

func (this Player) String() string {
	return fmt.Sprintf("Player %d marking '%c'", this.Priority, this.Mark)
}

func New(mark string) Player {
	var player Player
	switch mark {
	case "x", "X":
		player.Priority = 0
		player.Mark = 'X'
	default:
		player.Priority = 1
		player.Mark = 'O'
	}
	return player
}
