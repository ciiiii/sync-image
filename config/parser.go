package config

import (
	"os"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	c    Config
	once sync.Once
)

func Parser() Config {
	once.Do(func() {
		if os.Getenv("MIRROR_REGISTRY_USERNAME") == "" {
			if err := godotenv.Load(); err != nil {
				panic(err)
			}
		}
		if err := env.Parse(&c); err != nil {
			panic(err)
		}
	})
	return c
}

