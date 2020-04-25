package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	MirrorRegistry struct {
		Server string `env:"MIRROR_REGISTRY_SERVER"`
		Auth   struct {
			Username string `env:"MIRROR_REGISTRY_USERNAME"`
			Password string `env:"MIRROR_REGISTRY_PASSWORD"`
		}
		Docker string `env:"DOCKER"`
		Gcr    string `env:"GCR"`
		Quay   string `env:"QUAY"`
	}
}

var (
	c    Config
	once sync.Once
)

func Parser() *Config {
	once.Do(func() {
		if err := env.Parse(&c); err != nil {
			panic(err)
		}
	})
	return &c
}
