package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.M) {
	// Store environment
	env := os.Environ()
	os.Clearenv()
	defer func() {
		os.Clearenv()
		for _, e := range env {
			idx := strings.Index(e, "=")
			key := e[:idx]
			value := e[idx+1:]
			err := os.Setenv(key, value)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	code := t.Run()
	os.Exit(code)
}

func TestReadEnvAndMerge(t *testing.T) {
	defer os.Clearenv()
	var c, cmp Config
	// first: full env read test
	os.Clearenv()
	c = NewEmptyConfig()
	os.Setenv("TICTACTOE_GAME_VARIANT", "http")
	os.Setenv("TICTACTOE_PLAYERS_SIDES", "XO")
	os.Setenv("TICTACTOE_FIELD_SIZE", "3")
	os.Setenv("TICTACTOE_FIELD_WINSEQ", "3")
	c.ReadEnvAndMerge()
	cmp = Config{
		GameVariant:  "http",
		PlayersSides: "XO",
		FieldSize:    3,
		FieldWinSeq:  3,
	}
	if c != cmp {
		t.Fatal("full env read test FAIL")
	}
	// second: merge env test
	os.Clearenv()
	c = NewEmptyConfig()
	c.PlayersSides = "XO"
	os.Setenv("TICTACTOE_GAME_VARIANT", "cli")
	os.Setenv("TICTACTOE_FIELD_SIZE", "4")
	os.Setenv("TICTACTOE_FIELD_WINSEQ", "4")
	c.ReadEnvAndMerge()
	cmp = Config{
		GameVariant:  "cli",
		PlayersSides: "XO",
		FieldSize:    4,
		FieldWinSeq:  4,
	}
	if c != cmp {
		t.Fatal("merge env test FAIL")
	}
}

var conf = `{
	"game_variant":"http","players_sides":"OX","field_size":3,"field_winseq":3
}`
var confMerge = `{
	"game_variant":"http","players_sides":"OX","field_size":4
}`

func TestJsonMerge(t *testing.T) {
	var Buffer *bytes.Buffer
	var c, cmp Config
	// first: read full json config
	Buffer = bytes.NewBufferString(conf)
	c = NewEmptyConfig()
	c.ReadJsonAndMerge(Buffer)
	cmp = Config{
		GameVariant:  "http",
		PlayersSides: "OX",
		FieldSize:    3,
		FieldWinSeq:  3,
	}
	if c != cmp {
		t.Fatal("full json config read FAIL: ", c, cmp, Buffer)
	}
	// second: merge json config
	Buffer = bytes.NewBufferString(confMerge)
	c = NewEmptyConfig()
	c.FieldWinSeq = 3
	c.ReadJsonAndMerge(Buffer)
	cmp = Config{
		GameVariant:  "http",
		PlayersSides: "OX",
		FieldSize:    4,
		FieldWinSeq:  3,
	}
	if c != cmp {
		t.Fatal("merge json config read FAIL")
	}
}
