package config

import (
	"bytes"
	"fmt"
	"testing"
)

var conf = []byte(`{
  "game_variant":"http",
  "players_sides":"OX",
  "filed_size":3,
  "field_winSeq":3
}`)

func TestReadJson(t *testing.T) {
	Buffer := bytes.NewBuffer(conf)
	c := NewEmptyConfig()
	c.ReadJson(Buffer)
	fmt.Println(c)
}

func TestReadEnvAndMerge(t *testing.T) {
	c := NewEmptyConfig()
	c.ReadEnv()
	fmt.Println(c)
}
