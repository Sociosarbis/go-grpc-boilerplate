package config

import (
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

func New() (AppConfig, error) {
	var cfg AppConfig
	err := errgo.Wrap(cleanenv.ReadEnv(&cfg), "cleanenv.ReadEnv")
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
