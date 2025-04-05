package config

import (
	"log/slog"

	"github.com/BurntSushi/toml"
)

func New() (config Config) {
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		slog.Error("read config error")

		panic(err)
	}

	return config
}
