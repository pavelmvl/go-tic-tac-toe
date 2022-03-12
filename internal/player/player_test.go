package player

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	// t.Fatal("not implemented")
	var player Player
	player = New("X")
	if player.Priority != 0 || player.GetMark() != 'X' {
		t.Fatal("do not pass cmd -p X")
	}
	player = New("x")
	if player.Priority != 0 || player.GetMark() != 'X' {
		t.Fatal("do not pass cmd -p x")
	}
	player = New("O")
	if player.Priority != 1 || player.GetMark() != 'O' {
		t.Fatal("do not pass cmd -p o")
	}
	player = New("")
	if player.Priority != 1 || player.GetMark() != 'O' {
		t.Fatal("do not pass cmd -p P")
	}
}
