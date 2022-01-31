package initPlayer

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestInitPlayer(t *testing.T) {
	// t.Fatal("not implemented")
	var player Player
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0], "-p", "X"}
	player = InitPlayer()
	if player.Priority != 0 || player.Mark != 'X' {
		t.Fatal("do not pass cmd -p X")
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0], "-p", "x"}
	player = InitPlayer()
	if player.Priority != 0 || player.Mark != 'X' {
		t.Fatal("do not pass cmd -p x")
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0], "-p", "o"}
	player = InitPlayer()
	if player.Priority != 1 || player.Mark != 'O' {
		t.Fatal("do not pass cmd -p o")
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0], "-p", "P"}
	player = InitPlayer()
	if player.Priority != 1 || player.Mark != 'O' {
		t.Fatal("do not pass cmd -p P")
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0]}
	player = InitPlayer()
	fmt.Println(player)
}
