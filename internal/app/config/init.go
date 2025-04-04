package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

func New() (config Config) {
	err := cleanenv.ReadConfig("./config.yml", &config)
	if err != nil {
		slog.Error("read config error")

		panic(err)
	}

	return config
}
