package config

import (
	"bytes"
	"fmt"
	"testing"
)

var conf = []byte(`{
  "game_variant":"http",
  "players_sides":"OX",
  "field_size":3,
  "field_winseq":3
}`)

var confMerge = []byte(`{
  "game_variant":"http",
  "players_sides":"OX",
  "field_size":3,
}`)

func TestJsonMerge(t *testing.T) {
	Buffer := bytes.NewBuffer(conf)
	c := NewEmptyConfig()
	c.FieldSize = 4
	c.FieldWinSeq = 4
	fmt.Println(c)
	c.ReadJsonAndMerge(Buffer)
	fmt.Println(c)
}

func TestReadJsonAndMerge(t *testing.T) {
	Buffer := bytes.NewBuffer(conf)
	c := NewEmptyConfig()
	c.ReadJsonAndMerge(Buffer)
	fmt.Println(c)
}

func TestReadEnvAndMerge(t *testing.T) {
	c := NewEmptyConfig()
	c.ReadEnvAndMerge()
	fmt.Println(c)
}
