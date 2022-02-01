package initPlayer

import (
	"flag"
	"fmt"
)

type Player struct {
	Priority int
	Mark     rune
}

func (this Player) String() string {
	return fmt.Sprintf("Player %d marking '%c'", this.Priority, this.Mark)
}

func InitPlayer() Player {
	var player Player
	p := flag.String("p", "", "Setup player mark. Default is 'x'")
	flag.Parse()
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	switch *p {
	case "x", "X":
		player.Priority = 0
		player.Mark = 'X'
	default:
		player.Priority = 1
		player.Mark = 'O'
	}
	return player
}
