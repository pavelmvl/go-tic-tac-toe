package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Config struct {
	GameVariant  string `json:"game_variant"`
	PlayersSides string `json:"players_sides"`
	FieldSize    int    `json:"field_size"`
	FieldWinSeq  int    `json:"field_winseq"`
}

var envLoookup = map[string][]string{
	"game_variant":  {"TICTACTOE_GAME_VARIANT", "GAME_VARIANT"},
	"players_sides": {"TICTACTOE_PLAYERS_SIDES", "PLAYERS_SIDES"},
	"field_size":    {"TICTACTOE_FIELD_SIZE", "FIELD_SIZE"},
	"field_winseq":  {"TICTACTOE_FIELD_WINSEQ", "FIELD_WINSEQ"},
}

func lookup(key string) (string, bool) {
	slice, ok := envLoookup[key]
	if ok {
		for _, v := range slice {
			value, ok := os.LookupEnv(v)
			if ok {
				return value, ok
			}
		}
	}
	return "", false
}

func NewEmptyConfig() Config {
	return Config{}
}

func NewConfigFromEnv() Config {
	cfg := NewEmptyConfig()
	cfg.ReadEnvAndMerge()
	return cfg
}

func NewConfigFromJsonFile(filename string) (Config, error) {
	cfg := NewEmptyConfig()
	file, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	cfg.ReadJsonAndMerge(file)
	return cfg, nil
}

func (c *Config) ReadEnvAndMerge() {
	if v, ok := lookup("game_variant"); ok {
		c.GameVariant = v
	}
	if v, ok := lookup("players_sides"); ok {
		c.PlayersSides = v
	}
	if v, ok := lookup("field_size"); ok {
		if int_v, err := strconv.Atoi(v); err == nil {
			c.FieldSize = int_v
		} else {
			fmt.Println(err)
		}
	}
	if v, ok := lookup("field_winseq"); ok {
		if int_v, err := strconv.Atoi(v); err == nil {
			c.FieldWinSeq = int_v
		} else {
			fmt.Println(err)
		}
	}
}

func (c *Config) ReadJsonAndMerge(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}
