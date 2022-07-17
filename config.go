package main

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Socrates string `toml:"socrates"`
	Diogenes string `toml:"diogenes"`
	Benjamin string `toml:"benjamin"`
}

// Possible security issue: embedding into binary - don't share binaries publically; treat binaries as secrets!
//go:embed config.toml
var configFile string

func LoadConfigMust() Config {
	var config Config

	_, err := toml.Decode(configFile, &config)
	if err != nil {
		panic(fmt.Sprintf("error decoding toml config: %s", err))
	}

	return config
}
